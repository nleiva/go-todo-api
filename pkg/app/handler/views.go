package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/nleiva/go-todo-api/pkg/app/model"
	"github.com/nleiva/go-todo-api/pkg/app/types"
	"github.com/nleiva/go-todo-api/pkg/app/types/pagination"
	"github.com/nleiva/go-todo-api/pkg/jwt"
	"github.com/nleiva/go-todo-api/pkg/middleware/locals"
	"github.com/nleiva/go-todo-api/pkg/view"
	"github.com/nleiva/go-todo-api/utils"
	"gorm.io/gorm"
)

func (h *Handler) GetBaseData(c *fiber.Ctx) view.BaseData {
	account := &model.Account{}
	if locals.JwtPayload(c).Valid {
		h.accountService.FindAccountByID(account, locals.JwtPayload(c).AccountID)
	}

	return view.BaseData{
		IsAuthenticated: locals.JwtPayload(c).Valid,
		Account:         account,
	}
}

func (h *Handler) VIndex(c *fiber.Ctx) error {
	return adaptor.HTTPHandler(templ.Handler(view.IndexPage(h.GetBaseData(c))))(c)
}

func (h *Handler) VProfile(c *fiber.Ctx) error {
	// Get the account
	account := &model.Account{}
	accountID := locals.JwtPayload(c).AccountID

	err := h.accountService.FindAccountByID(account, accountID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &utils.NOT_FOUND
		}
		return &utils.INTERNAL_SERVER_ERROR
	}

	// Calculate todo statistics
	var totalTodos, completedTodos, pendingTodos int64

	// Get total todos count
	if err := h.db.Model(&model.Todo{}).Where("account_id = ?", accountID).Count(&totalTodos).Error; err != nil {
		return &utils.INTERNAL_SERVER_ERROR
	}

	// Get completed todos count
	if err := h.db.Model(&model.Todo{}).Where("account_id = ? AND completed = ?", accountID, true).Count(&completedTodos).Error; err != nil {
		return &utils.INTERNAL_SERVER_ERROR
	}

	// Calculate pending todos
	pendingTodos = totalTodos - completedTodos

	// Calculate completion rate
	completionRate := 0
	if totalTodos > 0 {
		completionRate = int((completedTodos * 100) / totalTodos)
	}

	// Prepare profile data
	profileData := types.ProfileData{
		Account: account,
		Stats: types.ProfileStats{
			TotalTodos:     totalTodos,
			CompletedTodos: completedTodos,
			PendingTodos:   pendingTodos,
			CompletionRate: completionRate,
		},
	}

	pageData := view.ProfilePageData{
		BaseData:    h.GetBaseData(c),
		ProfileData: profileData,
	}

	return adaptor.HTTPHandler(templ.Handler(view.ProfilePage(pageData)))(c)
}

func (h *Handler) VTodosIndex(c *fiber.Ctx) error {
	var meta = locals.Meta(c)

	meta.Order = append(meta.Order, pagination.OrderEntry{
		Key:       "created_at",
		Direction: "desc",
	})

	var todos = []model.Todo{}
	if err := h.FindWithMeta(&todos, &model.Todo{}, meta, h.db.Where("account_id = ?", locals.JwtPayload(c).AccountID)).Error; err != nil {
		return &utils.INTERNAL_SERVER_ERROR
	}

	return adaptor.HTTPHandler(templ.Handler(view.TodosIndexPage(h.GetBaseData(c), todos)))(c)
}

func (h *Handler) VTodosCreate(c *fiber.Ctx) error {
	todo := &model.Todo{}

	if err := ParseBodyAndValidate(c, todo, *h.validator); err != nil {
		return err
	}

	todo.AccountID = locals.JwtPayload(c).AccountID

	if err := h.todoService.CreateTodo(todo).Error; err != nil {
		return &utils.INTERNAL_SERVER_ERROR
	}

	return adaptor.HTTPHandler(templ.Handler(view.TodoItem(*todo)))(c)
}

func (h *Handler) VTodosComplete(c *fiber.Ctx) error {
	id := c.Params("id")

	todo := &model.Todo{}
	if err := h.todoService.FindTodoByID(todo, id, locals.JwtPayload(c).AccountID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &utils.NOT_FOUND
		}
		return &utils.INTERNAL_SERVER_ERROR
	}

	todo.Completed = !todo.Completed

	if err := h.todoService.UpdateTodo(todo).Error; err != nil {
		return &utils.INTERNAL_SERVER_ERROR
	}

	return adaptor.HTTPHandler(templ.Handler(view.TodoCompleteToggle(*todo)))(c)
}

func (h *Handler) VTodosDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	h.todoService.DeleteTodoByID(id)

	return c.Status(http.StatusOK).SendString("")
}

func (h *Handler) VLogin(c *fiber.Ctx) error {
	return adaptor.HTTPHandler(templ.Handler(view.LoginPage(h.GetBaseData(c))))(c)
}

