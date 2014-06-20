package main

import(
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)

type StarRepo struct {
  Name *string
  Indexed bool
  StarGazers *int
  Forks *int
  Description *string
  SearchField *string
}

func ReindexRepos(coll *mgo.Collection, batch int) *mgo.Iter {
  return coll.Find(bson.M{}).Batch(batch).Iter()
}

func StarRepoUpdate(coll *mgo.Collection, name string, repo *StarRepo) {
  coll.Upsert(bson.M{"name": name}, repo)
}