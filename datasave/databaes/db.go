package databaes

import 	"gopkg.in/mgo.v2"

func Init() *mgo.Session {
	dail_info := &mgo.DialInfo{
		Addrs:          nil,
		Direct:         false,
		Timeout:        0,
		FailFast:       false,
		Database:       "",
		ReplicaSetName: "",
		Source:         "",
		Service:        "",
		ServiceHost:    "",
		Mechanism:      "",
		Username:       "",
		Password:       "",
		PoolLimit:      0,
		DialServer:     nil,
		Dial:           nil,
	}
	session, err := mgo.DialWithInfo(dail_info)
	if err != nil {
		panic(err)
	}
	return session
}

