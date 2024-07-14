package tuya

import (
	"encoding/json"
	"fmt"
	"net/url"

	"go.uber.org/zap"
)

/*
Query the device details, including attributes and the latest status of a specified device.
https://developer.tuya.com/en/docs/cloud/device-management?id=K9g6rfntdz78a#title-1-Get%20device%20details
*/
func (c *TuyaClient) GetDevice(deviceId string) (*Device, error) {
	var body []byte
	baseURL := c.GetBaseUrl(1.0)
	endpointURL := fmt.Sprintf("%s/devices/%s", baseURL, deviceId)
	response, err := c.DoRequest(endpointURL, "GET", body)
	if err != nil {
		return nil, err
	}

	respBody := new(DeviceResponse)
	err = json.Unmarshal(response, respBody)
	if err != nil {
		c.logger.Errorw("json_decode_err",
			zap.String("error", err.Error()))
	}
	if !respBody.Success {
		c.logger.Errorf("tuya_success_false")
		return nil, fmt.Errorf("success false for response")
	}

	return &respBody.Result, nil
}

/*
Query the list of devices available to a specified user, including device attributes and the latest status.
https://developer.tuya.com/en/docs/cloud/device-management?id=K9g6rfntdz78a#title-10-Get%20a%20list%20of%20devices%20under%20a%20specified%20user
*/
func (c *TuyaClient) GetUserDevices(userId string, queryParams map[string]string) ([]Device, error) {

	var body []byte
	baseURL := c.GetBaseUrl(1.0)
	encodedParams := url.Values{}
	for key, value := range queryParams {
		encodedParams.Add(key, value)
	}
	endpointURL := fmt.Sprintf("%s/users/%s?%s", baseURL, userId, encodedParams.Encode())
	response, err := c.DoRequest(endpointURL, "GET", body)
	if err != nil {
		return nil, err
	}

	respBody := new(UserDeviceResponse)

	err = json.Unmarshal(response, respBody)
	if err != nil {
		c.logger.Errorw("json_decode_err",
			zap.String("error", err.Error()))
	}
	if !respBody.Success {
		c.logger.Errorf("tuya_success_false")
		return nil, fmt.Errorf("success false for response")
	}

	return respBody.Result, nil
}

/*
Query a list of devices parameters
https://developer.tuya.com/en/docs/cloud/device-management?id=K9g6rfntdz78a#title-19-Get%20a%20list%20of%20devices
*/
func (c *TuyaClient) GetDevices(pageNo, PageSize int, queryParams map[string]string) (*DevicesResult, error) {

	var body []byte
	baseURL := c.GetBaseUrl(1.0)
	encodedParams := url.Values{}
	for key, value := range queryParams {
		encodedParams.Add(key, value)
	}
	endpointURL := fmt.Sprintf("%s/devices?page_no=%d&page_size=%d&%s", baseURL, pageNo, PageSize, encodedParams.Encode())
	response, err := c.DoRequest(endpointURL, "GET", body)
	if err != nil {
		return nil, err
	}

	respBody := new(DevicesResponse)

	err = json.Unmarshal(response, respBody)
	if err != nil {
		c.logger.Errorw("json_decode_err",
			zap.String("error", err.Error()))
	}
	if !respBody.Success {
		c.logger.Errorf("tuya_success_false")
		return nil, fmt.Errorf("success false for response")
	}

	return &respBody.Result, nil
}

/*
Modify the name of a data point
https://developer.tuya.com/en/docs/cloud/device-management?id=K9g6rfntdz78a#title-28-Modify%20the%20name%20of%20a%20data%20point
*/
func (c *TuyaClient) ModifyDPName(deviceId, functionCode, newName string) (bool, error) {

	baseURL := c.GetBaseUrl(1.0)
	endpointURL := fmt.Sprintf("%s/devices/%s/functions/%s", baseURL, deviceId, functionCode)

	payload := map[string]string{
		"name": newName,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		c.logger.Errorw("json_encode_err", zap.String("error", err.Error()))
		return false, err
	}
	response, err := c.DoRequest(endpointURL, "GET", body)
	if err != nil {
		return false, err
	}

	respBody := new(BooleanResponse)
	err = json.Unmarshal(response, respBody)
	if err != nil {
		c.logger.Errorw("json_decode_err",
			zap.String("error", err.Error()))
	}
	if !respBody.Success {
		c.logger.Errorf("tuya_success_false")
		return false, fmt.Errorf("success false for response")
	}

	return respBody.Result, nil
}

