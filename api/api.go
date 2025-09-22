package api

import "cli/app/config"

type Api struct {
	Config config.Config
}

func NewApi(config *config.Config) *Api {
	return &Api{
		Config: *config,
	}
}
