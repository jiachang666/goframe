package config

type T struct {
	Mysql struct {
		User     string `json:"User"`
		Password string `json:"Password"`
		Host     string `json:"Host"`
		Port     int    `json:"Port"`
		Database string `json:"Database"`
	} `json:"mysql"`
	Redis struct {
		Addr     string `json:"Addr"`
		Password string `json:"Password"`
		DB       int    `json:"DB"`
	} `json:"redis"`
	Grpc struct {
		Host string `json:"Host"`
		Port int    `json:"Port"`
	} `json:"grpc"`
	Consul struct {
		Name string `json:"Name"`
		Host string `json:"Host"`
		Port int    `json:"Port"`
	} `json:"consul"`
	Aliyun struct {
		ID     string `json:"ID"`
		Secret string `json:"Secret"`
	} `json:"aliyun"`
	Qiniuyun struct {
		AK string `json:"AK"`
		SK string `json:"SK"`
	} `json:"qiniuyun"`
	Alipay struct {
		APPID      string `json:"APPID"`
		PrivateKey string `json:"PrivateKey"`
	} `json:"alipay"`
	Elasticsearch struct {
		Addr string `json:"Addr"`
	} `json:"elasticsearch"`
	Idcard struct {
		ID  string `json:"ID"`
		Key string `json:"Key"`
	} `json:"idcard"`
}

var Nas T
