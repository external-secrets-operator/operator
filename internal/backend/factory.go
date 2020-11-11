package backend

import (
	"context"
	"fmt"
	"github.com/external-secrets-operator/operator/api/v1alpha1"
	"github.com/external-secrets-operator/operator/generated/api"
	"google.golang.org/grpc"
)

type Factory interface {
	EstablishConnection(ctx context.Context, backend string, spec v1alpha1.ExternalBackendSpec, opts ...grpc.DialOption) error
	CloseConnection(ctx context.Context, backend string) error
	GetSecrets(ctx context.Context, keys []string) ([]api.SecretsResponse_Secret, error)
}

type factory struct {
	conns map[string]connection
}

type connection struct {
	client *grpc.ClientConn
	spec   v1alpha1.ExternalBackendSpec
}

func (f *factory) EstablishConnection(ctx context.Context, backend string, spec v1alpha1.ExternalBackendSpec, opts ...grpc.DialOption) error {
	ctx, _ = context.WithTimeout(ctx, spec.Timeout.Duration)
	opts = append(opts, grpc.WithBlock())
	if spec.Insecure {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.DialContext(ctx, spec.Target, opts...)
	if err != nil {
		return fmt.Errorf("failed to connect to %s backend; %w", backend, err)
	}
	f.conns[backend] = connection{
		client: conn,
		spec:   spec,
	}
	return nil
}

func (f *factory) CloseConnection(ctx context.Context, backend string) error {
	panic("implement me")
}

func (f *factory) GetSecrets(ctx context.Context, keys []string) ([]api.SecretsResponse_Secret, error) {
	panic("implement me")
}

func NewFactory() Factory {
	return &factory{
		conns: make(map[string]connection),
	}
}
