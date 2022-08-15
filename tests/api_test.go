package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"moments/api"
	"moments/db"
	"moments/ent"
	"moments/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegister(t *testing.T) {
	method := "POST"
	url := "/api/v1/auth/register"

	engine := api.RunEngine()

	dbc, cancel := db.New()
	defer cancel()
	defer dbc.Client.Close()

	user := ent.User{
		Username: "testuser",
		Email:    "testuser@email.com",
		Password: "testpass123",
	}

	jsonBody, err := json.Marshal(&user)
	assert.Nil(t, err)

	defer deleteUserByUsername(t, dbc, user.Username)

	response := httptest.NewRecorder()
	request, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	engine.ServeHTTP(response, request)

	var body api.AuthResponseSchema
	json.Unmarshal(response.Body.Bytes(), &body)

	assert.Equal(t, http.StatusCreated, response.Code)
	assert.True(t, utils.CheckPasswordHash(user.Password, body.User.Password))
	assert.Equal(t, user.Username, body.User.Username)
	assert.Equal(t, user.Email, body.User.Email)
}

func TestLogin(t *testing.T) {

	engine := api.RunEngine()

	method := "POST"
	url := "/api/v1/auth/login"

	dbc, cancel := db.New()
	defer cancel()
	defer dbc.Client.Close()

	user := ent.User{
		Username: "testuser",
		Email:    "testuser@email.com",
		Password: "testpass123",
	}
	createTestUserViaAPI(t, dbc, user.Username, user.Email)
	defer deleteUserByUsername(t, dbc, user.Username)

	requestBody, err := json.Marshal(&user)
	assert.Nil(t, err)

	response := httptest.NewRecorder()
	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	assert.Nil(t, err)
	engine.ServeHTTP(response, request)

	var schema api.AuthResponseSchema

	json.Unmarshal(response.Body.Bytes(), &schema)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.NotEmpty(t, schema.Token)
	assert.Equal(t, schema.User.Username, user.Username)
}

func TestCreatePrivateChat(t *testing.T) {

	engine := api.RunEngine()

	method := "POST"
	url := "/api/v1/private_chat/"

	dbc, cancel := db.New()
	defer cancel()
	defer dbc.Client.Close()

	testUser := createTestUserViaAPI(t, dbc, "testuser", "testuser@email.com")
	defer deleteUserByUsername(t, dbc, "testuser")

	createTestUserViaAPI(t, dbc, "testuser1", "testuser1@email.com")
	defer deleteUserByUsername(t, dbc, "testuser1")

	requestSchema := api.CreateRoomSchema{Usernames: []string{"testuser1"}}
	reqBody, err := json.Marshal(&requestSchema)

	response := httptest.NewRecorder()
	request, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	assert.Nil(t, err)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", testUser.Token))
	engine.ServeHTTP(response, request)

	var responseBody api.CreateRoomResponseSchema
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	t.Log(response.Body.String())

	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, "private", responseBody.CreatedRoom.Type.String())
	assert.Equal(t, 2, len(responseBody.Users))
}

func TestCreateDuplicatePrivateChatWithSameUsers(t *testing.T) {

	engine := api.RunEngine()

	method := "POST"
	url := "/api/v1/private_chat/"

	dbc, cancel := db.New()
	defer cancel()
	defer dbc.Client.Close()

	testUser := createTestUserViaAPI(t, dbc, "testuser", "testuser@email.com")
	defer deleteUserByUsername(t, dbc, "testuser")

	createTestUserViaAPI(t, dbc, "testuser1", "testuser1@email.com")
	defer deleteUserByUsername(t, dbc, "testuser1")

	requestSchema := api.CreateRoomSchema{Usernames: []string{"testuser1"}}
	reqBody, err := json.Marshal(&requestSchema)

	response := httptest.NewRecorder()
	request, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	assert.Nil(t, err)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", testUser.Token))
	engine.ServeHTTP(response, request)

	w := httptest.NewRecorder()
	request, err = http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	assert.Nil(t, err)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", testUser.Token))
	engine.ServeHTTP(w, request)

	assert.Equal(t, http.StatusConflict, w.Code)

}

func TestCreatePrivateChatWithReversedUsers(t *testing.T) {

	engine := api.RunEngine()

	method := "POST"
	url := "/api/v1/private_chat/"

	dbc, cancel := db.New()
	defer cancel()
	defer dbc.Client.Close()

	testUser := createTestUserViaAPI(t, dbc, "testuser", "testuser@email.com")
	defer deleteUserByUsername(t, dbc, "testuser")

	testUser1 := createTestUserViaAPI(t, dbc, "testuser1", "testuser1@email.com")
	defer deleteUserByUsername(t, dbc, "testuser1")

	requestSchema := api.CreateRoomSchema{Usernames: []string{"testuser"}}
	reqBody, err := json.Marshal(&requestSchema)

	response := httptest.NewRecorder()
	request, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	assert.Nil(t, err)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", testUser1.Token))
	engine.ServeHTTP(response, request)

	w := httptest.NewRecorder()
	request, err = http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	assert.Nil(t, err)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", testUser.Token))
	engine.ServeHTTP(w, request)

	assert.Equal(t, http.StatusConflict, w.Code)

}

func TestPostFollow(t *testing.T) {
	engine := api.RunEngine()

	method := "POST"
	url := "/api/v1/users/follow"

	dbc, cancel := db.New()
	defer cancel()
	defer dbc.Client.Close()

	firstUser := createTestUserViaAPI(t, dbc, "firstuser", "firstuser@email.com")
	defer deleteUserByUsername(t, dbc, firstUser.User.Username)

	secondUser := createTestUserViaAPI(t, dbc, "seconduser", "seconduser@email.com")
	defer deleteUserByUsername(t, dbc, secondUser.User.Username)

	usernames := []string{secondUser.User.Username}
	body, err := json.Marshal(&usernames)
	assert.Nil(t, err)
	response := httptest.NewRecorder()
	request, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", firstUser.Token))
	engine.ServeHTTP(response, request)

	var schema api.PostFollowSchema

	json.Unmarshal(response.Body.Bytes(), &schema)

	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, schema.FollowingList[0].Username, secondUser.User.Username)
}

func TestPostFollowDuplicate(t *testing.T) {
	engine := api.RunEngine()

	method := "POST"
	url := "/api/v1/users/follow"

	dbc, cancel := db.New()
	defer cancel()
	defer dbc.Client.Close()

	firstUser := createTestUserViaAPI(t, dbc, "firstuser", "firstuser@email.com")
	defer deleteUserByUsername(t, dbc, firstUser.User.Username)

	secondUser := createTestUserViaAPI(t, dbc, "seconduser", "seconduser@email.com")
	defer deleteUserByUsername(t, dbc, secondUser.User.Username)

	usernames := []string{secondUser.User.Username}
	body, err := json.Marshal(&usernames)
	assert.Nil(t, err)
	response := httptest.NewRecorder()
	request, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", firstUser.Token))
	engine.ServeHTTP(response, request)

	var schema api.PostFollowSchema

	json.Unmarshal(response.Body.Bytes(), &schema)

	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, schema.FollowingList[0].Username, secondUser.User.Username)

	response = httptest.NewRecorder()
	request, err = http.NewRequest(method, url, bytes.NewBuffer(body))
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", firstUser.Token))
	engine.ServeHTTP(response, request)

	json.Unmarshal(response.Body.Bytes(), &schema)

	t.Log(response.Body.String())

	assert.Equal(t, http.StatusConflict, response.Code)
	assert.Equal(t, schema.Message, "constraint error.")
}
