package main

import (
	"context"
	nethttp "net/http"

	"github.com/TRON-US/go-btfs-cmds/examples/adder"

	http "github.com/TRON-US/go-btfs-cmds/http"
)

type env struct{}

func (env) Context() context.Context {
	return context.TODO()
}

func main() {
	h := http.NewHandler(env{}, adder.RootCmd, http.NewServerConfig())

	// create http rpc server
	err := nethttp.ListenAndServe(":6798", h)
	if err != nil {
		panic(err)
	}
}
