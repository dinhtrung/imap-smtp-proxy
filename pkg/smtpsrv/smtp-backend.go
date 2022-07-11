package smtpsrv

// see: https://github.com/emersion/go-smtp-proxy/blob/master/backend.go

import (
	"crypto/tls"
	"errors"
	"net"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

type Security int

const (
	SecurityTLS Security = iota
	SecurityStartTLS
	SecurityNone
)

type SMTPProxyBackend struct {
	Addr      string
	Security  Security
	TLSConfig *tls.Config
	LMTP      bool
	Host      string
	LocalName string

	unexported struct{}
}

func NewSMTPProxyBackend(addr string) *SMTPProxyBackend {
	return &SMTPProxyBackend{Addr: addr, Security: SecurityStartTLS}
}

func NewSMTPProxyBackendWithTLS(addr string, tlsConfig *tls.Config) *SMTPProxyBackend {
	return &SMTPProxyBackend{
		Addr:      addr,
		Security:  SecurityTLS,
		TLSConfig: tlsConfig,
	}
}

func NewLMTP(addr string, host string) *SMTPProxyBackend {
	return &SMTPProxyBackend{
		Addr:     addr,
		Security: SecurityNone,
		LMTP:     true,
		Host:     host,
	}
}

func (be *SMTPProxyBackend) newConn() (*smtp.Client, error) {
	var conn net.Conn
	var err error
	if be.LMTP {
		if be.Security != SecurityNone {
			return nil, errors.New("smtp-proxy: LMTP doesn't support TLS")
		}
		conn, err = net.Dial("unix", be.Addr)
	} else if be.Security == SecurityTLS {
		conn, err = tls.Dial("tcp", be.Addr, be.TLSConfig)
	} else {
		conn, err = net.Dial("tcp", be.Addr)
	}
	if err != nil {
		return nil, err
	}

	var c *smtp.Client
	if be.LMTP {
		c, err = smtp.NewClientLMTP(conn, be.Host)
	} else {
		host := be.Host
		if host == "" {
			host, _, _ = net.SplitHostPort(be.Addr)
		}
		c, err = smtp.NewClient(conn, host)
	}
	if err != nil {
		return nil, err
	}

	if be.LocalName != "" {
		err = c.Hello(be.LocalName)
		if err != nil {
			return nil, err
		}
	}

	if be.Security == SecurityStartTLS {
		if err := c.StartTLS(be.TLSConfig); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (be *SMTPProxyBackend) login(username, password string) (*smtp.Client, error) {
	c, err := be.newConn()
	if err != nil {
		return nil, err
	}

	auth := sasl.NewPlainClient("", username, password)
	if err := c.Auth(auth); err != nil {
		return nil, err
	}

	return c, nil
}

func (be *SMTPProxyBackend) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	c, err := be.login(username, password)
	if err != nil {
		return nil, err
	}

	s := &session{
		c:  c,
		be: be,
	}
	return s, nil
}

func (be *SMTPProxyBackend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	c, err := be.newConn()
	if err != nil {
		return nil, err
	}

	s := &session{
		c:  c,
		be: be,
	}
	return s, nil
}
