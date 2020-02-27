package kernel

import (
	"github.com/sereiner/gbot/gbot"
	"github.com/sereiner/gbot/session"
	"github.com/sereiner/library/file"
	"os"
)

type Kernel struct {
	gbot gbot.Gbot
}

func NewKernel(gbot gbot.Gbot) *Kernel {
	return &Kernel{gbot: gbot}
}

func (k *Kernel) BootStrap() {
	k.checkEnvironment()
	k.registerProviders()
	k.initializeConfig()
	k.prepareSession()
	k.initializePath()

}
func (k *Kernel) checkEnvironment() {}
func (k *Kernel) registerProviders() {

	for _, r := range k.gbot.ServiceProviders() {
		k.gbot.Container().Register(r(k.gbot))
	}

}
func (k *Kernel) bootstrapException() {}
func (k *Kernel) initializeConfig() {
	if !file.Exists(k.gbot.Config().GetString("path")) {
		if err := os.MkdirAll(k.gbot.Config().GetString("path"), os.ModePerm); err != nil {
			panic(err)
		}
	}

	k.gbot.Config().Set("storage", "collection")

}
func (k *Kernel) prepareSession() {

	s := session.NewSession(k.gbot)
	sessionKey := s.CurrentSession()

	k.gbot.Config().Set("session", sessionKey)
	k.gbot.Config().Set("session_key", "session."+sessionKey)

}
func (k *Kernel) initializePath() {
	if !file.Exists(k.gbot.Config().GetString("path") + "/cookies") {
		if err := os.MkdirAll(k.gbot.Config().GetString("path")+"/cookies", 0755); err != nil {
			panic(err)
		}
	}

	if !file.Exists(k.gbot.Config().GetString("path") + "/users") {
		if err := os.MkdirAll(k.gbot.Config().GetString("path")+"/users", 0755); err != nil {
			panic(err)
		}
	}

	if !file.Exists(k.gbot.Config().GetString("download.emoticon_path")) {
		if err := os.MkdirAll(k.gbot.Config().GetString("download.emoticon_path"), 0755); err != nil {
			panic(err)
		}
	}

	k.gbot.Config().Set("cookie_file", k.gbot.Config().GetString("path")+"/cookies/"+k.gbot.Session())
	k.gbot.Config().Set("user_path", k.gbot.Config().GetString("path")+"/users/")

}
