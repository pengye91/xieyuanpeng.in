package db

import (
	"github.com/pengye91/xieyuanpeng.in/backend/configs"
	"github.com/pengye91/xieyuanpeng.in/backend/utils/log"
	"gopkg.in/mgo.v2"
)

const (
	Host         = configs.BASE_MONGOURL
	Database     = "xieyuanpeng"
	AuthDatabase = "admin"
	AuthUserName = configs.MONGO_AUTH_USERNAME
	AuthPassword = configs.MONGO_AUTH_PASSWORD
)

var (
	mainSession *mgo.Session
	mainDb      *mgo.Database
)

type MgoDb struct {
	Session *mgo.Session
	Db      *mgo.Database
	Col     *mgo.Collection
}

func init() {

	if mainSession == nil {
		url := "mongodb://" + AuthUserName + ":" + AuthPassword + "@" + Host + "/" + AuthDatabase

		var err error
		mainSession, err = mgo.Dial(url)

		if err != nil {
			log.LoggerSugar.Errorw("Dial Mongo Error",
				"module", "mongo",
				"error", err,
			)
			panic(err)
		}
		mainSession.SetMode(mgo.Monotonic, true)
		mainDb = mainSession.DB(Database)
	}
}

func (this *MgoDb) Init() *mgo.Session {

	this.Session = mainSession.Copy()
	this.Db = this.Session.DB(Database)

	return this.Session
}

func (this *MgoDb) C(collection string) *mgo.Collection {
	this.Col = this.Session.DB(Database).C(collection)
	return this.Col
}

func (this *MgoDb) Close() bool {
	defer this.Session.Close()
	return true
}

func (this *MgoDb) DropoDb() {
	err := this.Session.DB(Database).DropDatabase()
	if err != nil {
		panic(err)
	}
}

func (this *MgoDb) RemoveAll(collection string) bool {
	this.Session.DB(Database).C(collection).RemoveAll(nil)

	this.Col = this.Session.DB(Database).C(collection)
	return true
}

func (this *MgoDb) Index(collection string, keys []string) bool {

	index := mgo.Index{
		Key:        keys,
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := this.Db.C(collection).EnsureIndex(index)
	if err != nil {
		log.LoggerSugar.Errorw("Mongo EnsureIndex Error",
			"module", "mongo",
			"error", err,
		)
		//panic(err)
		return false
	}
	return true
}

func (this *MgoDb) IsDup(err error) bool {

	if mgo.IsDup(err) {
		return true
	}

	return false
}
