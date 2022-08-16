package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"moments/api"
	"moments/db"
	"moments/ent"
	"moments/user"
	"net/http"
	"net/http/httptest"
	"testing"
)

func followViaAPI(t *testing.T, dbc *db.DatabaseConnection, token string, usernames ...string) (api.PostFollowSchema, []api.AuthResponseSchema) {
	engine := api.RunEngine()
	method := "POST"
	url := "/api/v1/users/follow"
	dbc, cancel := db.New()
	defer cancel()
	defer dbc.Client.Close()

	createdUsers := []api.AuthResponseSchema{}

	for _, s := range usernames {
		createdUsers = append(createdUsers, createTestUserViaAPI(t, dbc, s, fmt.Sprintf("%s@email.com", s)))
	}

	body, err := json.Marshal(&usernames)
	assert.Nil(t, err)
	response := httptest.NewRecorder()
	request, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	engine.ServeHTTP(response, request)
	var schema api.PostFollowSchema
	json.Unmarshal(response.Body.Bytes(), &schema)
	return schema, createdUsers

}

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
