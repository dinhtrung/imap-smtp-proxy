package imapsrv

// see: https://github.com/emersion/go-imap-proxy/blob/master/backend.go

import (
	"crypto/tls"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/backend"
	"github.com/emersion/go-imap/client"
)

type Security int

const (
	SecurityNone Security = iota
	SecuritySTARTTLS
	SecurityTLS
)

type IMAPProxyBackend struct {
	Addr      string
	Security  Security
	TLSConfig *tls.Config

	unexported struct{}
}

func NewIMAPProxyBackend(addr string) *IMAPProxyBackend {
	return &IMAPProxyBackend{
		Addr:     addr,
		Security: SecuritySTARTTLS,
	}
}

func NewIMAPProxyBackendWithTLS(addr string, tlsConfig *tls.Config) *IMAPProxyBackend {
	return &IMAPProxyBackend{
		Addr:      addr,
		Security:  SecurityTLS,
		TLSConfig: tlsConfig,
	}
}

func (be *IMAPProxyBackend) login(username, password string) (*client.Client, error) {
	var c *client.Client
	var err error
	if be.Security == SecurityTLS {
		if c, err = client.DialTLS(be.Addr, be.TLSConfig); err != nil {
			return nil, err
		}
	} else {
		if c, err = client.Dial(be.Addr); err != nil {
			return nil, err
		}

		if be.Security == SecuritySTARTTLS {
			if err := c.StartTLS(be.TLSConfig); err != nil {
				return nil, err
			}
		}
	}

	if err := c.Login(username, password); err != nil {
		return nil, err
	}

	return c, nil
}

func (be *IMAPProxyBackend) Login(_ *imap.ConnInfo, username, password string) (backend.User, error) {
	c, err := be.login(username, password)
	if err != nil {
		return nil, err
	}

	u := &user{
		be:       be,
		c:        c,
		username: username,
	}
	return u, nil
}
