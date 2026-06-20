package mfa

import "github.com/pquerna/otp/totp"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GenerateSecret(
	username string,
) (string, error) {
	key, err := totp.Generate(
		totp.GenerateOpts{
			Issuer:      "CipherGate",
			AccountName: username,
		},
	)

	if err != nil {
		return "", err
	}

	return key.Secret(), nil
}

func (s *Service) Validate(
	code string,
	secret string,
) bool {
	return totp.Validate(
		code,
		secret,
	)
}

func (s *Service) Verify(
	code string,
	secret *string,
) bool {

	if secret == nil {
		return false
	}

	return totp.Validate(
		code,
		*secret,
	)
}
