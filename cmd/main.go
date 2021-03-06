package main

import (
	"github.com/rafaeljesus/composing-go-example/httpclient"
	"github.com/rafaeljesus/composing-go-example/user"
)

func main() {
	req := new(httpclient.Request)
	client := httpclient.New(req)

	_ = user.NewSync(client)
	_ = user.NewStore(client)

	// pass syncer and storer downstream to
	// http handler, messaging handler, etc...
}
