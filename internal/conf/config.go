package conf

import (
	_ "embed"
	"gopkg.in/yaml.v3"
)

var (
	//go:embed conf.yaml
	configFile string
)

type Config struct {
	WorkDir string       `yaml:"work_dir"`
	App     AppConfig    `yaml:"app"`
	Domain  DomainConfig `yaml:"domain"`
	Repo    RepoConfig   `yaml:"repo"`
}

func GetConfig() Config {
	var config Config
	err := yaml.Unmarshal([]byte(configFile), &config)
	if err != nil {
		panic(err)
	}
	return config
}

type AppConfig struct {
	RelativePath string `yaml:"relative_path"`
}

type DomainConfig struct {
	RelativePath string `yaml:"relative_path"`
}

type RepoConfig struct {
	RelativePath string `yaml:"relative_path"`
}
