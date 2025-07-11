package handler_test

import (
	"net/http"
	"testing"

	"github.com/nleiva/go-todo-api/pkg/app/model"
	"github.com/nleiva/go-todo-api/pkg/app/service"
	"github.com/nleiva/go-todo-api/pkg/jwt"
	"github.com/nleiva/go-todo-api/test"
	"gopkg.in/guregu/null.v4/zero"
)

func TestTodosHandlerExportCSV(t *testing.T) {
	// Setup
	pw, _ := model.HashPassword("123456")
	account := &model.Account{
		Email:     "todos.exportcsv@turbomeet.xyz",
		Password:  pw,
		Firstname: "Accounts",
		Lastname:  "ExportCSV",
	}
	accountService := service.NewAccountService(DB)
	accountService.CreateAccount(account)

	auth, _ := jwt.Generate(account)
	authToken := auth.Token

	todoService := service.NewTodoService(DB)
	todoService.CreateTodo(&model.Todo{
		Title:     zero.NewString("Todo 1", true),
		Completed: false,
		AccountID: account.ID,
	})

	todoService.CreateTodo(&model.Todo{
		Title:       zero.NewString("Todo 2", true),
		Completed:   false,
		AccountID:   account.ID,
		Description: zero.NewString("some description", true),
	})

	todoService.CreateTodo(&model.Todo{
		Title:       zero.NewString("Todo 3", true),
		Completed:   true,
		AccountID:   account.ID,
		Description: zero.NewString("some longer description with a \n line-break", true),
	})

	todoService.CreateTodo(&model.Todo{
		Title:     zero.NewString("Todo 4", true),
		Completed: true,
		AccountID: account.ID,
	})

	t.Run("should be unauthorized", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/todos/csv", nil)
		res, _ := App.Test(req)

		if res.StatusCode != 401 {
			t.Errorf("Expected status code 401, got %d", res.StatusCode)
		}
	})

	t.Run("should be authorized", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/todos/csv", nil)
		req.Header.Set("Authorization", "Bearer "+authToken)
		res, _ := App.Test(req)

		if res.StatusCode != 200 {
			t.Errorf("Expected status code 200, got %d", res.StatusCode)
		}

		expectedContentType := "text/csv"
		actualContentType := res.Header.Get("Content-Type")
		if actualContentType != expectedContentType {
			t.Errorf("Expected Content-Type to be %s, got %s", expectedContentType, actualContentType)
		}

		// With the following code you could write the response to a file
		// fo, err := os.Create(config.TEST_FILE_PATH + "todos-export.csv")
		// if err != nil {
		// 	panic(err)
		// }

		// // close fo on exit and check for its returned error
		// defer func() {
		// 	if err := fo.Close(); err != nil {
		// 		panic(err)
		// 	}
		// }()

		// // make a buffer to keep chunks that are read
		// buf := make([]byte, 1024)
		// for {
		// 	// read a chunk
		// 	n, err := res.Body.Read(buf)
		// 	if err != nil && err != io.EOF {
		// 		panic(err)
		// 	}
		// 	if n == 0 {
		// 		break
		// 	}

		// 	// write a chunk
		// 	if _, err := fo.Write(buf[:n]); err != nil {
		// 		panic(err)
		// 	}
		// }
	})

	// Cleanup
	test.ClearTables(DB, []string{"accounts"})
	test.ClearAllTables(DB)
}
