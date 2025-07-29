//go:build e2e
// +build e2e

package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func createToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte("gorestapikey"))
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

func TestPostComment(t *testing.T) {
	time.Sleep(time.Second * 1)
	t.Run("can post comment", func(t *testing.T) {
		client := resty.New()
		response, err := client.R().
			SetBody(`{"slug": "slug", "body": "body", "author": "author"}`).
			SetHeader("Authorization", "bearer "+createToken()).
			Post("http://localhost:8080/api/v1/comment")
		assert.NoError(t, err)
		assert.Equal(t, 200, response.StatusCode())
	})

	t.Run("cant post comment without JWT authorization", func(t *testing.T) {
		client := resty.New()
		response, err := client.R().
			SetBody(`{"slug": "slug", "body": "body", "author": "author"}`).
			Post("http://localhost:8080/api/v1/comment")
		assert.NoError(t, err)
		assert.Equal(t, 401, response.StatusCode())
	})
}
