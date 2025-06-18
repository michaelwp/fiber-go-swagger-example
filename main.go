package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"

	_ "fiber-go-swagger/docs" // Import generated docs
)

// @title Fiber Swagger API
// @version 1.0
// @description This is a sample API using Fiber and Swagger
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@swagger.io
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:3000
// @BasePath /api/v1
// @schemes http https
func main() {
	app := fiber.New()

	// Enable CORS
	app.Use(cors.New())

	// Swagger route
	app.Get("/swagger/*", swagger.HandlerDefault)

	// API routes
	api := app.Group("/api/v1")

	// User routes
	api.Get("/users", getUsers)
	api.Get("/users/:id", getUserByID)
	api.Post("/users", createUser)
	api.Put("/users/:id", updateUser)
	api.Delete("/users/:id", deleteUser)

	log.Fatal(app.Listen(":3000"))
}

// User represents a user in the system
type User struct {
	ID    int    `json:"id" example:"1"`
	Name  string `json:"name" example:"John Doe"`
	Email string `json:"email" example:"john@example.com"`
	Age   int    `json:"age" example:"30"`
}

// CreateUserRequest represents the request body for creating a user
type CreateUserRequest struct {
	Name  string `json:"name" example:"John Doe" validate:"required"`
	Email string `json:"email" example:"john@example.com" validate:"required,email"`
	Age   int    `json:"age" example:"30" validate:"required,min=1"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error" example:"Bad Request"`
	Message string `json:"message" example:"Invalid input data"`
}

// SuccessResponse represents a success response
type SuccessResponse struct {
	Message string      `json:"message" example:"Operation successful"`
	Data    interface{} `json:"data,omitempty"`
}

// getUsers godoc
// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {array} User
// @Failure 500 {object} ErrorResponse
// @Router /users [get]
func getUsers(c *fiber.Ctx) error {
	// Mock data
	users := []User{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Age: 30},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Age: 25},
	}

	return c.JSON(users)
}

// getUserByID godoc
// @Summary Get user by ID
// @Description Get a single user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [get]
func getUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	// Mock data - in real app, fetch from database
	user := User{
		ID:    1,
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   30,
	}

	if id == "1" {
		return c.JSON(user)
	}

	return c.Status(404).JSON(ErrorResponse{
		Error:   "Not Found",
		Message: "User not found",
	})
}

// createUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User data"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 422 {object} ErrorResponse
// @Router /users [post]
func createUser(c *fiber.Ctx) error {
	var req CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(ErrorResponse{
			Error:   "Bad Request",
			Message: "Invalid JSON format",
		})
	}

	// In a real app, save to database and return the created user
	response := SuccessResponse{
		Message: "User created successfully",
		Data: User{
			ID:    3,
			Name:  req.Name,
			Email: req.Email,
			Age:   req.Age,
		},
	}

	return c.Status(201).JSON(response)
}

// updateUser godoc
// @Summary Update an existing user
// @Description Update user information by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body CreateUserRequest true "Updated user data"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [put]
func updateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	log.Println("update user id:", id)

	var req CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(ErrorResponse{
			Error:   "Bad Request",
			Message: "Invalid JSON format",
		})
	}

	// Mock update - in real app, update in database
	response := SuccessResponse{
		Message: "User updated successfully",
		Data: User{
			ID:    1, // Would be parsed from id param
			Name:  req.Name,
			Email: req.Email,
			Age:   req.Age,
		},
	}

	return c.JSON(response)
}

// deleteUser godoc
// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [delete]
func deleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	log.Println("delete user id:", id)

	// Mock deletion - in real app, delete from database
	response := SuccessResponse{
		Message: "User deleted successfully",
	}

	return c.JSON(response)
}
