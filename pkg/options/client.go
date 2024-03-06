package options

import (
	"context"
	"crypto/tls"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"net"
	"net/http"
	"time"
)

type ClientOptions struct {
	opts *options.ClientOptions
}

func Client() *ClientOptions {
	return &ClientOptions{
		opts: options.Client(),
	}
}

func (c *ClientOptions) Validate() error {
	return c.opts.Validate()
}

func (c *ClientOptions) Get() options.ClientOptions {
	return *c.opts
}

func (c *ClientOptions) GetURI() string {
	return c.opts.GetURI()
}

func (c *ClientOptions) ApplyURI(uri string) *ClientOptions {
	c.opts = c.opts.ApplyURI(uri)
	return c
}

func (c *ClientOptions) SetAppName(s string) *ClientOptions {
	c.opts = c.opts.SetAppName(s)
	return c
}

func (c *ClientOptions) SetAuth(auth Credential) *ClientOptions {
	c.opts = c.opts.SetAuth(auth.credential)
	return c
}

func (c *ClientOptions) SetCompressors(comps []string) *ClientOptions {
	c.opts = c.opts.SetCompressors(comps)
	return c
}

func (c *ClientOptions) SetConnectTimeout(d time.Duration) *ClientOptions {
	c.opts = c.opts.SetConnectTimeout(d)
	return c
}

func (c *ClientOptions) SetDialer(d ContextDialer) *ClientOptions {
	c.opts = c.opts.SetDialer(d)
	return c
}

func (c *ClientOptions) SetDirect(b bool) *ClientOptions {
	c.opts = c.opts.SetDirect(b)
	return c
}

func (c *ClientOptions) SetHeartbeatInterval(d time.Duration) *ClientOptions {
	c.opts = c.opts.SetHeartbeatInterval(d)
	return c
}

func (c *ClientOptions) SetHosts(s []string) *ClientOptions {
	c.opts = c.opts.SetHosts(s)
	return c
}

func (c *ClientOptions) SetLoadBalanced(lb bool) *ClientOptions {
	c.opts = c.opts.SetLoadBalanced(lb)
	return c
}

func (c *ClientOptions) SetLocalThreshold(d time.Duration) *ClientOptions {
	c.opts = c.opts.SetLocalThreshold(d)
	return c
}

func (c *ClientOptions) SetLoggerOptions(opts *options.LoggerOptions) *ClientOptions {
	c.opts = c.opts.SetLoggerOptions(opts)
	return c
}

func (c *ClientOptions) SetMaxConnIdleTime(d time.Duration) *ClientOptions {
	c.opts = c.opts.SetMaxConnIdleTime(d)
	return c
}

func (c *ClientOptions) SetMaxPoolSize(u uint64) *ClientOptions {
	c.opts = c.opts.SetMaxPoolSize(u)
	return c
}

func (c *ClientOptions) SetMinPoolSize(u uint64) *ClientOptions {
	c.opts = c.opts.SetMinPoolSize(u)
	return c
}

func (c *ClientOptions) SetMaxConnecting(u uint64) *ClientOptions {
	c.opts = c.opts.SetMaxConnecting(u)
	return c
}

func (c *ClientOptions) SetPoolMonitor(m *event.PoolMonitor) *ClientOptions {
	c.opts = c.opts.SetPoolMonitor(m)
	return c
}

func (c *ClientOptions) SetMonitor(m *event.CommandMonitor) *ClientOptions {
	c.opts = c.opts.SetMonitor(m)
	return c
}

func (c *ClientOptions) SetServerMonitor(m *event.ServerMonitor) *ClientOptions {
	c.opts = c.opts.SetServerMonitor(m)
	return c
}

func (c *ClientOptions) SetReadConcern(rc *readconcern.ReadConcern) *ClientOptions {
	c.opts = c.opts.SetReadConcern(rc)
	return c
}

func (c *ClientOptions) SetReadPreference(rp *readpref.ReadPref) *ClientOptions {
	c.opts = c.opts.SetReadPreference(rp)
	return c
}

func (c *ClientOptions) SetBSONOptions(opts *options.BSONOptions) *ClientOptions {
	c.opts = c.opts.SetBSONOptions(opts)
	return c
}

func (c *ClientOptions) SetRegistry(registry *bsoncodec.Registry) *ClientOptions {
	c.opts = c.opts.SetRegistry(registry)
	return c
}

func (c *ClientOptions) SetReplicaSet(s string) *ClientOptions {
	c.opts = c.opts.SetReplicaSet(s)
	return c
}

func (c *ClientOptions) SetRetryWrites(b bool) *ClientOptions {
	c.opts = c.opts.SetRetryWrites(b)
	return c
}

func (c *ClientOptions) SetRetryReads(b bool) *ClientOptions {
	c.opts = c.opts.SetRetryReads(b)
	return c
}

func (c *ClientOptions) SetServerSelectionTimeout(d time.Duration) *ClientOptions {
	c.opts = c.opts.SetServerSelectionTimeout(d)
	return c
}

func (c *ClientOptions) SetSocketTimeout(d time.Duration) *ClientOptions {
	c.opts = c.opts.SetSocketTimeout(d)
	return c
}

func (c *ClientOptions) SetTimeout(d time.Duration) *ClientOptions {
	c.opts = c.opts.SetTimeout(d)
	return c
}

func (c *ClientOptions) SetTLSConfig(cfg *tls.Config) *ClientOptions {
	c.opts = c.opts.SetTLSConfig(cfg)
	return c
}

func (c *ClientOptions) SetHTTPClient(client *http.Client) *ClientOptions {
	c.opts = c.opts.SetHTTPClient(client)
	return c
}

func (c *ClientOptions) SetWriteConcern(wc *writeconcern.WriteConcern) *ClientOptions {
	c.opts = c.opts.SetWriteConcern(wc)
	return c
}

func (c *ClientOptions) SetZlibLevel(level int) *ClientOptions {
	c.opts = c.opts.SetZlibLevel(level)
	return c
}

func (c *ClientOptions) SetZstdLevel(level int) *ClientOptions {
	c.opts = c.opts.SetZstdLevel(level)
	return c
}
func (c *ClientOptions) SetAutoEncryptionOptions(opts *options.AutoEncryptionOptions) *ClientOptions {
	c.opts = c.opts.SetAutoEncryptionOptions(opts)
	return c
}

func (c *ClientOptions) SetDisableOCSPEndpointCheck(disableCheck bool) *ClientOptions {
	c.opts = c.opts.SetDisableOCSPEndpointCheck(disableCheck)
	return c
}

func (c *ClientOptions) SetServerAPIOptions(opts *options.ServerAPIOptions) *ClientOptions {
	c.opts = c.opts.SetServerAPIOptions(opts)
	return c
}

func (c *ClientOptions) SetServerMonitoringMode(mode string) *ClientOptions {
	c.opts = c.opts.SetServerMonitoringMode(mode)
	return c
}

func (c *ClientOptions) SetSRVMaxHosts(srvMaxHosts int) *ClientOptions {
	c.opts = c.opts.SetSRVMaxHosts(srvMaxHosts)
	return c
}

func (c *ClientOptions) SetSRVServiceName(srvName string) *ClientOptions {
	c.opts = c.opts.SetSRVServiceName(srvName)
	return c
}

type Credential struct {
	credential options.Credential
}

type ContextDialer interface {
	DialContext(ctx context.Context, network, address string) (net.Conn, error)
}
