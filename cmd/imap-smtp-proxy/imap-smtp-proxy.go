package main

//
//// main start both the SMTP and IMAP server at the same time and add a work group for monitoring them
//func main()  {
//	imapSettings, err := util.GetIMAPServerSettings(apiCfgData)
//	if err != nil {
//		log.Fatal(err)
//	}
//	imapBackend := imapsrv.NewIMAPXOAUTH2Backend(baseURL, db, upstreamData)
//	go imapsrv.StartIMAPServer(ctx, imapBackend, imapSettings)
//
//
//	// start SMTP server
//	smtpSettings, err := util.GetSMTPServerSettings(apiCfgData)
//	if err != nil {
//		log.Fatal(err)
//	}
//	smptBackend := smtpsrv.NewSMPTXOAuth2Backend(baseURL, db, upstreamData)
//	go smtpsrv.StartSMTPServer(ctx, smptBackend, smtpSettings)
//
//}
