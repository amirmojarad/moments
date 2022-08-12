package tests

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"moments/api"
	"moments/db"
	"moments/ent"
	"moments/user"
	"net/http"
	"net/http/httptest"
	"testing"
)

const AUTH_URL = "/api/v1/auth"

type temp struct {
}

func createTestUser(t *testing.T, connection *db.DatabaseConnection, username, email string) *ent.User {
	u := &ent.User{
		Username: username,
		Password: "testpass123",
		Email:    email,
	}
	createdUser, err := user.CreateUser(connection, u)
	if err != nil {
		t.Fatal(err)
	}
	return createdUser
}

func deleteUser(t *testing.T, connection *db.DatabaseConnection, u *ent.User) {
	if u != nil {
		err := user.DeleteUser(connection, u)
		assert.Nil(t, err)
	}
}

func TestRegister(t *testing.T) {
	method := "POST"
	url := AUTH_URL + "/register"
	testUser := ent.User{
		Username: "testuser",
		Email:    "testuser@gmail.com",
		Password: "testuser",
	}

	engine := api.RunEngine()

	payload, _ := json.Marshal(testUser)
	request, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	assert.Nil(t, err)

	response := httptest.NewRecorder()

	engine.ServeHTTP(response, request)

	assert.Equal(t, http.StatusCreated, response.Code)

	var res map[string]interface{}

	err = json.Unmarshal(response.Body.Bytes(), &res)
	if err != nil {
		panic(err)
	}
	assert.NotEmpty(t, res["token"])
	assert.NotEmpty(t, res["message"])
	assert.NotEmpty(t, res["user"])
}
