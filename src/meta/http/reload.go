package http

import (
	"../g"
	"net/http"
	"strings"
)

func configReloadRoutes(){

	http.HandleFunc("/config/reload", func(w http.ResponseWriter, r *http.Request){
		if strings.HasPrefix(r.RemoteAddr, "127.0.0.1"){
			err := g.ParseConfig(g.ConfigFile)
			AutoRender(w, g.Config(), err)
		}else {
			w.Write([]byte("no privilege"))
		}
	})
}