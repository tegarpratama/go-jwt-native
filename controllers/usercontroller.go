package controllers

import (
	"go-jwt/helpers"
	"go-jwt/models"
	"net/http"
)

func Me(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("userinfo").(*helpers.MyCustomClaims)
	userResponse := &models.MyProfile{
		ID:    user.ID,
		Name:  user.Email,
		Email: user.Email,
	}

	helpers.Response(w, 200, "My Profile", userResponse)
}
