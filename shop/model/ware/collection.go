/*
 * Revision History:
 *     Initial: 2018/04/27        Chen Yanchen
 */

package ware

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (TableWare = "ware")
type Collection struct {
	Id  bson.ObjectId `bson:"_id,omitempty"`
	Uid bson.ObjectId `bson:"uid"` // user id
	Wid bson.ObjectId `bson:"wid"` // ware id
	CreatedAt time.Time`bson:"createdAt"`
}