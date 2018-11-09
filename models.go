package main

import (
	"time"
)

// Link describes a link in the database
type Link struct {
	ID         string     `bson:"_id" json:"id"`
	URL        string     `bson:"url" json:"url"`
	Clicks     uint       `bson:"clicks" json:"clicks"`
	ShortID    string     `bson:"short_id" json:"short_id"`
	ShortIDInt uint       `bson:"short_id_int" json:"short_id_int" sql:",unique"`
	UpdatedAt  *time.Time `bson:"updated_at,omitempty" json:"updated_at"`
	CreatedAt  *time.Time `bson:"created_at,omitempty" json:"created_at"`
}

// Counter has the info about total counter so far
type Counter struct {
	ID        string     `bson:"_id" json:"id"`
	Count     uint       `bson:"count" json:"count"`
	UpdatedAt *time.Time `bson:"updated_at,omitempty" json:"updated_at"`
	CreatedAt *time.Time `bson:"created_at,omitempty" json:"created_at"`
	StatType  string     `bson:"stat_type" json:"stat_type"`
}
