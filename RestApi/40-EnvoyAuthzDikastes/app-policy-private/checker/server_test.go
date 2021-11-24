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
	req := &authz.CheckRequest{}

	resp, err := uut.Check(ctx, req)

	//Expect(resp).To(BeNil())
	Expect(err).To(BeNil())
	Expect(resp.GetStatus().GetCode()).To(Equal(UNAVAILABLE))
}
