package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the database.
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

// DBClient is the interface for the database client.
type DBClient interface {
	CreateUser(name, email, password, address string) error
}

// CreateUser is the handler for POST /users.
func CreateUser(dbClient DBClient) func(c *gin.Context) {
	return func(c *gin.Context) {
		user := new(User)
		if err := c.ShouldBindJSON(user); err != nil {
			log.Printf("failed to bind json: %v\n", err)

			c.JSON(http.StatusBadRequest, newErrorResponse(http.StatusBadRequest))
			return
		}

		if err := user.validate(); err != nil {
			log.Printf("failed to validate user: %v\n", err)

			c.JSON(http.StatusBadRequest, newErrorResponse(http.StatusBadRequest))
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("failed to hash password: %v\n", err)

			c.JSON(http.StatusBadRequest, newErrorResponse(http.StatusInternalServerError))
			return
		}

		if err := dbClient.CreateUser(user.Username, user.Email, string(hashedPassword), user.Address); err != nil {
			log.Fatalf("failed to create user: %v\n", err)

			c.JSON(http.StatusBadRequest, newErrorResponse(http.StatusInternalServerError))
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "pong",
		})
	}
}

// validate validates the user.
func (u *User) validate() error {
	if u.Username == "" {
		return fmt.Errorf("missing username")
	}
	if u.Email == "" {
		return fmt.Errorf("missing email")
	}
	if u.Password == "" {
		return fmt.Errorf("missing password")
	}
	if u.Address == "" {
		return fmt.Errorf("missing address")
	}

	return nil
}
