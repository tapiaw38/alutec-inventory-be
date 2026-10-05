package token

import (
	"errors"
	"strings"
	"testing"
	"time"
)

const secret = "test-secret"

func TestVerifyAcceptsFreshToken(t *testing.T) {
	raw, err := Issue("admin@example.com", secret, time.Hour)
	if err != nil {
		t.Fatalf("issue: %v", err)
	}

	claims, err := Verify(raw, secret)
	if err != nil {
		t.Fatalf("verify: %v", err)
	}
	if claims.Subject != "admin@example.com" {
		t.Fatalf("subject = %q, want admin@example.com", claims.Subject)
	}
}

func TestVerifyRejectsOtherSecret(t *testing.T) {
	raw, _ := Issue("admin@example.com", secret, time.Hour)

	if _, err := Verify(raw, "another-secret"); !errors.Is(err, ErrSignature) {
		t.Fatalf("err = %v, want ErrSignature", err)
	}
}

func TestVerifyRejectsTamperedPayload(t *testing.T) {
	raw, _ := Issue("viewer@example.com", secret, time.Hour)
	forged, _ := Issue("admin@example.com", "attacker-secret", time.Hour)

	// Swap in a payload the attacker controls, keeping the genuine signature.
	_, signature, _ := strings.Cut(raw, ".")
	payload, _, _ := strings.Cut(forged, ".")

	if _, err := Verify(payload+"."+signature, secret); !errors.Is(err, ErrSignature) {
		t.Fatalf("err = %v, want ErrSignature", err)
	}
}

func TestVerifyRejectsExpiredToken(t *testing.T) {
	raw, _ := Issue("admin@example.com", secret, -time.Second)

	if _, err := Verify(raw, secret); !errors.Is(err, ErrExpired) {
		t.Fatalf("err = %v, want ErrExpired", err)
	}
}

func TestVerifyRejectsGarbage(t *testing.T) {
	for _, raw := range []string{"", "no-separator", "!!!.???"} {
		if _, err := Verify(raw, secret); err == nil {
			t.Fatalf("Verify(%q) succeeded, want error", raw)
		}
	}
}
