package imapsrv

import (
	"context"
	"github.com/dinhtrung/imap-smtp-proxy/pkg/util"
	"github.com/emersion/go-imap/backend"
	"github.com/emersion/go-imap/server"
	"log"
	"time"
)

// StartIMAPServer start the IMAP server
func StartIMAPServer(ctx context.Context, bkd backend.Backend, settings *util.IMAPServerWO) {
	s := server.New(bkd)

	if settings.Addr != nil {
		s.Addr = *settings.Addr
	}
	if settings.AutoLogout != nil {
		if autoLogout, err := time.ParseDuration(*settings.AutoLogout); err == nil {
			s.AutoLogout = autoLogout
		}
	}
	if settings.AllowInsecureAuth != nil {
		s.AllowInsecureAuth = *settings.AllowInsecureAuth
	}

	shouldExit := make(chan interface{})
	log.Printf("Starting IMAP server at %s", s.Addr)
	// await for IMAP server
	go func() {
		for {
			select {
			case <-shouldExit:
				return
			case <-ctx.Done():
				log.Printf("shutting down IMAP server...")
				if err := s.Close(); err != nil {
					log.Printf("unable to shutdown IMAP server due to %s", err)
				}
			}
		}
	}()
	if err := s.ListenAndServe(); err != nil {
		log.Printf("unable to start IMAP server due to :%s", err)
		shouldExit <- true
		close(shouldExit)
	}
}
