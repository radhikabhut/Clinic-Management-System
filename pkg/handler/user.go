package handler

import (
	"clinic-management/pkg/db"
	"clinic-management/pkg/model"
	"clinic-management/pkg/utils"
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(req model.UserRequest) (string, int, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("password hashing failed", "error", err.Error())
		return "", 500, fmt.Errorf("internal server error")
	}

	data := model.User{
		Id:       uuid.NewString(),
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     req.Role,
	}
	err = db.CreateUser(data)
	if err != nil {

		slog.Error("error while creating record in database", "error", err.Error())
		return "", 500, fmt.Errorf("internal server error")
	}

	slog.Info("Patient record inserted successfully")
	return data.Id, 201, nil
}

func LoginUser(data model.LoginRequest) (string, error) {
	user, err := db.GetUserByEmail(data.Email)
	if err != nil {
		slog.Error("error while getting user from database", "error", err.Error())
		if err == pgx.ErrNoRows {
			return "", errors.New("user not found")
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		slog.Error("password mismatch", "error", err.Error())
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateJWT(user.Email, user.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}
