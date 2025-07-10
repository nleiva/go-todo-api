package model_test

import (
	"testing"
	"time"

	"github.com/nleiva/go-todo-api/pkg/app/model"
	"gopkg.in/guregu/null.v4"
	"gopkg.in/guregu/null.v4/zero"
)

func TestTodoModelWriteRemote(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		todo := model.Todo{}
		todo.New(model.Todo{})

		expectedTitle := zero.NewString("", false)
		if todo.Title != expectedTitle {
			t.Errorf("Expected Title to be %v, got %v", expectedTitle, todo.Title)
		}

		expectedDescription := zero.NewString("", false)
		if todo.Description != expectedDescription {
			t.Errorf("Expected Description to be %v, got %v", expectedDescription, todo.Description)
		}

		if todo.Completed != false {
			t.Errorf("Expected Completed to be false, got %v", todo.Completed)
		}

		expectedCompletedAt := null.NewTime(time.Time{}, false)
		if todo.CompletedAt != expectedCompletedAt {
			t.Errorf("Expected CompletedAt to be %v, got %v", expectedCompletedAt, todo.CompletedAt)
		}
	})

	t.Run("with data", func(t *testing.T) {
		testTime := time.Now()

		todo := model.Todo{}
		todo.New(model.Todo{
			Title:       zero.NewString("Title", true),
			Description: zero.NewString("Description", true),
			Completed:   true,
			CompletedAt: null.NewTime(testTime, true),
		})

		expectedTitle := zero.NewString("Title", true)
		if todo.Title != expectedTitle {
			t.Errorf("Expected Title to be %v, got %v", expectedTitle, todo.Title)
		}

		expectedDescription := zero.NewString("Description", true)
		if todo.Description != expectedDescription {
			t.Errorf("Expected Description to be %v, got %v", expectedDescription, todo.Description)
		}

		if todo.Completed != true {
			t.Errorf("Expected Completed to be true, got %v", todo.Completed)
		}

		expectedCompletedAt := null.NewTime(testTime, true)
		if todo.CompletedAt != expectedCompletedAt {
			t.Errorf("Expected CompletedAt to be %v, got %v", expectedCompletedAt, todo.CompletedAt)
		}
	})
}
