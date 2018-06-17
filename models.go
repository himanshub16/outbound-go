package main

import (
	"errors"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

var baseconv, _ = NewBaseConvertor(62)
var errInvalidLink = errors.New("Short link too large")

// Link describes a link in the database
type Link struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	URL        string        `bson:"url" json:"url"`
	Clicks     uint          `bson:"clicks" json:"clicks"`
	ShortID    string        `bson:"short_id" json:"short_id"`
	ShortIDInt uint          `bson:"short_id_int" json:"short_id_int"`
	UpdatedAt  *time.Time    `bson:"updated_at,omitempty" json:"updated_at"`
	CreatedAt  *time.Time    `bson:"created_at,omitempty" json:"created_at"`
}

// Counter has the info about total counter so far
type Counter struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Count     uint          `bson:"count" json:"count"`
	UpdatedAt *time.Time    `bson:"updated_at,omitempty" json:"updated_at"`
	CreatedAt *time.Time    `bson:"created_at,omitempty" json:"created_at"`
	StatType  string        `bson:"stat_type" json:"stat_type"`
}

func newLink(URL string) *Link {
	var timenow = time.Now().UTC()
	var counter = incrementGlobalCounter()
	var shortID = baseconv.Encode(counter)

	link := Link{
		ID:         bson.NewObjectId(),
		URL:        URL,
		Clicks:     0,
		ShortID:    shortID,
		ShortIDInt: counter,
		CreatedAt:  &timenow,
		UpdatedAt:  &timenow,
	}

	coll := db.C(config.LinksColl)
	err := coll.Insert(link)
	if err != nil {
		log.Fatal("Error inserting new record", err)
	}

	return &link
}

func getLinkForShortID(shortID string) (*Link, error) {
	if len(shortID) > 10 {
		return nil, errInvalidLink
	}
	coll := db.C(config.LinksColl)
	shortIDInt := baseconv.Decode(shortID)
	var link Link
	err := coll.Find(bson.M{"short_id_int": shortIDInt}).One(&link)
	if err != nil && err != mgo.ErrNotFound {
		log.Fatal("Error while finding link", err)
	}

	return &link, err
}

func incrementLinkCounter(link *Link) {
	var timenow = time.Now().UTC()
	coll := db.C(config.LinksColl)
	link.Clicks++
	link.UpdatedAt = &timenow
	err := coll.UpdateId(link.ID, link)
	if err != nil {
		// log.Println("error here", err)
		log.Fatal("Error incrementing counter for link", link, err)
	}
}

func incrementGlobalCounter() uint {
	var (
		counter Counter
		err     error
		// info    *mgo.ChangeInfo
		timenow = time.Now().UTC()
	)

	coll := db.C(config.CounterColl)

	err = coll.Find(nil).One(&counter)
	if err != nil && err != mgo.ErrNotFound {
		log.Fatal("Error getting counter from database", err)
	}

	if err == mgo.ErrNotFound {
		counter.ID = bson.NewObjectId()
		counter.Count = 0
		counter.CreatedAt = &timenow
		counter.StatType = "counter"
	}

	counter.Count++
	counter.UpdatedAt = &timenow

	_, err = coll.Upsert(bson.M{"stat_type": "counter"}, counter)
	if err != nil {
		log.Fatal("Failed updating counter", err)
	}

	return counter.Count
}
