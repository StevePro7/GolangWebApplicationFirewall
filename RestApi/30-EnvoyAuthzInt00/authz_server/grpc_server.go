package main

// #cgo CPPFLAGS: -I/usr/local/modsecurity/include
// #cgo LDFLAGS: /usr/local/modsecurity/lib/libmodsecurity.so
// #include "modsec.c"
import "C"

import (
	"encoding/json"
	"flag"
	"github.com/gogo/googleapis/google/rpc"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
	"unsafe"

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
	//log.Println(">>> GetId() >>  ", myhttp.GetId())
	//log.Println(">>> GetMethod() >>  ", myhttp.GetMethod())
	//log.Println(">>> GetHeaders() >>  ", myhttp.GetHeaders())
	//log.Println(">>> GetPath() >>  ", myhttp.GetPath())
	//log.Println(">>> GetHost() >>  ", myhttp.GetHost())
	//log.Println(">>> GetScheme() >>  ", myhttp.GetScheme())
	//log.Println(">>> GetQuery() >>  ", myhttp.GetQuery())
	//log.Println(">>> GetFragment() >>  ", myhttp.GetFragment())
	//log.Println(">>> GetSize() >>  ", myhttp.GetSize())
	//log.Println(">>> GetProtocol() >>  ", myhttp.GetProtocol())
	//log.Println(">>> GetBody() >>  ", myhttp.GetBody())
	//log.Println(">>> GetRawBody() >>  ", myhttp.GetRawBody())

	// TODO - is the same as previous
	// e.g.
	// /test/artists.php?artist=0+div+1+union%23foo*%2F*bar%0D%0Aselect%23foo%0D%0A1%2C2%2Ccurrent_user

	uri := myhttp.GetPath()
	log.Println(">>> url() >>  ", uri)
	httpMethod := myhttp.GetMethod()
	log.Println(">>> GetMethod() >>  ", httpMethod)

	tempProtocol := myhttp.GetProtocol()
	httpProtocol, httpVersion := split(tempProtocol, "/")
	log.Println(">>> 01() >>  ", httpProtocol)
	log.Println(">>> 02() >>  ", httpVersion)

	clientSocket := myhttp.GetHost()
	clientLink, tempPort := split(clientSocket, ":")
	clientPort, _ := strconv.Atoi(tempPort)
	log.Println(">>> 01() >>  ", clientLink)
	log.Println(">>> 02() >>  ", clientPort)

	serverSocket := myhttp.GetHost()
	serverLink, dataPort := split(serverSocket, ":")
	serverPort, _ := strconv.Atoi(dataPort)
	log.Println(">>> 01() >>  ", serverLink)
	log.Println(">>> 02() >>  ", serverPort)

	inter := modsec(uri, httpMethod, httpProtocol, httpVersion, clientLink, clientPort, serverLink, serverPort)
	if inter > 0 {
		log.Printf("==== Mod Security Blocked! ====")
		//http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		// TODO - why can't get this as before
		// e.g.
		// https://github.com/StevePro7/GolangWebApplicationFirewall/blob/main/RestApi/10-EnvoyAuthz/authz_server/grpc_server.go
		return &auth.CheckResponse{
			Status: &rpcstatus.Status{
				Code: int32(rpc.PERMISSION_DENIED),
			},
		}, nil
	}

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

func split(input, delim string) (left, right string) {
	split := strings.SplitN(input, delim, 2)
	left = split[0]
	right = split[1]
	return left, right
}
func InitModSec() {
	log.Println("initModSec start")
	C.MyCInit()
	log.Println("initModSec -end-")
}

func modsec(url, httpMethod, httpProtocol, httpVersion string, clientLink string, clientPort int, serverLink string, serverPort int) int {
	log.Println("modsec start ", url)
	Curi := C.CString(url)
	ChttpMethod := C.CString(httpMethod)
	ChttpProtocol := C.CString(httpProtocol)
	ChttpVersion := C.CString(httpVersion)
	CclientLink := C.CString(clientLink)
	CclientPort := C.int(clientPort)
	CserverLink := C.CString(serverLink)
	CserverPort := C.int(serverPort)

	defer C.free(unsafe.Pointer(Curi))
	defer C.free(unsafe.Pointer(ChttpMethod))
	defer C.free(unsafe.Pointer(ChttpProtocol))
	defer C.free(unsafe.Pointer(ChttpVersion))
	defer C.free(unsafe.Pointer(CclientLink))
	defer C.free(unsafe.Pointer(CserverLink))

	start := time.Now()
	inter := int(C.MyCProcess(Curi, ChttpMethod, ChttpProtocol, ChttpVersion, CclientLink, CclientPort, CserverLink, CserverPort))
	elapsed := time.Since(start)
	log.Printf("modsec()=%d, elapsed: %s", inter, elapsed)
	log.Println("modsec -end-")
	return inter
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
