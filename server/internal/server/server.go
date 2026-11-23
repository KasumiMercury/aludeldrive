package server

import (
	"context"
	"net/http"

	"connectrpc.com/connect/v2"
	"connectrpc.com/connect/v2/connecthttp"

	aludelv1 "github.com/KasumiMercury/aludeldrive/server/gen/aludel/v1"
	"github.com/KasumiMercury/aludeldrive/server/gen/aludel/v1/aludelv1connect"
)

func NewHandler(version string) http.Handler {
	srv := connect.NewServer()
	aludelv1connect.RegisterServerServiceHandler(srv, &serverService{version: version})

	mux := http.NewServeMux()
	connecthttp.Mount(mux, srv)

	return mux
}

type serverService struct {
	version string
}

func (s *serverService) GetServerInfo(context.Context, *aludelv1.GetServerInfoRequest) (*aludelv1.GetServerInfoResponse, error) {
	return &aludelv1.GetServerInfoResponse{Version: s.version}, nil
}

var _ aludelv1connect.ServerServiceHandler = (*serverService)(nil)
