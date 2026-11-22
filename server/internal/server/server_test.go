package server

import (
	"context"
	"net/http/httptest"
	"testing"

	"connectrpc.com/connect"

	aludelv1 "github.com/KasumiMercury/aludeldrive/server/gen/aludel/v1"
	"github.com/KasumiMercury/aludeldrive/server/gen/aludel/v1/aludelv1connect"
)

func TestGetServerInfo(t *testing.T) {
	ts := httptest.NewServer(NewHandler("1.2.3"))
	defer ts.Close()

	client := aludelv1connect.NewServerServiceClient(ts.Client(), ts.URL)
	res, err := client.GetServerInfo(context.Background(), connect.NewRequest(&aludelv1.GetServerInfoRequest{}))
	if err != nil {
		t.Fatal(err)
	}
	if got := res.Msg.GetVersion(); got != "1.2.3" {
		t.Fatalf("version = %q, want %q", got, "1.2.3")
	}
}
