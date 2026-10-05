package auth

import (
	"context"
	"crypto/subtle"

	"github.com/tapiaw38/alutec-inventory-be/internal/platform/config"
	apperrors "github.com/tapiaw38/alutec-inventory-be/internal/platform/errors"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/token"
)

type (
	// LoginUsecase checks the single credential pair the service is configured
	// with. It is deliberately the only place that knows where credentials come
	// from, so swapping in real user accounts touches nothing else.
	LoginUsecase interface {
		Execute(ctx context.Context, in LoginInput) (*LoginOutput, apperrors.ApplicationError)
	}

	loginUsecase struct{}

	LoginInput struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	LoginOutput struct {
		Data SessionData `json:"data"`
	}

	SessionData struct {
		Token string `json:"token"`
		Email string `json:"email"`
	}
)

func NewLoginUsecase() LoginUsecase {
	return &loginUsecase{}
}

func (u *loginUsecase) Execute(_ context.Context, in LoginInput) (*LoginOutput, apperrors.ApplicationError) {
	cfg := config.GetConfigService().AuthConfig

	// Without configured credentials every password would compare equal to the
	// empty string, so refuse to authenticate at all.
	if cfg.Email == "" || cfg.Password == "" || cfg.Secret == "" {
		return nil, apperrors.NewInternalError(errAuthNotConfigured)
	}

	// Compare both fields unconditionally: bailing out early on a wrong email
	// would leak which half was wrong through timing.
	emailOK := subtle.ConstantTimeCompare([]byte(in.Email), []byte(cfg.Email)) == 1
	passwordOK := subtle.ConstantTimeCompare([]byte(in.Password), []byte(cfg.Password)) == 1
	if !emailOK || !passwordOK {
		return nil, apperrors.NewUnauthorizedError("email o contraseña incorrectos")
	}

	signed, err := token.Issue(cfg.Email, cfg.Secret, cfg.TokenTTL)
	if err != nil {
		return nil, apperrors.NewInternalError(err)
	}

	return &LoginOutput{Data: SessionData{Token: signed, Email: cfg.Email}}, nil
}
