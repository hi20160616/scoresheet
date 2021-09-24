package configs

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

type ProjectName string

type Config struct {
	ProjectName ProjectName
	RootPath    string
	Debug       bool
	Verbose     bool
	LogName     string
	Raw         []byte
	Err         error
	DBPath      string   `json:"dbPath"`
	DBType      []string `json:"dbType"`
}

func NewConfig(projectName ProjectName) *Config {
	return setRootPath(&Config{ProjectName: projectName}).load()
}

func setRootPath(cfg *Config) *Config {
	cfg.RootPath, cfg.Err = os.Getwd()
	if cfg.Err != nil {
		return cfg
	}
	if strings.Contains(os.Args[0], ".test") {
		return rootPath4Test(cfg)
	}
	return cfg
}

func rootPath4Test(cfg *Config) *Config {
	ps := strings.Split(cfg.RootPath, string(cfg.ProjectName))
	n := 0
	if len(ps) > 1 {
		n = strings.Count(ps[1], string(os.PathSeparator))
	}
	for i := 0; i < n; i++ {
		cfg.RootPath = filepath.Join("../", "./")
	}
	return cfg
}

func (c *Config) load() *Config {
	if c.Err != nil {
		return c
	}
	cfgFile := filepath.Join(c.RootPath, "configs", "configs.json")
	c.Raw, c.Err = os.ReadFile(cfgFile)
	if c.Err != nil {
		if errors.Is(c.Err, os.ErrNotExist) {
			c.Err = errors.WithMessage(c.Err, "ReadFile error: no configs.json")
		}
		return c
	}
	cfgTemp := &Config{}
	if c.Err = json.Unmarshal(c.Raw, cfgTemp); c.Err != nil {
		c.Err = errors.WithMessage(c.Err, "Unmarshal configs.json error")
		return c
	}
	c.Debug = cfgTemp.Debug
	c.Verbose = cfgTemp.Verbose
	c.LogName = cfgTemp.LogName
	c.ProjectName = cfgTemp.ProjectName
	c.DBPath = cfgTemp.DBPath
	c.DBType = cfgTemp.DBType
	return c
}
