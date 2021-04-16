package conf

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	Database DatabaseConf `yaml:"database"`
}

type DatabaseConf struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func NewConf(filePath string) *Conf {
	conf := &Conf{}
	return conf.loadConf(filePath)
}

func (c *Conf) loadConf(filePath string) *Conf {

	if filePath == "from os vars" {
		c.Database.Dbname = os.Getenv("DB_NAME")
		c.Database.Host = os.Getenv("DB_HOST")
		c.Database.Port = os.Getenv("DB_PORT")
		c.Database.User = os.Getenv("DB_USER")
		c.Database.Password = os.Getenv("DB_PASSWORD")
		return c
	} else {

		yamlFile, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Printf("yamlFile.Get err   #%v ", err)
		}
		err = yaml.Unmarshal(yamlFile, &c)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}

		return c
	}
}
