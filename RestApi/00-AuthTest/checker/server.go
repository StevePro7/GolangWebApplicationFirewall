package checker

import (
	authz "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type authServer struct {
	age int
}

// NewServer creates a new authServer and returns a pointer to it.
//func NewServer(ctx context.Context) *authServer {
func NewServer() *authServer {
	s := &authServer{ 10 }
	return s
}

// Check applies the currently loaded policy to a network request and renders a policy decision.
func (as *authServer) Check(ctx context.Context, req *authz.CheckRequest) int {
	fields := log.Fields{
		"context":         ctx,
		"Req.Method":      req.GetAttributes().GetRequest().GetHttp().GetMethod(),
		"Req.Path":        req.GetAttributes().GetRequest().GetHttp().GetPath(),
		"Req.Protocol":    req.GetAttributes().GetRequest().GetHttp().GetProtocol(),
		"Req.Source":      req.GetAttributes().GetSource(),
		"Req.Destination": req.GetAttributes().GetDestination(),
	}
	log.WithFields(fields).Debug("Check start")

	return 17
}