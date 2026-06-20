package mfa

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GenerateKey(username string) (*otp.Key, error) {
	return totp.Generate(
		totp.GenerateOpts{
			Issuer:      "CipherGate CLI",
			AccountName: username,
		},
	)
}

func (s *Service) Validate(code string, secret *string) bool {
	if secret == nil {
		return false
	}

	return totp.Validate(code, *secret)
}
