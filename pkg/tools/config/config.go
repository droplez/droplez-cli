package config

import (
	"fmt"
	"io/ioutil"

	"github.com/droplez/droplez-go-proto/pkg/studio/projects"
	"github.com/droplez/droplez-go-proto/pkg/studio/versions"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Project  *projects.ProjectInfo `yaml:"project"`
	Versions []*versions.VersionInfo
}

var getFileName = func(path string) string {
	return fmt.Sprintf("%s/.droplez/droplez.yaml", path)
}

func ReadConfig(path string) (config *Config, err error) {
	file, err := ioutil.ReadFile(getFileName(path))
	if err != nil {
		return
	}

	config = &Config{}
	if err = yaml.Unmarshal(file, &config); err != nil {
		return
	}

	return
}

func WriteConfig(path string, conf *Config) (err error) {
	out, err := yaml.Marshal(conf)
	if err != nil {
		return
	}

	if err = ioutil.WriteFile(getFileName(path), out, 0755); err != nil {
		return
	}

	fmt.Println(string(out))

	return
}
