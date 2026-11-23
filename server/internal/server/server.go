package server

import (
	"context"
	"net/http"

	"connectrpc.com/connect"

	aludelv1 "github.com/KasumiMercury/aludeldrive/server/gen/aludel/v1"
	"github.com/KasumiMercury/aludeldrive/server/gen/aludel/v1/aludelv1connect"
)

func NewHandler(version string) http.Handler {
	mux := http.NewServeMux()
	mux.Handle(aludelv1connect.NewServerServiceHandler(&serverService{version: version}))
	return mux
}

type serverService struct {
	version string
}

func (s *serverService) GetServerInfo(context.Context, *connect.Request[aludelv1.GetServerInfoRequest]) (*connect.Response[aludelv1.GetServerInfoResponse], error) {
	return connect.NewResponse(&aludelv1.GetServerInfoResponse{Version: s.version}), nil
}

var _ aludelv1connect.ServerServiceHandler = (*serverService)(nil)
