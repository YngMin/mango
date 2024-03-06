package options

import (
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type DatabaseOptions struct {
	opts *options.DatabaseOptions
}

func Database() *DatabaseOptions {
	return &DatabaseOptions{
		opts: options.Database(),
	}
}

func (d *DatabaseOptions) Get() options.DatabaseOptions {
	return *d.opts
}

func (d *DatabaseOptions) SetReadConcern(rc *readconcern.ReadConcern) *DatabaseOptions {
	d.opts = d.opts.SetReadConcern(rc)
	return d
}

func (d *DatabaseOptions) SetWriteConcern(wc *writeconcern.WriteConcern) *DatabaseOptions {
	d.opts = d.opts.SetWriteConcern(wc)
	return d
}

func (d *DatabaseOptions) SetReadPreference(rp *readpref.ReadPref) *DatabaseOptions {
	d.opts = d.opts.SetReadPreference(rp)
	return d
}

func (d *DatabaseOptions) SetBSONOptions(opts *options.BSONOptions) *DatabaseOptions {
	d.opts = d.opts.SetBSONOptions(opts)
	return d
}

func (d *DatabaseOptions) SetRegistry(r *bsoncodec.Registry) *DatabaseOptions {
	d.opts = d.opts.SetRegistry(r)
	return d
}
