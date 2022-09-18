package main

import (
	"fmt"
	"github.com/memmaker/net/webdav"
	"net/http"
)

func main() {
	config := ParseConfig()

	wdHandler := &webdav.Handler{
		Prefix: config.Prefix,
		FileSystem: &Dir{
			Config: config,
		},
		LockSystem: webdav.NewMemLS(),
		Logger: func(request *http.Request, err error) {
			if err != nil {
				fmt.Println(err)
			}
		},
	}

	a := &App{
		Config:  config,
		Handler: wdHandler,
	}

	http.Handle("/", wrapRecovery(NewBasicAuthWebdavHandler(a), config))
	connAddr := fmt.Sprintf("%s:%s", config.Address, config.Port)

	if config.TLS != nil {
		fmt.Println("Listening on https://" + connAddr)
		fmt.Println(http.ListenAndServeTLS(connAddr, config.TLS.CertFile, config.TLS.KeyFile, nil))
	} else {
		fmt.Println("Listening on http://" + connAddr)
		fmt.Println(http.ListenAndServe(connAddr, nil))
	}
}

func wrapRecovery(handler http.Handler, config *Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				switch t := err.(type) {
				case string:
					fmt.Println(t)
				case error:
					fmt.Println(t.Error())
				}
			}
		}()

		if len(config.Cors.Origin) > 0 {
			w.Header().Set("Access-Control-Allow-Origin", config.Cors.Origin)
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Allow-Methods", "*")
			if config.Cors.Credentials {
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}
		}

		handler.ServeHTTP(w, r)
	})
}
