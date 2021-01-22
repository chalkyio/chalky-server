package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

type loginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

var signingKey []byte

func createAuthMiddleware() func(*fiber.Ctx) error {
	signingKey = []byte(os.Getenv("JWT_SIGNING_KEY"))
	return jwtware.New(jwtware.Config{
		SigningKey: signingKey,
	})
}

func handleLogin(c *fiber.Ctx) error {
	var data loginData
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Retrieve the password hash from the DB.
	var userID string
	var hash []byte
	if err := db.QueryRow(infiniteContext, "SELECT id, password_hash FROM users WHERE username = ?", userID, data.Username).Scan(&hash); err != nil {
		if err == pgx.ErrNoRows {
			// The user doesn't exist.
			resp, _ := json.Marshal(responseBody{
				Code:    errNotExists,
				Message: "A user with the provided username doesn't exist.",
			})
			return c.Status(http.StatusUnauthorized).Send(resp)
		}
		// There was an internal database error.
		return err
	}

	// Check if the password matches the hash.
	if err := bcrypt.CompareHashAndPassword(hash, []byte(data.Password)); err != nil {
		// The password didn't match.
		resp, _ := json.Marshal(responseBody{
			Code:    errIncorrectPassword,
			Message: "The password provided is incorrect.",
		})
		return c.Status(http.StatusUnauthorized).Send(resp)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tString, err := token.SignedString(signingKey)
	if err != nil {
		return err
	}

	resp, _ := json.Marshal(responseBody{
		Code:    okAuthenticated,
		Message: "Logged in.",
		Data: loginResponse{
			Token: tString,
		},
	})
	// Authenticated; we're done here.
	return c.Send(resp)
}

func handleRegistration(c *fiber.Ctx) error {
	return nil
}
