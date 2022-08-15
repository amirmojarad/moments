package tests

import (
	"github.com/stretchr/testify/assert"
	"moments/db"
	"moments/ent"
	r "moments/ent/room"
	"moments/messages"
	"moments/post"
	"moments/room"
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

	if err := user.DeleteUser(dbc, createdUser); err != nil {
		t.Fatal(err)
	}

}

// TestUserModelWithoutEmailAndPasswordAndUsername tests model when email field is empty
func TestUserModelWithoutEmailAndPasswordAndUsername(t *testing.T) {

	dbc, cancel := db.NewTestDB()
	defer dbc.Client.Close()
	defer cancel()

	u := &ent.User{}

	_, err := user.CreateUser(dbc, u)
	assert.Error(t, err, "asd")

}

func TestUserDelete(t *testing.T) {
	dbc, cancel := db.NewTestDB()
	defer dbc.Client.Close()
	defer cancel()

	u := &ent.User{
		Username: "testusername",
		Email:    "test@email.com",
		Password: "testpass123",
		IsAdmin:  false,
		IsStaff:  false,
	}

	createdUser, err := user.CreateUser(dbc, u)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, createdUser)

	err = user.DeleteUser(dbc, createdUser)
	assert.Nil(t, err)

	fetchedUser, err := user.GetUserByUsername(dbc, u.Username)
	assert.Nil(t, fetchedUser, err)
}

func TestGetUserByID(t *testing.T) {
	dbc, cancel := db.NewTestDB()
	defer dbc.Client.Close()
	defer cancel()

	u := &ent.User{
		Username: "testusername",
		Email:    "test@email.com",
		Password: "testpass123",
		IsAdmin:  false,
		IsStaff:  false,
	}

	createdUser, err := user.CreateUser(dbc, u)
	if err != nil {
		t.Fatal(err)
	}

	fetchedUser, err := user.GetUserByID(dbc, createdUser.ID)
	assert.NotNil(t, fetchedUser)

	err = user.DeleteUser(dbc, fetchedUser)
	assert.Nil(t, err)

}

func TestPostCreate(t *testing.T) {
	dbc, cancel := db.NewTestDB()
	defer dbc.Client.Close()
	defer cancel()

	createdUser := createTestUser(t, dbc, "firstUser", "firstUser@email.com")

	defer deleteUser(t, dbc, createdUser)

	p := &ent.Post{
		PlainText: "this is a sample plain text",
	}

	createdPost, err := post.CreatePost(dbc, createdUser, p)
	assert.Nil(t, err)
	assert.NotNil(t, createdPost)
	postOwner, err := post.GetPostOwner(dbc, createdPost)
	assert.Nil(t, err)
	assert.Equal(t, postOwner.Username, createdUser.Username)

}

func TestUpdatePost(t *testing.T) {
	dbc, cancel := db.NewTestDB()
	defer dbc.Client.Close()
	defer cancel()

	createdUser := createTestUser(t, dbc, "firstUser", "firstUser@email.com")
	defer deleteUser(t, dbc, createdUser)

	p := &ent.Post{
		PlainText: "this is a sample plain text",
	}

	createdPost, err := post.CreatePost(dbc, createdUser, p)

	assert.Nil(t, err)
	assert.NotNil(t, createdPost)

	link := "http://samplelink.com"

	updatedPost, err := post.UpdateLinkUrl(dbc, createdPost, link)
	assert.Nil(t, err)
	assert.NotNil(t, updatedPost)

	assert.Equal(t, updatedPost.LinkURL, link)
}

func TestFollow(t *testing.T) {
	dbc, cancel := db.NewTestDB()
	defer dbc.Client.Close()
	defer cancel()

	firstUser := createTestUser(t, dbc, "firstUser", "firstUser@email.com")
	defer deleteUser(t, dbc, firstUser)

	secondUser := createTestUser(t, dbc, "secondUser", "seconduser@email.com")
	defer deleteUser(t, dbc, secondUser)

	firstUser, err := user.AddFollowing(dbc, firstUser, secondUser)
	assert.Nil(t, err)
	assert.NotNil(t, firstUser)

	firstFollowers, err := user.GetAllFollowers(dbc, firstUser)
	assert.Equal(t, 0, len(firstFollowers))
	firstFollowing, err := user.GetAllFollowing(dbc, firstUser)
	assert.Equal(t, 1, len(firstFollowing))

	secondFollowing, err := user.GetAllFollowing(dbc, secondUser)
	assert.Equal(t, 0, len(secondFollowing))

	secondFollowers, err := user.GetAllFollowers(dbc, secondUser)
	assert.Equal(t, 1, len(secondFollowers))

	following, err := user.IsFollowing(dbc, firstUser, secondUser)
	assert.Nil(t, err)
	assert.NotNil(t, following)
	assert.Equal(t, following.Username, secondUser.Username)
}

func TestUnfollow(t *testing.T) {
	dbc, cancel := db.NewTestDB()
	defer dbc.Client.Close()
	defer cancel()

	firstUser := createTestUser(t, dbc, "firstUser", "firstUser@email.com")
	defer deleteUser(t, dbc, firstUser)

	secondUser := createTestUser(t, dbc, "secondUser", "seconduser@email.com")
	defer deleteUser(t, dbc, secondUser)

	firstUser, err := user.AddFollowing(dbc, firstUser, secondUser)
	assert.Nil(t, err)
	assert.NotNil(t, firstUser)
	following, err := user.GetAllFollowing(dbc, firstUser)
	assert.Equal(t, 1, len(following))

	firstUser, err = user.RemoveFollowing(dbc, firstUser, secondUser)
	assert.Nil(t, err)
	assert.NotNil(t, firstUser)
	following, err = user.GetAllFollowing(dbc, firstUser)
	assert.Equal(t, 0, len(following))

}

