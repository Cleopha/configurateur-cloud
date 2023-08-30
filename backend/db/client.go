package db

import (
	"backend/ent"
	"context"
	"fmt"
)

type Client struct {
	*ent.Client
}

func New() (*Client, error) {
	p := postgres{}

	client, err := p.Client()
	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(context.TODO()); err != nil {
		fmt.Println("ok")
		return nil, err
	}

	return client, nil
}
