package smtpsrv

import (
	"context"
	"github.com/dinhtrung/imap-smtp-proxy/pkg/util"
	"github.com/emersion/go-smtp"
	"log"
	"time"
)

// StartSMTPServer start the SMTP server
func StartSMTPServer(ctx context.Context, be smtp.Backend, cfg *util.SMTPServerWO) {
	s := smtp.NewServer(be)

	if cfg.Addr != nil {
		s.Addr = *cfg.Addr
	}
	if cfg.Domain != nil {
		s.Domain = *cfg.Domain
	}
	if cfg.ReadTimeout != nil {
		if readTimeout, err := time.ParseDuration(*cfg.ReadTimeout); err == nil {
			s.ReadTimeout = readTimeout
		}
	}
	if cfg.WriteTimeout != nil {
		if writeTimeout, err := time.ParseDuration(*cfg.WriteTimeout); err == nil {
			s.WriteTimeout = writeTimeout
		}
	}
	if cfg.MaxMessageBytes != nil {
		s.MaxMessageBytes = int(*cfg.MaxMessageBytes)
	}
	if cfg.MaxRecipients != nil {
		s.MaxRecipients = int(*cfg.MaxRecipients)
	}
	if cfg.AllowInsecureAuth != nil {
		s.AllowInsecureAuth = *cfg.AllowInsecureAuth
	}
	log.Printf("Starting SMTP server at %s", s.Addr)
	shouldExit := make(chan interface{})
	// await for IMAP server
	go func() {
		for {
			select {
			case <-shouldExit:
				return
			case <-ctx.Done():
				log.Printf("shutting down SMTP server...")
				if err := s.Close(); err != nil {
					log.Printf("unable to shutdown IMAP server due to %s", err)
				}
			}
		}
	}()
	if err := s.ListenAndServe(); err != nil {
		log.Printf("unable to start SMTP server : %s", err)
		shouldExit <- true
	}
}
