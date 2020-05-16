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
	"log"

	"gopkg.in/mgo.v2/bson"
)

var mDBCfg = dbs.NewMongoDB(AuthRelayDatabaseName, "Config")

//Config - Used for global configuration
type Config struct {
	ID     bson.ObjectId `json:"id" bson:"_id"`
	RowID  int           `json:"rowID"`
	Secret string        `json:"secret"`
}

//NewConfig - Application level configuration
func NewConfig() *Config {
	config := new(Config)
	config.ID = bson.NewObjectId()
	//We just need one configuration for the site.
	config.RowID = 1
	return config
}

//SaveConfig ...
func SaveConfig(config *Config) error {

	err := mDBCfg.Update(config, bson.M{"rowid": config.RowID})
	if err != nil {
		err = mDBCfg.Insert(config, bson.M{"rowid": config.RowID})
	}

	return err
}

//GetConfig ... It will either retrieve the configuration or a new Config Object
func GetConfig() *Config {
	cfg := NewConfig()
	err := mDBCfg.Find(cfg, bson.M{"rowid": 1})
	if err != nil {
		log.Printf("It must be the first instance of the Config object")
	}

	return cfg
}