func (h *Handler) VLoginPost(c *fiber.Ctx) error {
	remoteData := &types.LoginDTOBody{}

	if err := ParseBodyAndValidate(c, remoteData, *h.validator); err != nil {
		// Return login page with validation error
		baseData := h.GetBaseData(c)
		return adaptor.HTTPHandler(templ.Handler(view.LoginPageWithError(baseData, "Please check your input and try again.")))(c)
	}

	account := &model.Account{}
	if err := h.accountService.FindAccountByEmail(account, remoteData.Email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Return login page with "account not found" error
			baseData := h.GetBaseData(c)
			return adaptor.HTTPHandler(templ.Handler(view.LoginPageWithError(baseData, "No account found with this email address.")))(c)
		}
		// Return login page with generic error
		baseData := h.GetBaseData(c)
		return adaptor.HTTPHandler(templ.Handler(view.LoginPageWithError(baseData, "An error occurred. Please try again.")))(c)
	}

	if !model.CheckPasswordHash(remoteData.Password, account.Password) {
		// Return login page with "wrong password" error
		baseData := h.GetBaseData(c)
		return adaptor.HTTPHandler(templ.Handler(view.LoginPageWithError(baseData, "Incorrect email or password. Please try again.")))(c)
	}

	auth, err := jwt.Generate(account)
	if err != nil {
		// Return login page with generic error
		baseData := h.GetBaseData(c)
		return adaptor.HTTPHandler(templ.Handler(view.LoginPageWithError(baseData, "An error occurred. Please try again.")))(c)
	}

	//Set cookie and return 200
	c.Cookie(&fiber.Cookie{
		Name:     "go-todo-api_auth",
		Value:    auth.Token,
		HTTPOnly: true,
		Secure:   c.Protocol() == "https", // Only secure in HTTPS
		SameSite: "Lax",
		MaxAge:   3600, // 1 hour
	})

	c.Cookie(&fiber.Cookie{
		Name:     "go-todo-api_refresh",
		Value:    auth.RefreshToken,
		HTTPOnly: true,
		Secure:   c.Protocol() == "https",
		SameSite: "Lax",
		MaxAge:   86400, // 24 hours
	})

	c.Response().Header.Set("HX-Redirect", "/")

	return c.Status(http.StatusOK).SendString("")
}

func (h *Handler) VLogout(c *fiber.Ctx) error {
	// Clear auth cookies with proper settings
	c.Cookie(&fiber.Cookie{
		Name:     "go-todo-api_auth",
		Value:    "",
		HTTPOnly: true,
		Secure:   c.Protocol() == "https",
		SameSite: "Lax",
		MaxAge:   -1, // Expire immediately
		Expires:  time.Unix(0, 0),
	})

	c.Cookie(&fiber.Cookie{
		Name:     "go-todo-api_refresh",
		Value:    "",
		HTTPOnly: true,
		Secure:   c.Protocol() == "https",
		SameSite: "Lax",
		MaxAge:   -1, // Expire immediately
		Expires:  time.Unix(0, 0),
	})

	return c.Redirect("/")
}

func (h *Handler) VRegister(c *fiber.Ctx) error {
	return adaptor.HTTPHandler(templ.Handler(view.RegisterPage(h.GetBaseData(c))))(c)
}

func (h *Handler) VRegisterPost(c *fiber.Ctx) error {
	remoteData := &types.RegisterDTOBody{}

	if err := ParseBodyAndValidate(c, remoteData, *h.validator); err != nil {
		return err
	}

	// Check if passwords match
	if remoteData.Password != remoteData.ConfirmPassword {
		return utils.RequestErrorFrom(&utils.VALIDATION_ERROR, "Password and confirm password do not match")
	}

	// Check if user already exists
	existingAccount := &model.Account{}
	if err := h.accountService.FindAccountByEmail(existingAccount, remoteData.Email).Error; err == nil {
		return &utils.ACCOUNT_WITH_EMAIL_ALREADY_EXISTS
	}

	// Hash the password
	hashedPassword, err := model.HashPassword(remoteData.Password)
	if err != nil {
		return err
	}

	// Create new account
	account := &model.Account{
		Email:     remoteData.Email,
		Password:  hashedPassword,
		Firstname: remoteData.Firstname,
		Lastname:  remoteData.Lastname,
	}

	if err := h.accountService.CreateAccount(account).Error; err != nil {
		return err
	}

	// Generate JWT token for automatic login
	auth, err := jwt.Generate(account)
	if err != nil {
		return err
	}

	// Set cookies
	c.Cookie(&fiber.Cookie{
		Name:     "go-todo-api_auth",
		Value:    auth.Token,
		HTTPOnly: true,
	})

	c.Cookie(&fiber.Cookie{
		Name:     "go-todo-api_refresh",
		Value:    auth.RefreshToken,
		HTTPOnly: true,
	})

	// Redirect to home page
	c.Response().Header.Set("HX-Redirect", "/")

	return c.Status(http.StatusOK).SendString("")
}
