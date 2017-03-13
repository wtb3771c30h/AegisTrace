// Package policy implements short-lived Ed25519 capability tokens.
package policy

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Claims struct {
	Issuer    string    `json:"iss"`
	Subject   string    `json:"sub"`
	Scopes    []string  `json:"scopes"`
	IssuedAt  time.Time `json:"iat"`
	ExpiresAt time.Time `json:"exp"`
	Nonce     string    `json:"nonce"`
