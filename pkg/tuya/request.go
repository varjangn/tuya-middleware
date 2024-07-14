package tuya

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"go.uber.org/zap"
)

func (c *TuyaClient) GetBaseUrl(version float32) string {
	baseUrl := strings.TrimSuffix(c.cfg.Host, "/")
	verStr := fmt.Sprintf("v%.1f", version)
	return fmt.Sprintf("%s/%s", baseUrl, verStr)
}

func (c *TuyaClient) DoRequest(url, method string, body []byte) ([]byte, error) {

	RefreshingWg.Wait()

	token := c.GetActiveToken()
	if token.AccessToken == "" {
		c.logger.Errorw("request_err",
			zap.String("error", "access token is not generated"))
		err := fmt.Errorf("access token is not generated")
		return nil, err
	}

	c.logger.Infow("url_hit",
		zap.String("method", method),
		zap.String("url", url),
		zap.String("token", token.AccessToken))

	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	BuildRequestHeader(req, body, c.cfg.ClientId, c.cfg.Secret)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.logger.Errorw("request_err",
			zap.String("error", err.Error()))
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Errorw("decode_err",
			zap.String("error", err.Error()))
		return nil, err
	}

	return bs, nil
}
