package gbot

import (
	"github.com/sereiner/gbot/container"
	"github.com/spf13/viper"
)

type Gbot interface {
	Session() string
	Config() *viper.Viper
	Cache()
	ServiceProviders() []Resover
	Container() *container.Container
}

type Resover func(Gbot) interface{}
