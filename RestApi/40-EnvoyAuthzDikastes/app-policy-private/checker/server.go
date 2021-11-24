package checker

import (
	authz "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/status"
)

type authServer struct {
}

// NewServer creates a new authServer and returns a pointer to it.
func NewServer() *authServer {
	s := &authServer{}
	//go s.updateStores(ctx)
	return s
}

// Check applies the currently loaded policy to a network request and renders a policy decision.
func (as *authServer) Check(ctx context.Context, req *authz.CheckRequest) (*authz.CheckResponse, error) {

	log.WithFields(log.Fields{
		"context":         ctx,
		"Req.Method":      req.GetAttributes().GetRequest().GetHttp().GetMethod(),
		"Req.Path":        req.GetAttributes().GetRequest().GetHttp().GetPath(),
		"Req.Protocol":    req.GetAttributes().GetRequest().GetHttp().GetProtocol(),
		"Req.Source":      req.GetAttributes().GetSource(),
		"Req.Destination": req.GetAttributes().GetDestination(),
	}).Debug("Check start")
	resp := authz.CheckResponse{Status: &status.Status{Code: INTERNAL}}

	return &resp, nil
}

func (as *authServer) Check2(_ context.Context, req *authz.CheckRequest) (string, error) {

	path := req.GetAttributes().GetRequest().GetHttp().GetPath()

	return path, nil
}
