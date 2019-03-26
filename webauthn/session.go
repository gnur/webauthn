package webauthn

import (
	"github.com/gnur/webauthn/protocol"
)

type SessionData struct {
	Challenge            protocol.Challenge `json:"challenge"`
	UserID               []byte             `json:"user_id"`
	AllowedCredentialIDs [][]byte           `json:"allowed_credentials,omitempty"`
}
