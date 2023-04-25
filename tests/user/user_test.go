package user

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestList_Should_ReturnEmpty(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/users")

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGet_Should_ReturnStatusNotFound(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/users/Eryk")

	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}
