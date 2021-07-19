package adapters

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type ConfigStruct struct {
	Webport1         string `yaml:"Webport1"`
	Webport2         string `yaml:"Webport2"`
	DataBaseHost     string `yaml:"DataBaseHost"`
	DataBasePort     string `yaml:"DataBasePort"`
	DataBaseUsername string `yaml:"DataBaseUsername"`
	DataBaseDbname   string `yaml:"DataBaseDbname"`
	DataBaseSslmode  string `yaml:"DataBaseSslmode"`
	DataBasePassword string `yaml:"DataBasePassword"`
	AtExpires        string `yaml:"AtExpires"`
	RtExpires        string `yaml:"RtExpires"`
	AccessTokenKey   string `yaml:"AccessTokenKey"`
	RefreshTokenKey  string `yaml:"RefreshTokenKey"`
}

func (c *ConfigStruct) Parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}

func ParseConfig() ConfigStruct {
	yfile, err := ioutil.ReadFile("configs/config.yml")

	if err != nil {

		log.Fatal(err)
	}

	var config ConfigStruct
	if err := config.Parse(yfile); err != nil {
		log.Fatal(err)
	}
	return config
}
