package mango

import (
	"github.com/YngMin/mango/pkg/options"
	"github.com/YngMin/mango/pkg/sliceutil"
	"go.mongodb.org/mongo-driver/mongo"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	client *mongo.Client
}

func NewClient(ctx Context, opts ...*options.ClientOptions) (client *Client, err error) {
	mongoOpts := sliceutil.MapIf(opts, func(idx int) (extracted *mongoOptions.ClientOptions, ok bool) {
		if opts[idx] == nil {
			return
		}
		opt := opts[idx].Get()
		extracted = &opt
		ok = true
		return
	})

	var c *mongo.Client
	c, err = mongo.Connect(ctx, mongoOpts...)
	if err != nil {
		return
	}

	client = &Client{
		client: c,
	}
	return
}

func (c *Client) Database(name string, opts ...*options.DatabaseOptions) (database *Database) {
	mongoOpts := sliceutil.MapIf(opts, func(idx int) (value *mongoOptions.DatabaseOptions, ok bool) {
		if opts[idx] == nil {
			return
		}
		mongoOpt := opts[idx].Get()
		value = &mongoOpt
		ok = true
		return
	})
	d := c.client.Database(name, mongoOpts...)
	database = &Database{
		database: d,
	}
	return
}
