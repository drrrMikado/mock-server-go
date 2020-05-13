package conf

import (
	"bytes"
	"io/ioutil"

	"github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

var (
	Conf = &Config{}
)

type Config struct {
	Mysql *mysql.Config
}

func Init() error {
	b, err := ioutil.ReadFile("cmd/conf.yml")
	if err != nil {
		return err
	}
	return yaml.NewDecoder(bytes.NewReader(b)).Decode(&Conf)
}
