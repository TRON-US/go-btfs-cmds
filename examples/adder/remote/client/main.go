package main

import (
	"context"
	"os"

	"github.com/TRON-US/go-btfs-cmds/examples/adder"

	//cmdkit "github.com/ipfs/go-ipfs-cmdkit"
	cmds "github.com/TRON-US/go-btfs-cmds"
	cli "github.com/TRON-US/go-btfs-cmds/cli"
	http "github.com/TRON-US/go-btfs-cmds/http"
)

func main() {
	// parse the command path, arguments and options from the command line
	req, err := cli.Parse(context.TODO(), os.Args[1:], os.Stdin, adder.RootCmd)
	if err != nil {
		panic(err)
	}

	// create http rpc client
	client := http.NewClient(":6798")

	// send request to server
	res, err := client.Send(req)
	if err != nil {
		panic(err)
	}

	req.Options["encoding"] = cmds.Text

	// create an emitter
	re, err := cli.NewResponseEmitter(os.Stdout, os.Stderr, req)
	if err != nil {
		panic(err)
	}

	// copy received result into cli emitter
	if pr, ok := req.Command.PostRun[cmds.CLI]; ok {
		err = pr(res, re)
	} else {
		err = cmds.Copy(re, res)
	}
	if err != nil {
		re.CloseWithError(err)
	}

	os.Exit(re.Status())
}
