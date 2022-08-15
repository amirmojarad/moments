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

func createTestUserViaAPI(t *testing.T, dbc *db.DatabaseConnection, username, email string) api.AuthResponseSchema {
	method := "POST"
	url := "/api/v1/auth/register"

	engine := api.RunEngine()

	u := ent.User{
		Username: username,
		Email:    email,
		Password: "testpass123",
	}

	jsonBody, err := json.Marshal(&u)
	assert.Nil(t, err)

	response := httptest.NewRecorder()
	request, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	engine.ServeHTTP(response, request)

	var body api.AuthResponseSchema
	json.Unmarshal(response.Body.Bytes(), &body)

	return body
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

func deleteUserByUsername(t *testing.T, connection *db.DatabaseConnection, username string) {
	fetchedUser, err := user.GetUserByUsername(connection, username)
	assert.NotNil(t, fetchedUser)
	assert.Nil(t, err)
	deleteUser(t, connection, fetchedUser)
}
