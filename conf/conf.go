package conf

import (
	"bytes"
	"io/ioutil"

	"github.com/drrrMikado/mock-server-go/database/orm"
	"gopkg.in/yaml.v2"
)

var (
	Conf = &Config{}
)

type Config struct {
	Mysql *orm.Config
}

func Init() error {
	b, err := ioutil.ReadFile("cmd/conf.yml")
	if err != nil {
		return err
	}
	return yaml.NewDecoder(bytes.NewReader(b)).Decode(&Conf)
}
