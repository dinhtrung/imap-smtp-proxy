package smtpsrv

import (
	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	"io"
	"log"
)

type session struct {
	c  *smtp.Client
	be *SMTPProxyBackend
}

func (s *session) AuthPlain(username, password string) error {
	return s.c.Auth(sasl.NewPlainClient(username, username, password))
}

func (s *session) Reset() {
	if err := s.c.Reset(); err != nil {
		log.Printf("unable to reset session: %v", err)
	}
}

func (s *session) Mail(from string, opts *smtp.MailOptions) error {
	return s.c.Mail(from, opts)
}

func (s *session) Rcpt(to string) error {
	return s.c.Rcpt(to)
}

func (s *session) Data(r io.Reader) error {
	wc, err := s.c.Data()
	if err != nil {
		return err
	}

	_, err = io.Copy(wc, r)
	if err != nil {
		if err := wc.Close(); err != nil {
			log.Printf("unable to write data: %v", err)
		}
		return err
	}

	return wc.Close()
}

func (s *session) Logout() error {
	return s.c.Quit()
}
