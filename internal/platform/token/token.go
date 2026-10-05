// Package token issues and verifies the signed session tokens the API accepts.
//
// The format is a trimmed-down JWT: "<base64url payload>.<base64url HMAC>".
// A full JWT library buys nothing here because there is a single issuer, a
// single audience and one signing algorithm.
package token

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

var (
	ErrMalformed = errors.New("malformed token")
	ErrSignature = errors.New("invalid token signature")
	ErrExpired   = errors.New("token expired")
)

type Claims struct {
	Subject   string `json:"sub"`
	ExpiresAt int64  `json:"exp"`
}

func sign(payload []byte, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

// Issue returns a token for subject that stops being valid after ttl.
func Issue(subject, secret string, ttl time.Duration) (string, error) {
	payload, err := json.Marshal(Claims{
		Subject:   subject,
		ExpiresAt: time.Now().Add(ttl).Unix(),
	})
	if err != nil {
		return "", err
	}

	encoded := base64.RawURLEncoding.EncodeToString(payload)
	return encoded + "." + sign(payload, secret), nil
}

// Verify checks the signature before the expiry, so a tampered token is
// rejected as forged rather than reported as merely expired.
func Verify(raw, secret string) (*Claims, error) {
	encoded, signature, found := strings.Cut(raw, ".")
	if !found {
		return nil, ErrMalformed
	}

	payload, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		return nil, ErrMalformed
	}

	if !hmac.Equal([]byte(signature), []byte(sign(payload, secret))) {
		return nil, ErrSignature
	}

	var claims Claims
	if err := json.Unmarshal(payload, &claims); err != nil {
		return nil, ErrMalformed
	}
	if time.Now().Unix() >= claims.ExpiresAt {
		return nil, ErrExpired
	}

	return &claims, nil
}
