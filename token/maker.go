package token

import (
	"time"
)

// Make is an interface to manage tokens
type Maker interface {
	//Create new token for specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	//Checks if token is valid
	VerifyToken(token string) (*Payload, error)
}
