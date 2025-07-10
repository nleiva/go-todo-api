package handler_test

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/nleiva/go-todo-api/pkg/app/model"
	"github.com/nleiva/go-todo-api/pkg/app/service"
	"github.com/nleiva/go-todo-api/pkg/app/types"
	"github.com/nleiva/go-todo-api/pkg/jwt"
	"github.com/nleiva/go-todo-api/pkg/permission"
	"github.com/nleiva/go-todo-api/test"
)

func TestAccountsHandlerList(t *testing.T) {
	// Setup
	pw, _ := model.HashPassword("123456")
	account := &model.Account{
		Email:      "accounts.list@turbomeet.xyz",
		Password:   pw,
		Firstname:  "Accounts",
		Lastname:   "List",
		Permission: permission.ACCOUNTS_READ_ALL,
	}
	accountService := service.NewAccountService(DB)
	accountService.CreateAccount(account)

	auth, _ := jwt.Generate(account)
	authToken := auth.Token

	t.Run("should be unauthorized", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/accounts", nil)
		res, _ := App.Test(req)

		if res.StatusCode != 401 {
			t.Errorf("Expected status code 401, got %d", res.StatusCode)
		}
	})

	t.Run("should be authorized", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/accounts", nil)
		req.Header.Set("Authorization", "Bearer "+authToken)
		res, _ := App.Test(req)

		if res.StatusCode != 200 {
			t.Errorf("Expected status code 200, got %d", res.StatusCode)
		}

		bodyBytes, _ := io.ReadAll(res.Body)
		result := types.GetAccountsResponse{}
		if err := json.Unmarshal(bodyBytes, &result); err != nil {
			t.Fatalf("Error unmarshalling response: %v", err)
		}

		// Check if the account we created is in the response
		found := false
		for _, acc := range result.Accounts {
			if acc.Email == "accounts.list@turbomeet.xyz" &&
				acc.Password == "" && // Password should be empty in response
				acc.Firstname == "Accounts" &&
				acc.Lastname == "List" {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Expected account not found in response")
		}
	})

	// Cleanup
	test.ClearTables(DB, []string{"accounts"})
	test.ClearAllTables(DB)
}