func TestCreatePrivateRoom(t *testing.T) {
	dbc, cancel := db.NewTestDB()
	defer dbc.Client.Close()
	defer cancel()

	firstUser := createTestUser(t, dbc, "firstUser", "firstUser@email.com")
	defer deleteUser(t, dbc, firstUser)

	secondUser := createTestUser(t, dbc, "secondUser", "seconduser@email.com")
	defer deleteUser(t, dbc, secondUser)

	newPrivateRoom, err := room.CreatePrivateRoom(dbc, firstUser, secondUser)
	assert.Nil(t, err)
	assert.NotNil(t, newPrivateRoom)

	result, err := room.IsUserInRoom(dbc, firstUser, newPrivateRoom.ID)
	assert.Nil(t, err)
	assert.True(t, result)

	result, err = room.IsUserInRoom(dbc, secondUser, newPrivateRoom.ID)
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestCreateDuplicatePrivateRoom(t *testing.T) {
	dbc, cancel := db.NewTestDB()
	defer dbc.Client.Close()
	defer cancel()

	firstUser := createTestUser(t, dbc, "firstUser", "firstUser@email.com")
	defer deleteUser(t, dbc, firstUser)

	secondUser := createTestUser(t, dbc, "secondUser", "seconduser@email.com")
	defer deleteUser(t, dbc, secondUser)

	newPrivateRoom, err := room.CreatePrivateRoom(dbc, firstUser, secondUser)
	assert.Nil(t, err)
	assert.NotNil(t, newPrivateRoom)

	result, err := room.IsUserInRoom(dbc, firstUser, newPrivateRoom.ID)
	assert.Nil(t, err)
	assert.True(t, result)

	result, err = room.IsUserInRoom(dbc, secondUser, newPrivateRoom.ID)
	assert.Nil(t, err)
	assert.True(t, result)

	newPrivateRoom1, err := room.CreatePrivateRoom(dbc, firstUser, secondUser)
	assert.NotNil(t, err)
	assert.Nil(t, newPrivateRoom1)

}

func TestCreatePublicRoom(t *testing.T) {
	dbc, cancel := db.NewTestDB()
	defer dbc.Client.Close()
	defer cancel()

	firstUser := createTestUser(t, dbc, "firstUser", "firstUser@email.com")
	defer deleteUser(t, dbc, firstUser)

	secondUser := createTestUser(t, dbc, "secondUser", "seconduser@email.com")
	defer deleteUser(t, dbc, secondUser)

	thirdUser := createTestUser(t, dbc, "thirdUser", "thirdUser@email.com")
	defer deleteUser(t, dbc, thirdUser)

	newPublicRoom, err := room.CreatePublicRoom(dbc, "testGroup", firstUser, thirdUser, secondUser)
	assert.Nil(t, err)
	assert.NotNil(t, newPublicRoom)

	result, err := room.IsUserInRoom(dbc, firstUser, newPublicRoom.ID)
	assert.Nil(t, err)
	assert.True(t, result)

	result, err = room.IsUserInRoom(dbc, secondUser, newPublicRoom.ID)
	assert.Nil(t, err)
	assert.True(t, result)

	result, err = room.IsUserInRoom(dbc, thirdUser, newPublicRoom.ID)
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestCreateMessage(t *testing.T) {
	dbc, cancel := db.NewTestDB()
	defer dbc.Client.Close()
	defer cancel()

	firstUser := createTestUser(t, dbc, "firstUser", "firstUser@email.com")
	defer deleteUser(t, dbc, firstUser)

	secondUser := createTestUser(t, dbc, "secondUser", "seconduser@email.com")
	defer deleteUser(t, dbc, secondUser)

	newPrivateRoom, err := room.CreatePrivateRoom(dbc, firstUser, secondUser)
	assert.Nil(t, err)
	assert.NotNil(t, newPrivateRoom)

	newMessage, err := messages.Create(dbc, firstUser, newPrivateRoom, "test text")
	assert.Nil(t, err)
	assert.NotNil(t, newMessage)
	assert.Equal(t, newMessage.Text, "test text")

	senderUser, err := messages.GetMessageSenderUser(dbc, newMessage)
	assert.Nil(t, err)
	assert.NotNil(t, senderUser)
	assert.Equal(t, senderUser.Username, firstUser.Username)

	roomOwner, err := messages.GetMessageRoomOwner(dbc, newMessage)
	assert.Nil(t, err)
	assert.NotNil(t, senderUser)
	assert.Equal(t, newPrivateRoom.ID, roomOwner.ID)
}

func TestTemp(t *testing.T) {
	dbc, cancel := db.NewTestDB()
	defer dbc.Client.Close()
	defer cancel()

	firstUser := createTestUser(t, dbc, "firstUser", "firstUser@email.com")
	defer deleteUser(t, dbc, firstUser)

	secondUser := createTestUser(t, dbc, "secondUser", "seconduser@email.com")
	defer deleteUser(t, dbc, secondUser)

	room.CreatePublicRoom(dbc, "title", firstUser, secondUser)
	room.CreatePublicRoom(dbc, "title", firstUser, secondUser)

	_, err := dbc.Client.Room.Query().Where(r.TitleEQ("title")).Only(*dbc.Ctx)
	t.Log(err)
}
