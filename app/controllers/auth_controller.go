package controllers

import (
	"context"
	"time"

	""

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UserSignUp(c *fiber.Ctx) error {
	signUp := &models.SignUp{}

}