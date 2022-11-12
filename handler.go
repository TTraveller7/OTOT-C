package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func authorize(c *gin.Context) {
	abortWithErrorMsg := func() {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}

	// Extract username

	userCookieStr, err := c.Cookie("Authorization")
	if err != nil {
		log.Error("", err)
		abortWithErrorMsg()
		return
	}

	username, isValid := tryParseUserJwt(userCookieStr)
	if !isValid {
		abortWithErrorMsg()
		return
	}

	u, err := getUserByUsername(username)
	if err != nil {
		abortWithErrorMsg()
		return
	}

	c.Set("role", u.Role)
}

func staffAuthorize(c *gin.Context) {
	role := c.GetString("role")
	if role == ROLE_STAFF || role == ROLE_MANAGER {
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusUnauthorized)
	}
}

func managerAuthorize(c *gin.Context) {
	role := c.GetString("role")
	if role == ROLE_MANAGER {
		c.Status(http.StatusOK)
	} else if role == ROLE_STAFF {
		c.Status(http.StatusForbidden)
	} else {
		c.Status(http.StatusUnauthorized)
	}
}

func authenticate(c *gin.Context) {
	u := &User{}
	if err := c.ShouldBind(u); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := saveUser(u); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Generate jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": u.Username,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(), // Jwt expires after 30 days
		"iat": time.Now().Unix(),                          // issue time
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(SECRET)) // SECRET here should be a env variable e.g. []byte(os.Getenv("SECRET"))

	if err != nil {
		log.Error("", err)
		c.String(http.StatusBadRequest, "Failed to create JWT")
		return
	}

	// Sent jwt back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully generate JWT",
	})
}
