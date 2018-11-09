package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type MongoDBRepository struct {
	db *mgo.Database
}

func (m *MongoDBRepository) FindCounterById(id string) (*Counter, error) {
	coll := m.db.C(config.CounterColl)
	var counter *Counter
	err := coll.Find(bson.M{"_id": id}).One(&counter)
	if err != nil && err != mgo.ErrNotFound {
		log.Fatal("Error getting counter from database", err)
	}

	return counter, err
}

func (m *MongoDBRepository) UpsertCounter(counter Counter) error {
	coll := m.db.C(config.CounterColl)
	_, err := coll.UpsertId(counter.ID, counter)
	return err
}

func (m *MongoDBRepository) InsertLink(link Link) (error) {
	coll := m.db.C(config.LinksColl)
	return coll.Insert(link)
}

func (m *MongoDBRepository) FindLinkByShortIdInt(id uint) (*Link, error) {
	coll := m.db.C(config.LinksColl)
	var link Link
	err := coll.Find(bson.M{"short_id_int": id}).One(&link)
	if err != nil && err != mgo.ErrNotFound {
		log.Fatal("Error while finding link", err)
	}

	return &link, err
}

func (m *MongoDBRepository) UpdateLink(link Link) error {
	coll := m.db.C(config.LinksColl)
	return coll.UpdateId(link.ID, link)
}

func (m *MongoDBRepository) close() {
	m.db.Logout()
}

func NewMongoDBRepository() *MongoDBRepository {
	session, err := mgo.Dial(config.DBURL)
	if err != nil {
		log.Fatal(err)
	}
	db := session.DB(config.DBName)

	index := mgo.Index{
		Key:        []string{"short_id_int"},
		Unique:     true,
		DropDups:   true,
		Background: true,
	}

	err = db.C(config.LinksColl).EnsureIndex(index)
	if err != nil {
		log.Println("Failed to ensure index on short_id_int", err)
		log.Println("Try setting up index manually")
	}

	return &MongoDBRepository{db}
}
