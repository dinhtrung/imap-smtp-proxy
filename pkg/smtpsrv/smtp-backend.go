package smtpsrv

// see: https://github.com/emersion/go-smtp-proxy/blob/master/backend.go

import (
	"crypto/tls"
	"net"

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
}

func NewSMTPProxyBackend(addr string) *SMTPProxyBackend {
	return &SMTPProxyBackend{Addr: addr, Security: SecurityStartTLS}
}

// newSMTPClient try to construct the SMTP client based on the configuration
func (be *SMTPProxyBackend) newSMTPClient(conn net.Conn) (*smtp.Client, error) {
	var smtpClient *smtp.Client
	var err error
	if be.LMTP {
		smtpClient, err = smtp.NewClientLMTP(conn, be.Host)
	} else {
		host := be.Host
		if host == "" {
			host, _, _ = net.SplitHostPort(be.Addr)
		}
		smtpClient, err = smtp.NewClient(conn, host)
	}
	if err != nil {
		return nil, err
	}

	if be.LocalName != "" {
		err = smtpClient.Hello(be.LocalName)
		if err != nil {
			return nil, err
		}
	}

	if be.Security == SecurityStartTLS {
		if err := smtpClient.StartTLS(be.TLSConfig); err != nil {
			return nil, err
		}
	}
	return smtpClient, nil
}

// NewSession implements smpt.Session interface
func (be *SMTPProxyBackend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	smtpClient, err := be.newSMTPClient(c.Conn())
	if err != nil {
		return nil, err
	}
	s := &session{
		c:  smtpClient,
		be: be,
	}
	return s, nil
}
