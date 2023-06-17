package controller

import (
	"DevelopHub/AuthServer/initializers"
	"DevelopHub/AuthServer/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
			// "error": "Failed to hash password",
			"error": "Already created a username or email",
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

func Login(context *gin.Context) {
	// username, email, password를 request body에서 얻는다.
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

	// 요청한 body가 User모델(테이블)에 있는지 확인한다.
	var user models.User
	if body.Username != "" {
		initializers.DB.First(&user, "username = ?", body.Username)
	} else if body.Email != "" {
		initializers.DB.First(&user, "email = ?", body.Email)
	}

	if user.ID == 0 /*Empty set*/ {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invaild username or email or password",
		})

		return
	}

	// 보낸 body의 password와 모델(테이블)의 password를 비교한다.
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or passsowrd",
		})

		return
	}

	// JWT 토큰 생성
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})

		return
	}

	// Token을 되돌려 주다.
	context.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	context.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
