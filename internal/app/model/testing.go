package model

import "testing"

func TestUser(_ *testing.T) *User {
	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}
