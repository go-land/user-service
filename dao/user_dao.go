package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	userDb          = "user"
	userProfileName = "user_profile"
)

type UserProfile struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Alias     string        `json:"alias" bson:"alias"`
	FirstName string        `json:"firstName" bson:"firstName"`
	LastName  string        `json:"lastName" bson:"lastName"`
}

type UserDao interface {
	GetAll() error
}

type UserMongoDao struct {
	session *mgo.Session
}

func (repo *UserMongoDao) GetAll() error {

	var userProfiles *[]UserProfile

	repo.collection().Find(nil).All(&userProfiles)

	return nil
}

func (repo *UserMongoDao) collection() *mgo.Collection {
	return repo.session.DB(userDb).C(userProfileName)
}

func NewUserDao(session *mgo.Session) UserDao {
	userDao := UserMongoDao{
		session: session,
	}

	return &userDao
}
