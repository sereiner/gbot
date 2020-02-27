package gbot

import (
	"flag"
	container2 "gbot/container"
	"gbot/gbot"
	"gbot/kernel"
	"gbot/log"
	"github.com/spf13/viper"
)

var session string

type GbotApp struct {
	container        *container2.Container
	config           *viper.Viper
	serviceProviders []gbot.Resover
	*option
}

func init() {
	flag.StringVar(&session, "session", "", "set session name")
}

type App interface {
	Serve()
}

func New(opts ...Option) App {
	return newApp(opts...)
}

func newApp(opts ...Option) *GbotApp {
	flag.Parse()
	a := &GbotApp{
		container: container2.NewContainer(),
		config:    viper.New(),
		option:    &option{path: "./tmp/", download: DownloadConf{EmoticonPath: "./tmp/emoticons/"}},
		serviceProviders: []gbot.Resover{
			log.NewLog,
		},
	}

	for _, op := range opts {
		op(a.option)
	}
	a.session = session

	a.initializeConfig()

	kernel.NewKernel(a).BootStrap()

	return a
}

func (a *GbotApp) initializeConfig() {
	a.config.Set("path", a.option.path)
	a.config.Set("download.emoticon_path", a.option.download.EmoticonPath)
}

func (a *GbotApp) Session() string {
	return a.session
}

func (a *GbotApp) Config() *viper.Viper {
	return a.config
}
func (a *GbotApp) Cache() {

}

func (a *GbotApp) ServiceProviders() []gbot.Resover {
	return a.serviceProviders
}

func (a *GbotApp) Container() *container2.Container {
	return a.container
}

func (a *GbotApp) Serve() {

}
