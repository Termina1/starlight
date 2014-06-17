package main

import(
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)

const DBName = "starlight"

const Collection = "repos"

type StarRepo struct {
  Name string
  Indexed bool
}

func ReindexRepos(session *mgo.Session, batch int) *mgo.Iter {
  coll := session.DB(DBName).C(Collection)
  return coll.Find(bson.M{}).Batch(batch).Iter()
}