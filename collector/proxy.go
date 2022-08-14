package collector

import (
	"DouBanUpdater/config"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Proxy struct {
	Anonymous  string `json:"anonymous"`
	CheckCount int    `json:"check_count"`
	FailCount  int    `json:"fail_count"`
	HTTPS      bool   `json:"https"`
	LastStatus bool   `json:"last_status"`
	LastTime   string `json:"last_time"`
	Proxy      string `json:"proxy"`
	Region     string `json:"region"`
	Source     string `json:"source"`
}

func GetProxy() (Proxy, error) {
	req, err := http.NewRequest("GET", config.Config.Proxy.Url, nil)
	if err != nil {
		return Proxy{}, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Proxy{}, err
	}
	var proxy Proxy
	root, err := ioutil.ReadAll(res.Body)
	if err := json.Unmarshal(root, &proxy); err != nil {
		return Proxy{}, err
	}
	return proxy, nil
}
