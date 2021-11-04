package main

// #include "modsec.c"
import "C"

import (
	"encoding/json"
	"flag"
	"github.com/gogo/googleapis/google/rpc"
	"log"
	"net"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
)

var (
	grpcport = flag.String("grpcport", ":50051", "grpcport")
)

type AuthorizationServer struct{}

func (a *AuthorizationServer) Check(_ context.Context, req *auth.CheckRequest) (*auth.CheckResponse, error) {
	log.Println("Auth Server Check() method start")

	b, err := json.MarshalIndent(req.Attributes.Request.Http.Headers, "", "  ")
	if err == nil {
		log.Println("Inbound Headers: ")
		log.Println(string(b))
	}

	ct, err := json.MarshalIndent(req.Attributes.ContextExtensions, "", "  ")
	if err == nil {
		log.Println("Context Extensions: ")
		log.Println(string(ct))
	}

	authHeader, ok := req.Attributes.Request.Http.Headers["authorization"]
	var splitToken []string

	if ok {
		splitToken = strings.Split(authHeader, "Bearer ")
	}

	log.Println("Auth header= " + authHeader)
	if len(splitToken) == 2 {
		log.Println("Split token: " + splitToken[1])
	}

	myhttp := req.GetAttributes().GetRequest().GetHttp()
	log.Println(">>> GetId() >>  ", myhttp.GetId())
	log.Println(">>> GetMethod() >>  ", myhttp.GetMethod())
	log.Println(">>> GetHeaders() >>  ", myhttp.GetHeaders())
	log.Println(">>> GetPath() >>  ", myhttp.GetPath())
	log.Println(">>> GetHost() >>  ", myhttp.GetHost())
	log.Println(">>> GetScheme() >>  ", myhttp.GetScheme())
	log.Println(">>> GetQuery() >>  ", myhttp.GetQuery())
	log.Println(">>> GetFragment() >>  ", myhttp.GetFragment())
	log.Println(">>> GetSize() >>  ", myhttp.GetSize())
	log.Println(">>> GetProtocol() >>  ", myhttp.GetProtocol())
	log.Println(">>> GetBody() >>  ", myhttp.GetBody())
	log.Println(">>> GetRawBody() >>  ", myhttp.GetRawBody())

	log.Println("Auth Server Check() method -end-")
	return &auth.CheckResponse{
		Status: &rpcstatus.Status{
			Code: int32(rpc.OK),
		},
		HttpResponse: &auth.CheckResponse_OkResponse{
			OkResponse: &auth.OkHttpResponse{
				Headers: []*core.HeaderValueOption{
					{
						Header: &core.HeaderValue{
							Key:   "x-custom-header-from-authz",
							Value: "some value",
						},
					},
				},
			},
		},
	}, nil
}

func InitModSec() {
	log.Println("initModSec start")
	C.MyCInit()
	log.Println("initModSec -end-")
}

func main() {

	lis, err := net.Listen("tcp", *grpcport)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	InitModSec()

	opts := []grpc.ServerOption{grpc.MaxConcurrentStreams(10)}
	opts = append(opts)

	s := grpc.NewServer(opts...)
	auth.RegisterAuthorizationServer(s, &AuthorizationServer{})

	log.Printf("Starting gRPC Server at %s", *grpcport)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}
}
