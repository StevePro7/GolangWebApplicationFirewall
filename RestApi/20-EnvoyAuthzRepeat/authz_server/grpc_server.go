package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gogo/googleapis/google/rpc"
	"log"
	"net"
	"os"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	_ "google.golang.org/grpc/codes"
	//healthpb "google.golang.org/grpc/health/grpc_health_v1"
	_ "google.golang.org/grpc/status"

	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"

	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	//envoytype "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	//"github.com/gogo/googleapis/google/rpc"
)

var (
	grpcport = flag.String("grpcport", ":50051", "grpcport")
	//conn     *grpc.ClientConn
	//hs       *health.Server
)

const (
//address string = ":50051"
)

//type healthServer struct{}

//func (s *healthServer) Check(ctx context.Context, in *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
//	log.Printf("Handling grpc Check request")
//	// yeah, right, open 24x7, like 7-11
//	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
//}
//
//func (s *healthServer) Watch(in *healthpb.HealthCheckRequest, srv healthpb.Health_WatchServer) error {
//	return status.Error(codes.Unimplemented, "Watch is not implemented")
//}

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

	log.Println("Auth header = " + authHeader)

	if len(splitToken) == 2 {
		log.Println("Split token" + splitToken[1])
	}

	//if len(splitToken) == 2 {
	//	token := splitToken[1]
	//
	//	if token == "foo" {
	//		return &auth.CheckResponse{
	//			Status: &rpcstatus.Status{
	//				Code: int32(rpc.OK),
	//			},
	//			HttpResponse: &auth.CheckResponse_OkResponse{
	//				OkResponse: &auth.OkHttpResponse{
	//					Headers: []*core.HeaderValueOption{
	//						{
	//							Header: &core.HeaderValue{
	//								Key:   "x-custom-header-from-authz",
	//								Value: "some value",
	//							},
	//						},
	//					},
	//				},
	//			},
	//		}, nil
	//	} else {
	//		return &auth.CheckResponse{
	//			Status: &rpcstatus.Status{
	//				Code: int32(rpc.PERMISSION_DENIED),
	//			},
	//			HttpResponse: &auth.CheckResponse_DeniedResponse{
	//				DeniedResponse: &auth.DeniedHttpResponse{
	//					Status: &envoytype.HttpStatus{
	//						Code: envoytype.StatusCode_Unauthorized,
	//					},
	//					Body: "PERMISSION_DENIED",
	//				},
	//			},
	//		}, nil
	//	}
	//}
	//return &auth.CheckResponse{
	//	Status: &rpcstatus.Status{
	//		Code: int32(rpc.UNAUTHENTICATED),
	//	},
	//	HttpResponse: &auth.CheckResponse_DeniedResponse{
	//		DeniedResponse: &auth.DeniedHttpResponse{
	//			Status: &envoytype.HttpStatus{
	//				Code: envoytype.StatusCode_Unauthorized,
	//			},
	//			Body: "Authorization Header malformed or not provided",
	//		},
	//	},
	//}, nil

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

func main() {

	flag.Parse()

	if *grpcport == "" {
		fmt.Fprintln(os.Stderr, "missing -grpcport flag (:50051)")
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
	//healthpb.RegisterHealthServer(s, &healthServer{})

	log.Printf("Starting gRPC Server at %s", *grpcport)
	s.Serve(lis)

}
