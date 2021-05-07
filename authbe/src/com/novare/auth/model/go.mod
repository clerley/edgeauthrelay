module com/novare/auth/model

go 1.15

require (
	com/novare/dbs v0.0.0
	com/novare/utils v0.0.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)

replace com/novare/dbs v0.0.0 => ../../dbs

replace com/novare/utils v0.0.0 => ../../utils
