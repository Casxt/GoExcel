package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/Casxt/GoExcel/config"
	"github.com/Casxt/GoExcel/restfulexcel"
	"github.com/Casxt/TimeLine/api"
	"github.com/Casxt/TimeLine/components/index"
)

func route(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path //strings.SplitN(req.URL.Path, "?", 2)[0]
	log.Println(req.RemoteAddr[0:strings.LastIndex(req.RemoteAddr, ":")], req.Method, path)
	switch {
	case strings.HasPrefix(strings.ToLower(path), "/api"):
		api.Route(res, req)
	case strings.HasSuffix(strings.ToLower(path), ".xlsx"):
		restfulexcel.Route(res, req)
	default:
		index.Route(res, req)
	}
}

func main() {
	log.Println("Start Session server...")
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(route))

	if config.TLS.Cert != "" && config.TLS.Key != "" {
		log.Println("TLS Enable")
		log.Println("Start Https Server @ 443 ...")
		if err := http.ListenAndServeTLS(":443", config.TLS.Cert, config.TLS.Key, mux); err != nil {
			log.Fatalln(err.Error())
		}
	} else {
		log.Println("TLS Disable")
		log.Println("Start Http Server @ 80 ...")
		if err := http.ListenAndServe(":80", mux); err != nil {
			log.Fatalln(err.Error())
		}
	}
}
