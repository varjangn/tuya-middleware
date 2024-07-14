package tuya

type BaseResponse struct {
	Code    string        `json:"code,omitempty"`
	Success bool          `json:"success"`
	Msg     string        `json:"msg,omitempty"`
	T       TuyaTimestamp `json:"t,omitempty"`
}

type DeviceResponse struct {
	BaseResponse
	Result Device `json:"result"`
}

type UserDeviceResponse struct {
	BaseResponse
	Result []Device `json:"result"`
}

type DevicesResult struct {
	Total   int64    `json:"total"`
	Devices []Device `json:"devices"`
	LastId  string   `json:"last_id"`
}

type DevicesResponse struct {
	BaseResponse
	Result DevicesResult `json:"result"`
}

type BooleanResponse struct {
	BaseResponse
	Result bool `json:"result"`
}

type StringResponse struct {
	BaseResponse
	Result string `json:"result"`
}

type SubDeviceResponse struct {
	BaseResponse
	Result []SubDevice `json:"result"`
}

type FactoryInfoResponse struct {
	BaseResponse
	Result []FactoryInfo `json:"result"`
}

type DeviceUserResponse struct {
	BaseResponse
	Result DeviceUser `json:"result"`
}

type DeviceUsersResponse struct {
	BaseResponse
	Result []DeviceUser `json:"result"`
}

type MODeviceNamesResponse struct {
	BaseResponse
	Result []MODeviceName `json:"result"`
}
