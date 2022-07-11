# IMAP + SMTP Proxy Server

This application will bring up an IMAP and / or SMTP server, which connect and retrieve / send email on another server.

## Build

```shell
$ make
```

## Usage

1. Edit the `configs/imap-proxy.json` and `configs/smtp-proxy.json` accordingly
2. Run the binary with `-c` flag point to the `.json` file