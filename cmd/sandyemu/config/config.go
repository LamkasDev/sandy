package config

import (
	"encoding/json"
	"os"
	"os/user"
	"path/filepath"
)

type PlaystationConfig struct {
	CPU PlaystationConfigCPU `json:"cpu"`
}

type PlaystationConfigCPU struct {
	Speed uint32 `json:"speed"`
}

func NewPlaystationConfig() PlaystationConfig {
	return PlaystationConfig{
		CPU: PlaystationConfigCPU{
			Speed: 1000,
		},
	}
}

func ReadConfig(config *PlaystationConfig) error {
	u, _ := user.Current()
	path := filepath.Join(u.HomeDir, "Documents", "sandy", "config", "emu.json")
	raw, err := os.ReadFile(path)
	if err != nil {
		if raw, err = json.Marshal(*config); err != nil {
			return err
		}
		if err = os.WriteFile(path, raw, 0755); err != nil {
			return err
		}
	}
	err = json.Unmarshal(raw, config)
	if err != nil {
		return err
	}

	return nil
}
