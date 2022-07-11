package util

// SMTPServerWO provide settings for SMTPServerWOserver
type SMTPServerWO struct {
	Addr              *string `json:"addr" valid:"required"`
	Upstream          *string `json:"upstream" valid:"required"`
	AllowInsecureAuth *bool   `json:"allowInsecureAuth"`
	Domain            *string `json:"domain" valid:"required,host"`
	MaxMessageBytes   *uint   `json:"maxMessageBytes"`
	MaxRecipients     *uint   `json:"maxRecipients"`
	ReadTimeout       *string `json:"readTimeout"`
	WriteTimeout      *string `json:"writeTimeout"`
}
