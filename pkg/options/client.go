package options

import "go.mongodb.org/mongo-driver/mongo/options"

type ClientOptions options.ClientOptions

func Client() *ClientOptions {
	client := ClientOptions(*options.Client())
	return &client
}

func (c *ClientOptions) ApplyURI(uri string) *ClientOptions {
	clientOptions := options.ClientOptions(*c)
	opts := ClientOptions(*(&clientOptions).ApplyURI(uri))
	return &opts
}
