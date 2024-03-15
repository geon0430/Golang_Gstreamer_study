package util

import (
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
)

func ReadConfig(path string) (*PipelineConfig, error) {
	var config PipelineConfig
	cfg, err := ini.Load(path)
	if err != nil {
		return nil, err
	}

	err = cfg.Section("General").MapTo(&config.General)
	if err != nil {
		return nil, err
	}
	err = cfg.Section("Encoder").MapTo(&config.Encoder)
	if err != nil {
		return nil, err
	}
	err = cfg.Section("Decoder").MapTo(&config.Decoder)
	if err != nil {
		return nil, err
	}
	err = cfg.Section("NOX").MapTo(&config.Decoder)
	if err != nil {
		return nil, err
	}

	logrus.Debugf("ReadConfig successfully")

	return &config, nil
}
