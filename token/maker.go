package token

import "time"

// Maker is an interface for managing token, we can choose which technology (JWT, PASETO) to use
type Maker interface {
	// CreateToken creates a new token for specific username and duration
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	// VerifyToken check token is valid or not
	VerifyToken(token string) (*Payload, error)
}
