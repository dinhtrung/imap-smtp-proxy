package util

import (
	"time"
)

// IMAPServerWO provide settings for IMAP server
type IMAPServerWO struct {
	Addr              *string        `json:"addr" valid:"required"`
	AllowInsecureAuth *bool          `json:"allowInsecureAuth"`
	AutoLogout        *time.Duration `json:"autoLogout"`
}
