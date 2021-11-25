package checker

import (
	"context"
	authz "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	. "github.com/onsi/gomega"
	"testing"
)

func TestCheckNoStore(t *testing.T) {
	RegisterTestingT(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	uut := NewServer()
	req := &authz.CheckRequest{Attributes: &authz.AttributeContext{
		Request: &authz.AttributeContext_Request{
			Http: &authz.AttributeContext_HttpRequest{
				Path: "/foo.com",
				},
			},
		},
	}

	//resp, err := uut.Check(ctx, req)
	path, err := uut.Check2(ctx, req)

	Expect(err).To(BeNil())
	Expect(path).To(Equal("blah"))
}
