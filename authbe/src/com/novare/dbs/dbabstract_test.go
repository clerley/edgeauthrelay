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
	"testing"
	"unicode/utf8"

	"gopkg.in/mgo.v2/bson"
)

const database = "TESTDB"

func TestDBAbstraction(t *testing.T) {

	db := NewMongoDB(database, "MyCollection")

	type testStruct struct {
		ID     bson.ObjectId `json:"id" bson:"_id"`
		Value1 string
		Value2 string
		Value3 string
		Value4 string
	}

	var test testStruct
	test.ID = bson.NewObjectId()
	test.Value1 = "Value 1"
	test.Value2 = "Value 2"
	test.Value3 = "Value 3"
	test.Value4 = "Value 4"

	err := db.Initialize()
	if err != nil {
		t.Errorf("The following error occurred: [%s]", err)
		return
	}

	err = db.Insert(&test, bson.M{"_id": test.ID})
	if err != nil {
		t.Errorf("The insert operation failed with error: [%s]", err)
		return
	}

	var test2 testStruct
	err = db.Find(&test2, bson.M{"_id": test.ID})
	if err != nil {
		t.Errorf("The record with ID:[%s] was not found", test.ID.Hex())
		return
	}

	if test2.ID != test.ID {
		t.Errorf("The expected ID and the ID retrieved do not match!")
		return
	}

	var lst []testStruct
	err = db.List(&lst, nil)
	if err != nil {
		t.Errorf("There was an error retrieving all the records from the database: [%s]", err)
		return
	}

	if len(lst) == 0 {
		t.Errorf("The list returned is empty! That should not have happend!")
		return
	}

	for i := range lst {
		t1 := lst[i]
		if utf8.RuneCountInString(t1.Value1) == 0 {
			t.Errorf("An error occurred, the value of Value1 should not be zero length")
		} else {
			t.Logf("ID for the testStruct %s Index: %d", t1.ID.Hex(), i)
		}
	}

	for i := range lst {
		t1 := lst[i]
		err = db.Remove(bson.M{"_id": t1.ID})
		if err != nil {
			t.Errorf("The following error occurred: %s", err)
		}
	}
}
