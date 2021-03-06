package smtpsrv

import (
	"github.com/emersion/go-smtp"
	"io"
)

type session struct {
	c  *smtp.Client
	be *SMTPProxyBackend
}

func (s *session) Reset() {
	s.c.Reset()
}

func (s *session) Mail(from string, opts smtp.MailOptions) error {
	return s.c.Mail(from, &opts)
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
		wc.Close()
		return err
	}

	return wc.Close()
}

func (s *session) Logout() error {
	return s.c.Quit()
}
