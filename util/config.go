package util

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

type Config struct { //配置文件信息
	Cron   string `yaml:"cron"`
	Qcloud struct {
		SecretId  string `yaml:"secretId"`
		SecretKey string `yaml:"secretKey"`
		List      []struct {
			Domain    string `yaml:"domain"`
			SubDomain string `yaml:"subDomain"`
			Type      string `yaml:"type"`
		} `yaml:"list"`
	} `yaml:"qcloud"`
}

func Setting() (newConfig *Config) {
	config, err := ioutil.ReadFile("config.yaml")
	CheckErr(err)
	newConfig = new(Config)
	err = yaml.Unmarshal(config, newConfig)
	CheckErr(err)
	return newConfig
}
