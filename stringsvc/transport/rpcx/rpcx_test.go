package rpcx_test

import (
	"context"
	"log"
	"play/stringsvc/transport"
	"testing"

	"github.com/smallnest/rpcx/client"
)

func TestUppercase(t *testing.T) {
	sd, _ := client.NewPeer2PeerDiscovery("tcp@localhost:6000", "")
	xcli := client.NewXClient("StringServiceRPCX", client.Failtry, client.RandomSelect, sd, client.DefaultOption)
	defer xcli.Close()

	resp := &transport.UppercaseResponse{}
	err := xcli.Call(context.Background(), "Uppercase", transport.UppercaseRequest{S: "hello world"}, resp)
	if err != nil {
		t.Error(err)
	}

	log.Printf("resp.V=%s", resp.V)
}
