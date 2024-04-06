package mango

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	client *mongo.Client
}

func NewClient(ctx Context, opts ...*options.ClientOptions) (client *Client, err error) {
	var c *mongo.Client
	c, err = mongo.Connect(ctx, opts...)
	if err != nil {
		return
	}

	client = &Client{
		client: c,
	}
	return
}

func (c *Client) Database(name string, opts ...*options.DatabaseOptions) (database *mongo.Database) {
	return c.client.Database(name, opts...)
}
