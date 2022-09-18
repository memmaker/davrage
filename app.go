package main

import (
	"github.com/memmaker/net/webdav"
)

// App holds configuration information and the webdav handler.
type App struct {
	Config  *Config
	Handler *webdav.Handler
}
