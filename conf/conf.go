package conf

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"

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
	filename := "cmd/conf.yml"
	workDirEnv := os.Getenv("WORK_DIR")
	if workDirEnv != "" {
		workDirEnv = strings.TrimSuffix(workDirEnv, "/")
		filename = workDirEnv + "/cmd/conf.yml"
	}
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return yaml.NewDecoder(bytes.NewReader(b)).Decode(&Conf)
}
