package ymgo

import (
	"go.mongodb.org/mongo-driver/mongo"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
	"ymgo/pkg/errors"
	"ymgo/pkg/options"
)

type Client struct {
	client *mongo.Client
}

func NewClient(ctx Context, opts ...*options.ClientOptions) (client *Client, err error) {

	clientOptions := make([]*mongoOptions.ClientOptions, len(opts))
	for i, opt := range opts {
		if opt == nil {
			err = errors.ErrNotSupportedOption
			return
		}
		clientOption := mongoOptions.ClientOptions(*opt)
		clientOptions[i] = &clientOption
	}

	var c *mongo.Client
	c, err = mongo.Connect(ctx, clientOptions...)
	if err != nil {
		return
	}

	client = &Client{
		client: c,
	}
	return
}

// Database TODO Support DatabaseOption
func (c *Client) Database(name string) (database *Database) {
	d := c.client.Database(name)
	database = &Database{
		database: d,
	}
	return
}