/*
Restore to factory defaults
https://developer.tuya.com/en/docs/cloud/device-management?id=K9g6rfntdz78a#title-42-Restore%20to%20factory%20defaults
*/
func (c *TuyaClient) FactoryResetDevice(deviceId string) (bool, error) {
	var body []byte
	baseURL := c.GetBaseUrl(1.0)
	endpointURL := fmt.Sprintf("%s/devices/%s/reset-factory", baseURL, deviceId)

	response, err := c.DoRequest(endpointURL, "PUT", body)
	if err != nil {
		return false, err
	}

	respBody := new(BooleanResponse)
	err = json.Unmarshal(response, respBody)
	if err != nil {
		c.logger.Errorw("json_decode_err",
			zap.String("error", err.Error()))
	}
	if !respBody.Success {
		c.logger.Errorf("tuya_success_false")
		return false, fmt.Errorf("success false for response")
	}

	return respBody.Result, nil
}

/*
Delete a specified device
https://developer.tuya.com/en/docs/cloud/device-management?id=K9g6rfntdz78a#title-49-Delete%20a%20specified%20device
*/
func (c *TuyaClient) DeleteDevice(deviceId string) (bool, error) {

	var body []byte
	baseURL := c.GetBaseUrl(1.0)
	endpointURL := fmt.Sprintf("%s/devices/%s", baseURL, deviceId)

	response, err := c.DoRequest(endpointURL, "DELETE", body)
	if err != nil {
		return false, err
	}

	respBody := new(BooleanResponse)
	err = json.Unmarshal(response, respBody)
	if err != nil {
		c.logger.Errorw("json_decode_err",
			zap.String("error", err.Error()))
	}
	if !respBody.Success {
		c.logger.Errorf("tuya_success_false")
		return false, fmt.Errorf("success false for response")
	}

	return respBody.Result, nil
}

/*
Query sub devices
https://developer.tuya.com/en/docs/cloud/device-management?id=K9g6rfntdz78a#title-55-Query%20a%20list%20of%20devices%20under%20a%20gateway
*/
func (c *TuyaClient) GetSubDevices(deviceId string) ([]SubDevice, error) {
	var body []byte
	baseURL := c.GetBaseUrl(1.0)
	endpointURL := fmt.Sprintf("%s/devices/%s/sub-devices", baseURL, deviceId)

	response, err := c.DoRequest(endpointURL, "DELETE", body)
	if err != nil {
		return []SubDevice{}, err
	}
	respBody := new(SubDeviceResponse)
	err = json.Unmarshal(response, respBody)
	if err != nil {
		c.logger.Errorw("json_decode_err",
			zap.String("error", err.Error()))
	}
	if !respBody.Success {
		c.logger.Errorf("tuya_success_false")
		return []SubDevice{}, fmt.Errorf("success false for response")
	}
	return respBody.Result, nil
}

/*
Query factory info
https://developer.tuya.com/en/docs/cloud/device-management?id=K9g6rfntdz78a#title-63-Query%20the%20factory%20information%20of%20a%20device
*/
func (c *TuyaClient) GetFactoryInfo(deviceIds string) ([]FactoryInfo, error) {
	var body []byte
	baseURL := c.GetBaseUrl(1.0)
	endpointURL := fmt.Sprintf("%s/devices/factory-infos?device_ids=%s", baseURL, deviceIds)
	response, err := c.DoRequest(endpointURL, "DELETE", body)
	if err != nil {
		return []FactoryInfo{}, err
	}
	respBody := new(FactoryInfoResponse)
	err = json.Unmarshal(response, respBody)
	if err != nil {
		c.logger.Errorw("json_decode_err",
			zap.String("error", err.Error()))
	}
	if !respBody.Success {
		c.logger.Errorf("tuya_success_false")
		return []FactoryInfo{}, fmt.Errorf("success false for response")
	}
	return respBody.Result, nil
}

