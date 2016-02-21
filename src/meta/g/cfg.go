package g

import (
	"encoding/json"
	"fmt"
	"github.com/toolkits/file"
	"log"
	"sync"
)

type HttpConfig struct {
	Enabled bool   `json:"enabled"`
	Listen  string `json:"listen"`
}

type AgentDefaultConfig struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Tarball string `json:"tarball"`
	Md5     string `json:"md5"`
	Cmd     string `json:"cmd"`
}

type AgentOtherConfig struct {
	Prefix  string `json:"prefix"`
	Version string `json:"version"`
	Tarball string `json:"tarball"`
	Md5     string `json:"md5"`
	Cmd     string `json:"cmd"`
}

type InheritConfig struct {
	Default *AgentDefaultConfig `json:"default"`
	Others  []*AgentOtherConfig `json:"others"`
}
type GlobalConfig struct {
	Debug      bool             `json:"debug"`
	TarballDir string           `json:"tarballDir"`
	Http       *HttpConfig      `json:"http"`
	Agents     []*InheritConfig `json:"agents"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	lock       = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return config
}

func ParseConfig(cfg string) error {
	if cfg == "" {
		return fmt.Errorf("configuration file is blank")
	}

	if !file.IsExist(cfg) {
		return fmt.Errorf("configuration file is nonexistent")
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Println("ToTrimString file error")
		return err
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Println("json.Unmarshal file error")
		return err
	}

	lock.Lock()
	defer lock.Unlock()

	config = &c

	log.Println("read config file:", cfg, "successfully")

	return nil
}
