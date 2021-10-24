package handler

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"go-study/web/user"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	assert := assert.New(t)

	req := httptest.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()

	mux := NewHttpRouter()
	mux.ServeHTTP(resp, req)

	assert.Equal(http.StatusOK, resp.Code)
	data, err := ioutil.ReadAll(resp.Body)
	assert.Nil(err)
	assert.Equal("Hello World", string(data))
}

func TestFooHandler_WithoutNameParameter(t *testing.T) {
	assert := assert.New(t)

	req := httptest.NewRequest("GET", "/foo", nil)
	resp := httptest.NewRecorder()

	mux := NewHttpRouter()
	mux.ServeHTTP(resp, req)

	assert.Equal(http.StatusOK, resp.Code)
	data, err := ioutil.ReadAll(resp.Body)
	assert.Nil(err)
	assert.Equal("Hello World", string(data))
}

func TestFooHandler_WithNameParameter(t *testing.T) {
	assert := assert.New(t)

	req := httptest.NewRequest("GET", "/foo?name=Foo", nil)
	resp := httptest.NewRecorder()

	mux := NewHttpRouter()
	mux.ServeHTTP(resp, req)

	assert.Equal(http.StatusOK, resp.Code)
	data, err := ioutil.ReadAll(resp.Body)
	assert.Nil(err)
	assert.Equal("Hello Foo", string(data))
}

func TestCreateUserHandler(t *testing.T) {
	assert := assert.New(t)

	username := "foo"
	userEmail := "foo@email.com"
	newUser := user.NewUser(username, userEmail)
	data, _ := json.Marshal(newUser)
	reader := strings.NewReader(string(data))

	req := httptest.NewRequest("POST", "/user", reader)
	resp := httptest.NewRecorder()

	mux := NewHttpRouter()
	mux.ServeHTTP(resp, req)

	assert.Equal(http.StatusCreated, resp.Code)

	createdUser := new(user.User)
	err := json.NewDecoder(resp.Body).Decode(createdUser)

	assert.Nil(err)
	assert.Equal(username, createdUser.Name)
	assert.Equal(userEmail, createdUser.Email)
	assert.NotNil(createdUser.CreatedAt)
}