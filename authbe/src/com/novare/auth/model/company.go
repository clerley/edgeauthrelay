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
	"log"
	"time"
	"unicode/utf8"

	"gopkg.in/mgo.v2/bson"
)

var mDBCompany = dbs.NewMongoDB(AuthRelayDatabaseName, "Companies")

//This constant define the valid password duration unit
const (
	//PassUnitDay - This comment is to avoid warnings from the IDE.
	PassUnitDay string = "Day"

	//PassUnitWeek - This comment is to avoid warnings from the IDE.
	PassUnitWeek string = "Week"

	//PassUnitMonth - This comment is to avoid warnings from the IDE.
	PassUnitMonth string = "Month"

	//PassUnitYear - This comment is to avoid warnings from the IDE.
	PassUnitYear string = "Year"
)

//CompanySettings ... All the settings related to a company
type CompanySettings struct {
	JWTDuration    int64  `json:"jwtDuration"`    //The number of minutes a JWT token should be granted 0 = Never expires
	PassExpiration int64  `json:"passExpiration"` //Password expiration... 0 means no expiration
	PassUnit       string `json:"passUnit"`       //Year, Month, Week, Days
}

//SetPasswordPolicy - It sets a policy on when the password should expire
func (settings *CompanySettings) SetPasswordPolicy(value int, unit string) {

	switch unit {

	case PassUnitDay:
		settings.PassExpiration = time.Now().AddDate(0, 0, value).Unix()
	case PassUnitMonth:
		settings.PassExpiration = time.Now().AddDate(0, value, 0).Unix()
	case PassUnitWeek:
		settings.PassExpiration = time.Now().AddDate(0, 0, value*7).Unix()
	default:
		log.Printf("Adding [%d] years to the password expiration", value)
		settings.PassExpiration = time.Now().AddDate(value, 0, 0).Unix()

	}

}

/*
Company -
Every account starts with a company. Sites can have one or more accounts.
*/
type Company struct {
	ID              bson.ObjectId   `bson:"_id" json:"id"`
	Name            string          `json:"name"`
	Address1        string          `json:"address1"`
	Address2        string          `json:"address2"`
	City            string          `json:"city"`
	State           string          `json:"state"`
	Zip             string          `json:"zip"`
	IsInLocation    bool            `json:"isInLocation"`    //Specifies if a company is also a location. Used with the
	RemotelyManaged bool            `json:"remotelyManaged"` //Is this Auth system managed remotely
	AuthRelay       string          `json:"authRelay"`       //If it is remotely managed, we need the path to it.
	UniqueID        string          `json:"uniqueID"`        //This must be provided in the request
	APIKey          string          `json:"apiKey"`          //The key required to access remote servers
	GroupOwnerID    string          `json:"groupOwnerID"`    //Group Owner ID
	MemberOfGroups  []string        `json:"memberOfGroups"`  //Groups this Company Belongs to
	Settings        CompanySettings `json:"settings"`        //Settings
}

/*
NewCompany - Constructor for Company.
*/
func NewCompany() *Company {
	company := new(Company)
	company.ID = bson.NewObjectId()
	return company
}

func isValidCompany(company *Company) bool {

	if company == nil {
		return false
	}

	if utf8.RuneCountInString(company.UniqueID) == 0 {
		log.Printf("A company without an UniqueID is not valid")
		return false
	}

	return true
}

//SaveCompany - Given a company, it will save it
//to the database. Note that the ID must be an existing
//ID in the database
func SaveCompany(company *Company) error {

	if !isValidCompany(company) {
		return errors.New("InvalidCompany")
	}

	return mDBCompany.Update(company, bson.M{"_id": company.ID})
}

//InsertCompany - Add a company to the database
func InsertCompany(company *Company) error {

	if !isValidCompany(company) {
		return errors.New("InvalidCompany")
	}

	return mDBCompany.Insert(company, bson.M{"_id": company.ID})
}

func findCompanyWithCondition(condition interface{}) (*Company, error) {
	company := NewCompany()

	err := mDBCompany.Find(company, condition)
	if err != nil {
		return nil, err
	}

	return company, nil
}

//FindCompanyByID - Return a company if it can find it
//or return an error
func FindCompanyByID(ID string) (*Company, error) {

	if !bson.IsObjectIdHex(ID) {
		return nil, errors.New("InvalidID")
	}

	return findCompanyWithCondition(bson.M{"_id": bson.ObjectIdHex(ID)})
}

//FindCompanyByUniqueID ...
func FindCompanyByUniqueID(uniqueID string) (*Company, error) {

	if utf8.RuneCountInString(uniqueID) == 0 {
		return nil, errors.New("UniqueIDTooShort")
	}

	return findCompanyWithCondition(bson.M{"uniqueid": uniqueID})
}

//RemoveCompanyByID - Remove a company if it can find it
func RemoveCompanyByID(ID string) error {

	company, err := FindCompanyByID(ID)
	if err != nil {
		return err
	}

	return mDBCompany.Remove(bson.M{"_id": company.ID})
}

//ListCompanies - List all companies.
//Initially there should not be that many. Eventually we will have to add
//another function to limit the number of records returned.
func ListCompanies() ([]Company, error) {
	var companies []Company
	err := mDBCompany.List(&companies, bson.M{})
	return companies, err
}
