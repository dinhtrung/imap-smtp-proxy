package util

// IMAPServerWO provide settings for IMAP server
type IMAPServerWO struct {
	Addr              *string `json:"addr" valid:"required"`
	Upstream          *string `json:"upstream" valid:"required"`
	AllowInsecureAuth *bool   `json:"allowInsecureAuth"`
	AutoLogout        *string `json:"autoLogout"`
}
