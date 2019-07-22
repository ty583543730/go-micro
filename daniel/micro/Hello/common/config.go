package common

type GlobalCfg struct {
	Consul map[string]interface{} `json:"consul"`
	Mysql map[string]interface{} `json:"mysql"`
}