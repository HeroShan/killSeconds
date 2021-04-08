package main

import (
	"SecondsKill/server"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		server.GinHttp()
	}()

	http.ListenAndServe("127.0.0.1:6060", nil)
}
