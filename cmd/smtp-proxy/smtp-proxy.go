package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/dinhtrung/imap-smtp-proxy/pkg/smtpsrv"
	"github.com/dinhtrung/imap-smtp-proxy/pkg/util"
	"log"
	"os"
)

var configFile string

// main try to start an IMAP Proxy server with given configuration
func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	flag.StringVar(&configFile, "c", "configs/smtp-proxy.json", "Configuration for SMTP server")
	flag.Parse()

	configFileData, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal(fmt.Errorf("unable to read config file %s: %s", configFile, err))
	}

	// start SMTP server
	smtpSettings := &util.SMTPServerWO{}
	if err := json.Unmarshal(configFileData, smtpSettings); err != nil {
		log.Fatal(fmt.Errorf("unable to parse config file %s: %s", configFile, err))
	}
	if _, err := govalidator.ValidateStruct(smtpSettings); err != nil {
		log.Fatal(fmt.Errorf("unable to validate config file %s: %s", configFile, err))
	}

	smtpBackend := smtpsrv.NewSMTPProxyBackend(*smtpSettings.Addr)
	smtpsrv.StartSMTPServer(context.Background(), smtpBackend, smtpSettings)

}
