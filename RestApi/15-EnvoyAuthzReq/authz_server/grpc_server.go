package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"

	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	envoytype "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	"github.com/gogo/googleapis/google/rpc"
)

var (
	grpcport = flag.String("grpcport", ":50051", "grpcport")
)

type AuthorizationServer struct{}

func (a *AuthorizationServer) Check(ctx context.Context, req *auth.CheckRequest) (*auth.CheckResponse, error) {
	log.Println(">>> stevepro Authorization called check() start")

	log.Println(">>> stevepro request information: start")

	//fields := mylog.Fields{
	//	"context":         ctx,
	//	"":      req.GetAttributes().GetRequest().GetHttp().GetMethod(),
	//	"Req.Path":        req.GetAttributes().GetRequest().GetHttp().GetPath(),
	//	"Req.Protocol":    req.GetAttributes().GetRequest().GetHttp().GetProtocol(),
	//	"Req.Source":      req.GetAttributes().GetSource(),
	//	"Req.Destination": req.GetAttributes().GetDestination(),
	//}
	//mylog.WithFields(fields).Debug("stevepro Check start")

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

	log.Println(">>> stevepro request information: -end-")

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
	if len(splitToken) == 2 {
		token := splitToken[1]

		if token == "foo" {
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
		} else {
			return &auth.CheckResponse{
				Status: &rpcstatus.Status{
					Code: int32(rpc.PERMISSION_DENIED),
				},
				HttpResponse: &auth.CheckResponse_DeniedResponse{
					DeniedResponse: &auth.DeniedHttpResponse{
						Status: &envoytype.HttpStatus{
							Code: envoytype.StatusCode_Unauthorized,
						},
						Body: "PERMISSION_DENIED",
					},
				},
			}, nil

		}

	}
	return &auth.CheckResponse{
		Status: &rpcstatus.Status{
			Code: int32(rpc.UNAUTHENTICATED),
		},
		HttpResponse: &auth.CheckResponse_DeniedResponse{
			DeniedResponse: &auth.DeniedHttpResponse{
				Status: &envoytype.HttpStatus{
					Code: envoytype.StatusCode_Unauthorized,
				},
				Body: "Authorization Header malformed or not provided",
			},
		},
	}, nil
}

func main() {

	flag.Parse()

	if *grpcport == "" {
		_, err := fmt.Fprintln(os.Stderr, "missing -grpcport flag (:50051)")
		if err != nil {
			log.Fatal(err)
			return
		}
		flag.Usage()
		os.Exit(2)
	}

	lis, err := net.Listen("tcp", *grpcport)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{grpc.MaxConcurrentStreams(10)}
	opts = append(opts)

	s := grpc.NewServer(opts...)

	auth.RegisterAuthorizationServer(s, &AuthorizationServer{})

	log.Printf("Starting gRPC Server at %s", *grpcport)
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
		return
	}

}
