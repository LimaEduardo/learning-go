package mongo_test

import (
	"log"
	"mongodb-example/pkg"
	"mongodb-example/pkg/mongo"
	"testing"
)

const (
	mongoUrl           = "localhost:27017"
	dbName             = "teste_db_go"
	userCollectionName = "user"
)

func Test_UserService(t *testing.T) {
	t.Run("CreateUser", createUser_should_insert_user_into_mongo)
}

func createUser_should_insert_user_into_mongo(t *testing.T) {
	session, err := mongo.NewSession(mongoUrl)
	if err != nil {
		log.Fatalf("Unable to connect to mongo: %s", err)
	}
	defer func() {
		session.DropDatabase(dbName)
		session.Close()
	}()

	userService := mongo.NewUserService(session.Copy(), dbName, userCollectionName)

	testUsername := "integration_test_user"
	testPassword := "integration_test_password"
	user := root.User{
		Username: testUsername,
		Password: testPassword,
	}

	err = userService.Create(&user)

	if err != nil {
		t.Error("Unable to create user: %T", err)
	}

	var results []root.User
	session.GetCollection(dbName, userCollectionName).Find(nil).All(&results)

	count := len(results)

	if count != 1 {
		t.Error("Incorrect number of results. Expected `1`, got `%i`", count)
	}

	if results[0].Username != user.Username {
		// t.Error("error in username")
		t.Error("Incorrect Username. Expected `%T`, Got: `%T`", testUsername, results[0].Username)
	}
}
