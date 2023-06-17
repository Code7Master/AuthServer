package controller

import (
	"DevelopHub/AuthServer/initializers"
	"DevelopHub/AuthServer/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(context *gin.Context) {

	// Email과 Password를 request body에서 얻습니다.
	var body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if context.BindJSON(&body) != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// password를 Hash 알고리즘을 이용해 Hashing 합니다.
	hashed, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	// User 모델(테이블)에 레코드를 생성합니다.
	user := models.User{Username: body.Username, Email: body.Email, Password: string(hashed)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user recode",
		})

		return
	}

	// response
	context.JSON(http.StatusOK, gin.H{})
}
