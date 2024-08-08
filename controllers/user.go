package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/nrednav/cuid2"
	"github.com/yash91989201/go_cart/configs"
	"github.com/yash91989201/go_cart/internal/database"
	"github.com/yash91989201/go_cart/models"
	"github.com/yash91989201/go_cart/utils"
)

type UserControllers struct {
	DB *database.Queries
}

func (c *UserControllers) CreateUser(w http.ResponseWriter, r *http.Request) {
	reqBody := models.CreateUserReq{}
	err := json.NewDecoder(r.Body).Decode(&reqBody)

	if err != nil {
		utils.RespondWithError(w, 400, "Incorrect data")
	}

	hashedPassword, err := utils.HashPassword(reqBody.Password)
	if err != nil {
		utils.RespondWithError(w, 400, "Incorrect data")
	}

	newUser, err := c.DB.InsertUser(r.Context(), database.InsertUserParams{
		ID:        cuid2.Generate(),
		CreatedAt: time.Now().UTC(),
		Name:      reqBody.Name,
		Email:     reqBody.Email,
		Password:  hashedPassword,
	})
	log.Print(err)
	if err != nil {
		utils.RespondWithError(w, 400, "Incorrect data1")
		return
	}

	utils.RespondWithJson(w, 201, newUser)
}

func (c *UserControllers) LoginUser(w http.ResponseWriter, r *http.Request) {

	reqBody := models.LoginUserReq{}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		utils.RespondWithError(w, 400, "Incorrect data")
		return
	}

	user, err := c.DB.GetUserByEmail(r.Context(), reqBody.Email)
	if err != nil {
		utils.RespondWithError(w, 404, "Invalid Credentials")
		return
	}

	if ok := utils.VerifyPassword(reqBody.Password, user.Password); !ok {

		utils.RespondWithError(w, 404, "Invalid Credentials")
		return
	}

	sessionId, err := c.DB.CreateSession(r.Context(), database.CreateSessionParams{
		ID:        cuid2.Generate(),
		ExpiresAt: time.Now().Add(time.Hour * 24),
		UserID:    user.ID,
	})

	if err != nil {
		utils.RespondWithError(w, 404, "Unable to login")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth-session",
		Value:    sessionId,
		Path:     "/",
		MaxAge:   int(time.Now().Add(time.Hour * 24).Unix()),
		HttpOnly: configs.GetEnv().ENV == "prod",
		Secure:   configs.GetEnv().ENV == "prod",
		SameSite: http.SameSiteLaxMode,
	})

}