/*
Modify a device name
https://developer.tuya.com/en/docs/cloud/device-management?id=K9g6rfntdz78a#title-71-Modify%20a%20device%20name
*/
func (c *TuyaClient) SetDeviceName(deviceId, newName string) (bool, error) {

	baseURL := c.GetBaseUrl(1.0)
	endpointURL := fmt.Sprintf("%s/devices/%s", baseURL, deviceId)

	payload := map[string]string{
		"name": newName,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		c.logger.Errorw("json_encode_err", zap.String("error", err.Error()))
		return false, err
	}
	response, err := c.DoRequest(endpointURL, "PUT", body)
	if err != nil {
		return false, err
	}

	respBody := new(BooleanResponse)
	err = json.Unmarshal(response, respBody)
	if err != nil {
		c.logger.Errorw("json_decode_err",
			zap.String("error", err.Error()))
	}
	if !respBody.Success {
		c.logger.Errorf("tuya_success_false")
		return false, fmt.Errorf("success false for response")
	}
	return respBody.Result, nil
}

/*
Add a user
https://developer.tuya.com/en/docs/cloud/device-management?id=K9g6rfntdz78a#title-78-Add%20a%20user
*/
func (c *TuyaClient) AddUser(deviceId string, userInfo map[string]interface{}) (string, error) {
	baseURL := c.GetBaseUrl(1.0)
	endpointURL := fmt.Sprintf("%s/devices/%s/user", baseURL, deviceId)

	_, ok := userInfo["nick_name"]
	if !ok {
		return "", fmt.Errorf("nick_name can not be empty")
	}
	_, ok = userInfo["sex"]
	if !ok {
		return "", fmt.Errorf("nick_name can not be empty")
	}

	body, err := json.Marshal(userInfo)
	if err != nil {
		c.logger.Errorw("json_encode_err", zap.String("error", err.Error()))
		return "", err
	}
	response, err := c.DoRequest(endpointURL, "POST", body)
	if err != nil {
		return "", err
	}

	respBody := new(StringResponse)
	err = json.Unmarshal(response, respBody)
	if err != nil {
		c.logger.Errorw("json_decode_err",
			zap.String("error", err.Error()))
	}
	if !respBody.Success {
		c.logger.Errorf("tuya_success_false")
		return "", fmt.Errorf("success false for response")
	}
	return respBody.Result, nil
}

/*
Modify a user
https://developer.tuya.com/en/docs/cloud/device-management?id=K9g6rfntdz78a#title-85-Modify%20a%20user
*/
func (c *TuyaClient) ModifyUser(deviceId, userId string, userInfo map[string]interface{}) (string, error) {
	baseURL := c.GetBaseUrl(1.0)
	endpointURL := fmt.Sprintf("%s/devices/%s/users/%s", baseURL, deviceId, userId)

	_, ok := userInfo["nick_name"]
	if !ok {
		return "", fmt.Errorf("nick_name can not be empty")
	}
	_, ok = userInfo["sex"]
	if !ok {
		return "", fmt.Errorf("nick_name can not be empty")
	}

	body, err := json.Marshal(userInfo)
	if err != nil {
		c.logger.Errorw("json_encode_err", zap.String("error", err.Error()))
		return "", err
	}
	response, err := c.DoRequest(endpointURL, "PUT", body)
	if err != nil {
		return "", err
	}

	respBody := new(StringResponse)
	err = json.Unmarshal(response, respBody)
	if err != nil {
		c.logger.Errorw("json_decode_err",
			zap.String("error", err.Error()))
	}
	if !respBody.Success {
		c.logger.Errorf("tuya_success_false")
		return "", fmt.Errorf("success false for response")
	}
	return respBody.Result, nil
}

/*
Delete device user
https://developer.tuya.com/en/docs/cloud/device-management?id=K9g6rfntdz78a#title-92-Delete%20a%20user
*/
func (c *TuyaClient) DeleteDeviceUser(deviceId, userId string) (bool, error) {
	baseURL := c.GetBaseUrl(1.0)
	endpointURL := fmt.Sprintf("%s/devices/%s/users/%s", baseURL, deviceId, userId)
	var body []byte
	response, err := c.DoRequest(endpointURL, "DELETE", body)
	if err != nil {
		return false, err
	}

	respBody := new(BooleanResponse)
	err = json.Unmarshal(response, respBody)
	if err != nil {
		c.logger.Errorw("json_decode_err",
			zap.String("error", err.Error()))
	}
	if !respBody.Success {
		c.logger.Errorf("tuya_success_false")
		return false, fmt.Errorf("success false for response")
	}

	return respBody.Result, nil
}

