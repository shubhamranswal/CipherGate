package mfa

import "testing"

func TestGenerateKey(t *testing.T) {

	service := &Service{}
	key, err := service.GenerateKey("shubham")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if key.Secret() == "" {
		t.Fatal("expected secret")
	}
}

func TestVerifyNilSecret(t *testing.T) {
	service := &Service{}
	result := service.Validate("123456", nil)
	if result {
		t.Fatal("expected false")
	}
}

func TestVerifyInvalidCode(t *testing.T) {
	service := &Service{}
	secret := "JBSWY3DPEHPK3PXP"
	result := service.Validate("000000", &secret)
	if result {
		t.Fatal("expected false")
	}
}
