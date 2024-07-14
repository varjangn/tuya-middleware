package tuya

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	ExpireTime   int    `json:"expire_time"`
	RefreshToken string `json:"refresh_token"`
	UID          string `json:"uid"`
}

type TokenResponse struct {
	Result  Token `json:"result"`
	Success bool  `json:"success"`
	T       int64 `json:"t"`
}

var (
	token           Token
	TokenExpiringAt time.Time
	RefreshingWg    sync.WaitGroup
)

func BuildRequestHeader(req *http.Request, body []byte, clientId, secret string) {
	req.Header.Set("client_id", clientId)
	req.Header.Set("sign_method", "HMAC-SHA256")

	ts := fmt.Sprint(time.Now().UnixNano() / 1e6)
	req.Header.Set("t", ts)

	if token.AccessToken != "" {
		req.Header.Set("access_token", token.AccessToken)
	}

	sign := buildSign(req, body, ts, clientId, secret)
	req.Header.Set("sign", sign)
}

func buildSign(req *http.Request, body []byte, t string, clientId string, secret string) string {
	headers := getHeaderStr(req)
	urlStr := getUrlStr(req)
	contentSha256 := Sha256(body)
	stringToSign := req.Method + "\n" + contentSha256 + "\n" + headers + "\n" + urlStr
	signStr := clientId + token.AccessToken + t + stringToSign
	sign := strings.ToUpper(HmacSha256(signStr, secret))
	return sign
}

func Sha256(data []byte) string {
	sha256Contain := sha256.New()
	sha256Contain.Write(data)
	return hex.EncodeToString(sha256Contain.Sum(nil))
}

func getUrlStr(req *http.Request) string {
	url := req.URL.Path
	keys := make([]string, 0, 10)

	query := req.URL.Query()
	for key := range query {
		keys = append(keys, key)
	}
	if len(keys) > 0 {
		url += "?"
		sort.Strings(keys)
		for _, keyName := range keys {
			value := query.Get(keyName)
			url += keyName + "=" + value + "&"
		}
	}

	if url[len(url)-1] == '&' {
		url = url[:len(url)-1]
	}
	return url
}

func getHeaderStr(req *http.Request) string {
	signHeaderKeys := req.Header.Get("Signature-Headers")
	if signHeaderKeys == "" {
		return ""
	}
	keys := strings.Split(signHeaderKeys, ":")
	headers := ""
	for _, key := range keys {
		headers += key + ":" + req.Header.Get(key) + "\n"
	}
	return headers
}

func HmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func (c *TuyaClient) FetchToken() error {

	method := "GET"
	body := []byte(``)
	req, _ := http.NewRequest(method, c.cfg.Host+"/v1.0/token?grant_type=1", bytes.NewReader(body))

	BuildRequestHeader(req, body, c.cfg.ClientId, c.cfg.Secret)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.logger.Infow("request err", zap.String("err", err.Error()))
		return err
	}
	defer resp.Body.Close()
	bs, _ := io.ReadAll(resp.Body)
	ret := TokenResponse{}
	err = json.Unmarshal(bs, &ret)
	if err != nil {
		c.logger.Infow("decode err", zap.String("err", err.Error()))
		return err
	}

	if v := ret.Result.AccessToken; v != "" {
		tkn := SetActiveToken(&ret)
		expireTime := tkn.ExpireTime
		// expireTime := 10

		TokenExpiringAt = time.Now().Local().Add(time.Second * time.Duration(expireTime))
		c.logger.Infow("token", zap.String("token", v))
	}

	return nil
}

func (c *TuyaClient) RefreshToken() error {
	if token.RefreshToken == "" {
		return fmt.Errorf("initial token not found")
	}
	body := []byte(``)

	tokenUrl := fmt.Sprintf("%s/v1.0/token/%s", c.cfg.Host, token.RefreshToken)

	c.logger.Infof("refresh token url: %s", tokenUrl)

	req, err := http.NewRequest("GET", tokenUrl, bytes.NewReader(body))
	if err != nil {
		c.logger.Infow("request err", zap.String("err", err.Error()))
		return err
	}

	req.Header.Set("client_id", c.cfg.ClientId)
	req.Header.Set("sign_method", "HMAC-SHA256")

	ts := fmt.Sprint(time.Now().UnixNano() / 1e6)
	req.Header.Set("t", ts)
	headers := getHeaderStr(req)
	urlStr := getUrlStr(req)
	contentSha256 := Sha256(body)
	stringToSign := req.Method + "\n" + contentSha256 + "\n" + headers + "\n" + urlStr
	signStr := c.cfg.ClientId + ts + "" + stringToSign
	sign := strings.ToUpper(HmacSha256(signStr, c.cfg.Secret))
	req.Header.Set("sign", sign)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.logger.Infow("request err", zap.String("err", err.Error()))
		return err
	}
	defer resp.Body.Close()
	bs, _ := io.ReadAll(resp.Body)

	tokenResponse := new(TokenResponse)
	err = json.Unmarshal(bs, tokenResponse)
	if err != nil {
		c.logger.Infow("decode err", zap.String("err", err.Error()))
		return err
	}

	if v := tokenResponse.Result.AccessToken; v != "" {
		tkn := SetActiveToken(tokenResponse)

		expireTime := tkn.ExpireTime
		// expireTime := 10

		TokenExpiringAt = time.Now().Local().Add(time.Second * time.Duration(expireTime))

		c.logger.Infof("new access token: %s", tkn.AccessToken)
		c.logger.Infof("new expiry: %s", TokenExpiringAt)
		return nil
	}

	return fmt.Errorf("failed to refresh token")
}

func (c *TuyaClient) GetActiveToken() Token {
	return token
}

func SetActiveToken(tokenResp *TokenResponse) Token {
	token = tokenResp.Result
	return token
}

func (c *TuyaClient) AutoRefreshToken() error {
	for {
		RefreshingWg.Add(1)

		if token.AccessToken == "" {
			c.logger.Info("generating first token")
			if err := c.FetchToken(); err != nil {
				RefreshingWg.Done()
				return err
			}
		} else if token.AccessToken != "" && token.RefreshToken != "" {
			c.logger.Infow("%s token is expired", token.AccessToken)
			if err := c.RefreshToken(); err != nil {
				RefreshingWg.Done()
				return err
			}
		} else {
			RefreshingWg.Done()
			break
		}

		currentTime := time.Now().Local()
		diff := TokenExpiringAt.Sub(currentTime)
		diff = diff - (time.Duration(120) * time.Second)
		c.logger.Infof("sleeping for %f hours", diff.Hours())
		RefreshingWg.Done()
		time.Sleep(diff)
	}
	return fmt.Errorf("can not auto refresh token")
}
