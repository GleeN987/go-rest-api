//go:build e2e
// +build e2e

package tests

import (
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckEndpoint(t *testing.T) {
	time.Sleep(time.Second * 1)
	client := resty.New()
	response, err := client.R().Get("http://localhost:8080/alive")
	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode())
}
