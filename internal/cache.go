package internal

import (
	"context"

	"github.com/valkey-io/valkey-go"
)

var defaultConfig = valkey.ClientOption{
	InitAddress: []string{"127.0.0.1:6379"},
}

type Cache interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Close()
}

type Valkey struct {
	client valkey.Client
}

func (v *Valkey) Get(key string) (string, error) {
	ctx := context.Background()

	return v.client.Do(ctx, v.client.B().Get().Key(key).Build()).ToString()
}

func (v *Valkey) Set(key string, value string) error {
	ctx := context.Background()

	err := v.client.Do(ctx, v.client.B().Set().Key(key).Value(value).Nx().Build()).Error()
	if err != nil {
		return err
	}

	v.client.B().Flushall()
	return nil
}

func (v *Valkey) Close() {
	v.client.Close()
}

func NewValKey() (Cache, error) {
	opts := defaultConfig

	client, err := valkey.NewClient(opts)
	if err != nil {
		return &Valkey{}, err
	}

	return &Valkey{
		client: client,
	}, nil
}
