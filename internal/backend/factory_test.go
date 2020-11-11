package backend

import (
	"context"
	"github.com/external-secrets-operator/operator/api/v1alpha1"
	"github.com/external-secrets-operator/operator/generated/api"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/test/bufconn"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"net"
	"testing"
	"time"
)

func Test_ShouldEstablishConnection(t *testing.T) {
	f := NewFactory()
	err := f.EstablishConnection(context.TODO(), "test", v1alpha1.ExternalBackendSpec{
		Target:   "",
		Timeout:  metav1.Duration{Duration: 1 * time.Second},
		Insecure: true,
	}, grpc.WithContextDialer(dialer(statsHandler{})))
	assert.NoError(t, err)
}

func dialer(statsHandler stats.Handler) func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer(grpc.StatsHandler(statsHandler))
	api.RegisterBackendService(server, &api.BackendService{})

	go func() {
		if err := server.Serve(listener); err != nil {
			panic(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

type statsHandler struct{}

func (s statsHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	log.Printf("%+v", info)
	return ctx
}

func (s statsHandler) HandleRPC(ctx context.Context, rpcStats stats.RPCStats) {
	log.Printf("%+v", rpcStats)
}

func (s statsHandler) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	log.Printf("%+v", info)
	return ctx
}

func (s statsHandler) HandleConn(ctx context.Context, connStats stats.ConnStats) {
	log.Printf("%+v", connStats)
}
