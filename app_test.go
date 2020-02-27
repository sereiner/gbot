package gbot

import (
	"testing"
)

func TestNewApp(t *testing.T) {
	a := New()
	a.Serve()
}