/*
Query user information
https://developer.tuya.com/en/docs/cloud/device-management?id=K9g6rfntdz78a#title-99-Query%20user%20information
*/
func (c *TuyaClient) GetDeviceUser(deviceId, userId string) (*DeviceUser, error) {
	baseURL := c.GetBaseUrl(1.0)
	endpointURL := fmt.Sprintf("%s/devices/%s/users/%s", baseURL, deviceId, userId)
	var body []byte
	response, err := c.DoRequest(endpointURL, "GET", body)
	if err != nil {
		return nil, err
	}
	respBody := new(DeviceUserResponse)
	err = json.Unmarshal(response, respBody)
	if err != nil {
		c.logger.Errorw("json_decode_err",
			zap.String("error", err.Error()))
	}
	if !respBody.Success {
		c.logger.Errorf("tuya_success_false")
		return nil, fmt.Errorf("success false for response")
	}

	return &respBody.Result, nil
}

/*
Query user information
https://developer.tuya.com/en/docs/cloud/device-management?id=K9g6rfntdz78a#title-99-Query%20user%20information
*/
func (c *TuyaClient) GetDeviceUsers(deviceId string) ([]DeviceUser, error) {
	baseURL := c.GetBaseUrl(1.0)
	endpointURL := fmt.Sprintf("%s/devices/%s/users", baseURL, deviceId)
	var body []byte
	response, err := c.DoRequest(endpointURL, "GET", body)
	if err != nil {
		return nil, err
	}
	respBody := new(DeviceUsersResponse)
	err = json.Unmarshal(response, respBody)
	if err != nil {
		c.logger.Errorw("json_decode_err",
			zap.String("error", err.Error()))
	}
	if !respBody.Success {
		c.logger.Errorf("tuya_success_false")
		return nil, fmt.Errorf("success false for response")
	}

	return respBody.Result, nil
}

/*
Modify names of a multi-outlet device
https://developer.tuya.com/en/docs/cloud/device-management?id=K9g6rfntdz78a#title-113-Modify%20names%20of%20a%20multi-outlet%20device
*/
func (c *TuyaClient) ModifyMODeviceName(deviceId, identifier, name string) (bool, error) {
	baseURL := c.GetBaseUrl(1.0)
	endpointURL := fmt.Sprintf("%s/devices/%s/multiple-name", baseURL, deviceId)

	payload := map[string]string{
		"identifier": identifier,
		"name":       name,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		c.logger.Errorw("json_encode_err", zap.String("error", err.Error()))
		return false, err
	}
	response, err := c.DoRequest(endpointURL, "PUT", body)
	if err != nil {
		return false, err
	}

	respBody := new(BooleanResponse)
	err = json.Unmarshal(response, respBody)
	if err != nil {
		c.logger.Errorw("json_decode_err",
			zap.String("error", err.Error()))
	}
	if !respBody.Success {
		c.logger.Errorf("tuya_success_false")
		return false, fmt.Errorf("success false for response")
	}
	return respBody.Result, nil
}

/*
Modify names of a multi-outlet device
https://developer.tuya.com/en/docs/cloud/device-management?id=K9g6rfntdz78a#title-113-Modify%20names%20of%20a%20multi-outlet%20device
*/
func (c *TuyaClient) GetMODeviceNames(deviceId, identifier, name string) ([]MODeviceName, error) {
	var body []byte
	baseURL := c.GetBaseUrl(1.0)
	endpointURL := fmt.Sprintf("%s/devices/%s/multiple-names", baseURL, deviceId)

	response, err := c.DoRequest(endpointURL, "GET", body)
	if err != nil {
		return []MODeviceName{}, err
	}

	respBody := new(MODeviceNamesResponse)
	err = json.Unmarshal(response, respBody)
	if err != nil {
		c.logger.Errorw("json_decode_err",
			zap.String("error", err.Error()))
	}
	if !respBody.Success {
		c.logger.Errorf("tuya_success_false")
		return []MODeviceName{}, fmt.Errorf("success false for response")
	}
	return respBody.Result, nil
}
