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

package model

import (
	"com/novare/dbs"
	"errors"

	"gopkg.in/mgo.v2/bson"
)

var mDB = dbs.NewMongoDB(AuthRelayDatabaseName, "Companies")

/*
Company -
Every account starts with a company. Sites can have one or more accounts.
*/
type Company struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	Name       string        `json:"name"`
	Address1   string        `json:"address1"`
	Address2   string        `json:"address2"`
	City       string        `json:"city"`
	State      string        `json:"state"`
	Zip        string        `json:"zip"`
	IsLocation bool          `json:"isLocation"` //Specifies if a company is also a location. Used with the
	//AuthRelay - If the company is using a distributed authorization
	//environment.
	AuthRelay string `json:"authRelay"`
}

/*
NewCompany - Constructor for Company.
*/
func NewCompany() *Company {
	company := new(Company)
	company.ID = bson.NewObjectId()
	return company
}

//SaveCompany - Given a company, it will save it
//to the database. Note that the ID must be an existing
//ID in the database
func SaveCompany(company *Company) error {

	if company == nil {
		return errors.New("InvalidCompany")
	}

	return mDB.Update(company, bson.M{"_id": company.ID})
}

//InsertCompany - Add a company to the database
func InsertCompany(company *Company) error {

	if company == nil {
		return errors.New("InvalidCompany")
	}

	return mDB.Insert(company, bson.M{"_id": company.ID})
}

//FindCompanyByID - Return a company if it can find it
//or return an error
func FindCompanyByID(ID string) (*Company, error) {
	company := NewCompany()

	if !bson.IsObjectIdHex(ID) {
		return nil, errors.New("InvalidID")
	}

	err := mDB.Find(company, bson.M{"_id": bson.ObjectIdHex(ID)})
	if err != nil {
		return nil, err
	}

	return company, nil
}

//RemoveCompanyByID - Remove a company if it can find it
func RemoveCompanyByID(ID string) error {

	company, err := FindCompanyByID(ID)
	if err != nil {
		return err
	}

	return mDB.Remove(bson.M{"_id": company.ID})
}

//ListCompanies - List all companies.
//Initially there should not be that many. Eventually we will have to add
//another function to limit the number of records returned.
func ListCompanies() ([]Company, error) {
	var companies []Company
	err := mDB.List(&companies, bson.M{})
	return companies, err
}
