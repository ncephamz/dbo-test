package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ncephamz/efishery-be-test/api/auth"
	"github.com/ncephamz/efishery-be-test/api/models"
	"github.com/ncephamz/efishery-be-test/api/responses"
	"github.com/ncephamz/efishery-be-test/api/utils/formaterror"

	"golang.org/x/crypto/bcrypt"
)

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userExist, _ := user.FindUsersByPhone(server.STHQ)
	if len(userExist) > 0 {
		formattedError := formaterror.FormatError("Phone")
		responses.ERROR(w, http.StatusBadRequest, formattedError)
		return
	}

	userCreated, err := user.SaveUser(server.STHQ)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	resp := models.Login{
		Phone:    userCreated.Phone,
		Password: userCreated.Password,
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	responses.JSON(w, http.StatusCreated, resp)
}

func (server *Server) GenerateToken(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("token")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userExist, err := user.FindUsersByPhone(server.STHQ)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	err = models.VerifyPassword(user.Password, userExist[0].Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	token, err := auth.CreateToken(userExist[0].Phone, userExist[0].Name, userExist[0].Role, userExist[0].CreatedAt)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	resp := models.Token{
		token,
	}

	responses.JSON(w, http.StatusOK, resp)
}

func (server *Server) ClaimToken(w http.ResponseWriter, r *http.Request) {
	claim, err := auth.ExtractTokenClaims(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	responses.JSON(w, http.StatusOK, claim)
}
