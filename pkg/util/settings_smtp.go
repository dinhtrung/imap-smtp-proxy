package util

import (
	"time"
)

// SMTPServerWO provide settings for SMTPServerWOserver
type SMTPServerWO struct {
	Addr              *string        `json:"addr" valid:"required"`
	AllowInsecureAuth *bool          `json:"allowInsecureAuth"`
	Domain            *string        `json:"domain" valid:"required,host"`
	MaxMessageBytes   *uint          `json:"maxMessageBytes"`
	MaxRecipients     *uint          `json:"maxRecipients"`
	ReadTimeout       *time.Duration `json:"readTimeout"`
	WriteTimeout      *time.Duration `json:"writeTimeout"`
}
