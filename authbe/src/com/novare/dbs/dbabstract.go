/*
MIT License

Copyright (c) 2020 Clerley Silveira

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package dbs

import (
	"errors"
	"log"
	"os"
	"unicode/utf8"

	"gopkg.in/mgo.v2"
)

/*
MongoDB ...
I am not worried about making this too generic for now
*/
type MongoDB struct {
	//Private stuff
	dbURL      string
	mgoSession *mgo.Session

	//Public stuff
	DBName   string
	CollName string //Collection name
}

/*
NewMongoDB ...
Constructor for the MongoDB instance.
Create a new instance of MongoDB
*/
func NewMongoDB(dbName string, collName string) *MongoDB {
	mongo := new(MongoDB)
	mongo.DBName = dbName
	//We will not support passing the URL initially
	//If there is an use case later on, we can reconsider.
	mongo.dbURL = ""
	mongo.CollName = collName
	return mongo
}

/*Initialize -
Look for the environment variable AUTH_DB_URL
If that is found use otherwise, assume it is for testing purposes only
and connect straight to localhost.
*/
func (mongo *MongoDB) Initialize() error {
	mgoURL := os.Getenv("AUTH_DB_URL")
	if utf8.RuneCountInString(mgoURL) == 0 {
		log.Printf("The AUTH_DB_URL parameter is not defined, defaulting to localhost.")
		mongo.dbURL = "localhost"
	} else {
		mongo.dbURL = mgoURL
	}

	if utf8.RuneCountInString(mongo.CollName) == 0 {
		log.Printf("The collection name must be defined or the abstraction will not work!")
		return errors.New("InvalidCollectionName")
	}
	//Let the called decide what to do next
	return mongo.dialDatabase()
}

func (mongo *MongoDB) dialDatabase() error {
	tmp, err := mgo.Dial(mongo.dbURL)
	if err != nil {
		log.Printf("The database connection failed")
		return errors.New("InvalidDBConnection")
	}
	mongo.mgoSession = tmp
	return nil
}

/*getMongoSession - This function will copy the mongoDB session*/
func (mongo *MongoDB) getMongoSession() (*mgo.Session, error) {
	var sess *mgo.Session
	if mongo.mgoSession != nil {
		//log.Printf("The session is existent, trying and copying the session now")
		sess = mongo.mgoSession.Copy()
		return sess, nil
	}

	err := mongo.dialDatabase()
	if err != nil {
		log.Printf("The connection to the database failed %s", err)
		return nil, errors.New("InvalidDbConnection")
	}

	sess = mongo.mgoSession.Copy()
	return sess, nil
}

//Insert an object into the database, expecting a pointer
func (mongo *MongoDB) Insert(obj interface{}, condition interface{}) error {

	mongoCon, err := mongo.getMongoSession()
	if err != nil {
		log.Printf("The Database session has not been found! ERROR: [%s]", err)
		return errors.New("GenericDatabaseFailure")
	}
	defer mongoCon.Close()

	// Optional. Switch the session to a monotonic behavior.
	mongoCon.SetMode(mgo.Monotonic, true)

	c := mongoCon.DB(mongo.DBName).C(mongo.CollName)

	cnt, err := mongoCon.DB(mongo.DBName).C(mongo.CollName).Find(condition).Count()
	if err != nil && cnt > 0 {
		return errors.New("ConditionFailed")
	}

	// Insert the customer into the database
	err = c.Insert(obj)

	if err != nil {
		log.Printf("Database insertion failed with ERROR: [%s]", err)
		return errors.New("GenericDatabaseFailure")
	}

	return nil
}

//Find - Find a customer or return with an error
func (mongo *MongoDB) Find(resObj interface{}, condition interface{}) error {

	sess, err := mongo.getMongoSession()
	if err != nil {
		log.Printf("Connection is not valid ")
		return errors.New("Invalid database connection")
	}
	defer sess.Close()

	mgoCursor := sess.DB(mongo.DBName).C(mongo.CollName)
	err = mgoCursor.Find(condition).One(resObj)
	if err != nil {
		log.Printf("Error retrieving data from datasase [%s]", err)
		return errors.New("Notfound")
	}

	return nil
}

//Update -
func (mongo *MongoDB) Update(obj interface{}, condition interface{}) error {
	sess, err := mongo.getMongoSession()
	if err != nil {
		log.Printf("Generic Database failure [%s]", err)
		return errors.New("GenericDatabaseFailure")
	}
	defer sess.Close()

	// Optional. Switch the session to a monotonic behavior.
	sess.SetMode(mgo.Monotonic, true)

	c := sess.DB(mongo.DBName).C(mongo.CollName)

	// Update the customer into the database
	err = c.Update(condition, obj)

	if err != nil {
		log.Printf("Update failed! (%s) ", err)
		return errors.New("NotFound")
	}

	return nil
}

//List - List all for customer
func (mongo *MongoDB) List(objs interface{}, condition interface{}) error {

	sess, err := mongo.getMongoSession()
	if err != nil {
		log.Printf("The DB has not been found [%s]", err)
		return errors.New("GenericDatabaseFailure")
	}
	defer sess.Close()

	// Optional. Switch the session to a monotonic behavior.
	sess.SetMode(mgo.Monotonic, true)

	c := sess.DB(mongo.DBName).C(mongo.CollName)

	err = c.Find(condition).All(objs)
	if err != nil {
		log.Printf("Object - Finding error [%s]", err)
		return errors.New("NotFound")
	}

	return nil
}

//Remove - Find a customer or return with an error
func (mongo *MongoDB) Remove(condition interface{}) error {

	sess, err := mongo.getMongoSession()
	if err != nil {
		log.Printf("Connection is not valid ")
		return errors.New("Invalid database connection")
	}
	defer sess.Close()

	mgoCursor := sess.DB(mongo.DBName).C(mongo.CollName)
	err = mgoCursor.Remove(condition)
	if err != nil {
		log.Printf("Error removing data from database [%s]", err)
		return errors.New("Not found")
	}

	return nil
}
