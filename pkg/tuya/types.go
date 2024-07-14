package tuya

type TuyaTimestamp int64

type DeviceStatus struct {
	Code  string `json:"code"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type Device struct {
	Id          string         `json:"id"`
	Name        string         `json:"name"`
	Uid         string         `json:"uid"`
	LocalKey    string         `json:"local_key"`
	Category    string         `json:"category"`
	ProductId   string         `json:"product_id"`
	ProductName string         `json:"product_name"`
	Sub         bool           `json:"sub"`
	UUID        string         `json:"uuid"`
	OwnerId     string         `json:"owner_id"`
	Online      bool           `json:"online"`
	Status      []DeviceStatus `json:"status"`
	ActiveTime  TuyaTimestamp  `json:"active_time"`
	BizType     int64          `json:"biz_type"`
	Icon        string         `json:"icon"`
	Ip          string         `json:"ip"`
	CreateTime  TuyaTimestamp  `json:"create_time,omitempty"`
	UpdateTime  TuyaTimestamp  `json:"update_time,omitempty"`
	TimeZone    string         `json:"time_zone,omitempty"`
}

type SubDevice struct {
	Id         string        `json:"id"`
	Name       string        `json:"name"`
	Online     bool          `json:"online"`
	OwnerId    string        `json:"owner_id"`
	Category   string        `json:"category"`
	ProductId  string        `json:"product_id"`
	ActiveTime TuyaTimestamp `json:"active_time"`
	UpdateTime TuyaTimestamp `json:"update_time"`
}

type FactoryInfo struct {
	Id   string `json:"id"`
	UUID string `json:"uuid"`
	Sn   string `json:"sn"`
	Mac  string `json:"mav"`
}

type DeviceUser struct {
	DeviceId string        `json:"device_id"`
	NickName string        `json:"nick_name"`
	Sex      int           `json:"sex"`
	Birthday TuyaTimestamp `json:"birthday"`
	Height   int           `json:"height"`
	Weight   int           `json:"weight"`
	Contact  string        `json:"contact"`
}

// multi outlet device name
type MODeviceName struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}
