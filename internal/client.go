package internal

import (
	"context"

	containerd "github.com/containerd/containerd/v2/client"
	"github.com/containerd/containerd/v2/pkg/namespaces"
)

func NewContainerdClient(containerdOpts ...containerd.Opt) (*containerd.Client, context.Context, error) {
	ctx := namespaces.WithNamespace(context.Background(), vesselNamespace)

	client, err := containerd.New(containerdSock, containerdOpts...)
	if err != nil {
		return nil, nil, err
	}

	return client, ctx, nil
}

func ShutdownClient(client *containerd.Client) error {
	return client.Close()
}
