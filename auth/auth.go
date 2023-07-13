package auth

import (
	
	"net/http"
	"orm/models"
	
	"time"

	jwt "github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
)
const(
	USER="admin"
	PASSWORD="password123"
	SECRET="secret"
)
func LoginHandler(c *gin.Context)  {
	var user models.Credential
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
	}

	if user.Username != USER {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message":"user invalid",
		})
	} else if  PASSWORD != user.Password {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message":"password invalid",
			})
	}else{

		claim := jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute *1).Unix(),
			Issuer: "test",
			IssuedAt: time.Now().Unix(),
		}
	
		sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
		token, err := sign.SignedString([]byte(SECRET))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			c.Abort()
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "succes",
			"token": token,
		})

	}
	

	
}