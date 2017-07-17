package main

import (
	"gopkg.in/mgo.v2/bson"
)

type user struct {
	ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Email string        `json:"email"`
}

func insertUser(email string) (user, error) {
	result := user{bson.NewObjectId(), email}

	s := session.Clone()
	defer s.Close()

	c := s.DB("golang").C("user")
	err := c.Insert(result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func findUserByID(id string) (user, error) {
	result := user{}

	s := session.Clone()
	defer s.Close()

	c := s.DB("golang").C("user")
	err := c.FindId(bson.ObjectIdHex(id)).One(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}
