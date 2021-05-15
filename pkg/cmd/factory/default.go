package factory

import (
	"github.com/bharat-rajani/GoHunter/internal/config"
	"github.com/bharat-rajani/GoHunter/pkg/cmdutil"
	"github.com/bharat-rajani/GoHunter/pkg/iostreams"
	"net/http"
)

// This is a factory of factory
func New() *cmdutil.CMDContainerFactory {
	io := iostreams.System()
	var cachedConfig *config.Configuration
	var configError error
	configFunc := func() (*config.Configuration, error) {
		if cachedConfig != nil || configError != nil {
			return cachedConfig, configError
		}
		cachedConfig, configError = config.ReadConfig()
		return cachedConfig, configError
	}
	httpClientFunc := func() (*http.Client, error) {
		cfg, err := configFunc()
		if err != nil {
			return nil, err
		}

		return NewHttpClient(io, cfg), nil
	}

	return &cmdutil.CMDContainerFactory{io, configFunc, httpClientFunc}
}
