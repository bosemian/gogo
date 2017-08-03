package model

import mgo "gopkg.in/mgo.v2"

var mongoSession *mgo.Session

const database = "gonews"

// Init func to connect mongoSession
func Init(mongoURL string) error {
	var err error
	mongoSession, err = mgo.Dial(mongoURL)
	if err != nil {
		return err
	}
	return nil
}
