package user

import "testing"

func TestValidateUsernameTooShort(t *testing.T) {
	service := &Service{}
	err := service.ValidateUsername("ab")
	if err == nil {
		t.Fatal("expected validation error")
	}
}

func TestValidateUsernameValid(t *testing.T) {
	service := &Service{}
	err := service.ValidateUsername("shubham")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidatePasswordTooShort(t *testing.T) {
	service := &Service{}
	err := service.ValidatePassword("123")
	if err == nil {
		t.Fatal("expected validation error")
	}
}

func TestValidatePasswordValid(t *testing.T) {
	service := &Service{}
	err := service.ValidatePassword("Password@123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
