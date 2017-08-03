package model

import (
	"fmt"
	"unicode/utf8"

	"gopkg.in/mgo.v2"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// Register user from request
func Register(username, password string) error {
	if utf8.RuneCountInString(username) < 4 {
		return fmt.Errorf("username must be > 4 character")
	}
	if utf8.RuneCountInString(password) < 6 {
		return fmt.Errorf("password must be > 6 character")
	}
	hpwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	s := mongoSession.Copy()
	defer s.Close()
	s.DB(database).C("users").Insert(bson.M{
		"username": username,
		"password": string(hpwd),
	})
	if err != nil {
		return err
	}

	return nil
}

// Login login account by uname password
func Login(username, password string) (string, error) {
	s := mongoSession.Copy()
	defer s.Close()
	var user bson.M
	err := s.DB(database).
		C("users").
		Find(bson.M{"username": username}).
		One(&user)
	if err == mgo.ErrNotFound {
		return "", fmt.Errorf("username or password wrong")
	}
	if err != nil {
		return "", err
	}
	hpwd := user["password"].(string)
	err = bcrypt.CompareHashAndPassword([]byte(hpwd), []byte(password))
	if err != nil {
		return "", fmt.Errorf("username or password invalid")
	}
	userID, _ := user["_id"].(bson.ObjectId)

	return userID.Hex(), nil

}

// CheckUserID check id for set cookies
func CheckUserID(id string) (bool, error) {
	s := mongoSession.Copy()
	defer s.Close()
	if !bson.IsObjectIdHex(id) {
		return false, fmt.Errorf("invalid id")
	}
	objectID := bson.ObjectIdHex(id)
	n, err := s.
		DB(database).
		C("users").
		FindId(objectID).
		Count()
	if err != nil {
		return false, err
	}
	if n <= 0 {
		return false, nil
	}
	return true, nil
}
