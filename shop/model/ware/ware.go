/*
 * Revision History:
 *     Initial: 2018/04/27        Chen Yanchen
 */

package ware

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Ware struct {
	Id      bson.ObjectId `bson:"_id,omitempty"`
	Sid     bson.ObjectId `bson:"sid"`
	Name    string        `bson:"name"`
	Brief   string        `bson:"brief"`
	Price   float64       `bson:"price"`
	Stock   int           `bson:"stock"`
	Imgs    []string      `bson:"imgs"`
	Detail  string        `bson:"detail"` // save as .md file
	Created time.Time     `bson:"created"`
	State   bool          `bson:"state"`
}
