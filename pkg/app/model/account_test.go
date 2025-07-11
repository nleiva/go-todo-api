package model_test

import (
	"testing"

	"github.com/nleiva/go-todo-api/pkg/app/model"
)

func TestAccountModelWriteRemote(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		account := model.Account{}
		account.New(model.Account{})

		if account.Email != "" {
			t.Errorf("Expected Email to be empty, got %s", account.Email)
		}
		if account.Password != "" {
			t.Errorf("Expected Password to be empty, got %s", account.Password)
		}
		if account.Firstname != "" {
			t.Errorf("Expected Firstname to be empty, got %s", account.Firstname)
		}
		if account.Lastname != "" {
			t.Errorf("Expected Lastname to be empty, got %s", account.Lastname)
		}
		if account.TokenSecret != "" {
			t.Errorf("Expected TokenSecret to be empty, got %s", account.TokenSecret)
		}
	})

	t.Run("with data", func(t *testing.T) {
		account := model.Account{}
		account.New(model.Account{
			Email:       "test@turbomeet.xyz",
			Password:    "password",
			Firstname:   "Firstname",
			Lastname:    "Lastname",
			TokenSecret: "token",
		})

		if account.Email != "test@turbomeet.xyz" {
			t.Errorf("Expected Email to be 'test@turbomeet.xyz', got %s", account.Email)
		}
		if account.Password != "" {
			t.Errorf("Expected Password to be empty (should not be written), got %s", account.Password)
		}
		if account.Firstname != "Firstname" {
			t.Errorf("Expected Firstname to be 'Firstname', got %s", account.Firstname)
		}
		if account.Lastname != "Lastname" {
			t.Errorf("Expected Lastname to be 'Lastname', got %s", account.Lastname)
		}
		if account.TokenSecret != "" {
			t.Errorf("Expected TokenSecret to be empty (should not be written), got %s", account.TokenSecret)
		}
	})
}
