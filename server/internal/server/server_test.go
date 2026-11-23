package server

import (
	"context"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect/v2"
	"connectrpc.com/connect/v2/connecthttp"

	aludelv1 "github.com/KasumiMercury/aludeldrive/server/gen/aludel/v1"
	"github.com/KasumiMercury/aludeldrive/server/gen/aludel/v1/aludelv1connect"
)

func TestGetServerInfo(t *testing.T) {
	ts := httptest.NewServer(NewHandler("1.2.3"))
	defer ts.Close()

	client := aludelv1connect.NewServerServiceClient(connect.NewClient(connecthttp.NewTransport(ts.Client(), ts.URL)))

	res, err := client.GetServerInfo(context.Background(), &aludelv1.GetServerInfoRequest{})
	if err != nil {
		t.Fatal(err)
	}

	if got := res.GetVersion(); got != "1.2.3" {
		t.Fatalf("version = %q, want %q", got, "1.2.3")
	}
}
