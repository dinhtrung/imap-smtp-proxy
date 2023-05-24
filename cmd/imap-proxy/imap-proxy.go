package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/dinhtrung/imap-smtp-proxy/pkg/imapsrv"
	"github.com/dinhtrung/imap-smtp-proxy/pkg/util"
	"log"
	"os"
)

var configFile string

// main try to start an IMAP Proxy server with given configuration
func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	flag.StringVar(&configFile, "c", "configs/imap-proxy.json", "Configuration for IMAP server")
	flag.Parse()

	configFileData, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal(fmt.Errorf("unable to read config file %s: %s", configFile, err))
	}

	imapSettings := &util.IMAPServerWO{}
	if err := json.Unmarshal(configFileData, imapSettings); err != nil {
		log.Fatal(fmt.Errorf("unable to parse config file %s: %s", configFile, err))
	}

	if _, err := govalidator.ValidateStruct(imapSettings); err != nil {
		log.Fatal(fmt.Errorf("unable to validate config file %s: %s", configFile, err))
	}

	imapBackend := imapsrv.NewIMAPProxyBackend(*imapSettings.Addr)
	imapsrv.StartIMAPServer(context.Background(), imapBackend, imapSettings)

}
