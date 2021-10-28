package checker

import (
	"context"
	. "github.com/onsi/gomega"
	"testing"

	authz "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
)

func TestCheckNoStore(t *testing.T) {
	RegisterTestingT(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//uut := NewServer(ctx)
	uut := NewServer()
	req := &authz.CheckRequest{}
	val := uut.Check(ctx, req)

	Expect(val).To(Equal(18))
}