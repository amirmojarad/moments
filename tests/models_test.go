package tests

import (
	"github.com/stretchr/testify/assert"
	"moments/db"
	"moments/ent"
	"moments/user"
	"testing"
)

func init() {
	loadDbEnv()
}

//TestUserModel test user model and its create method
func TestUserModel(t *testing.T) {

	dbc, cancel := db.NewTestDB()
	defer dbc.Client.Close()
	defer cancel()
	u := &ent.User{
		Username: "testuser",
		Password: "testpass123",
		Email:    "testuser@email.com",
	}
	createdUser, err := user.CreateUser(dbc, u)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, createdUser.Email, u.Email)
	assert.Equal(t, createdUser.Username, u.Username)
	assert.NotEqual(t, createdUser.Password, u.Password)

	if err := user.DeleteUser(dbc, createdUser); err != nil {
		t.Fatal(err)
	}
}

func TestSuperUserModel(t *testing.T) {
	dbc, cancel := db.NewTestDB()
	defer dbc.Client.Close()
	defer cancel()
	u := &ent.User{
		Username: "testuser",
		Password: "testpass123",
		Email:    "testuser@email.com",
		IsAdmin:  true,
		IsStaff:  true,
	}
	createdUser, err := user.CreateSuperUser(dbc, u)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, createdUser.Email, u.Email)
	assert.Equal(t, createdUser.Username, u.Username)
	assert.NotEqual(t, createdUser.Password, u.Password)
	assert.True(t, createdUser.IsStaff)
	assert.True(t, createdUser.IsAdmin)

}

// TestUserModelWithoutEmailAndPasswordAndUsername tests model when email field is empty
func TestUserModelWithoutEmailAndPasswordAndUsername(t *testing.T) {
	defer db.Clear()

	dbc, cancel := db.NewTestDB()
	defer dbc.Client.Close()
	defer cancel()

	u := &ent.User{}

	_, err := user.CreateUser(dbc, u)
	assert.Error(t, err, "asd")
}
