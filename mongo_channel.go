package main

import (
  "labix.org/v2/mgo"
  "time"
)
var counting int = 0

type MongoConf struct {
  Url string
  Db string
  Collection string
  Batch int
  Time int
}

func CreateMongoClient(conf MongoConf) *mgo.Collection {
  session, _ := mgo.Dial(conf.Url)
  return session.DB(conf.Db).C(conf.Collection)
}

func SubscribeMongo(iter *mgo.Iter) <-chan string {
  tube := make(chan string)
  var repo StarRepo
  go func() {
    for iter.Next(&repo) == true {
      tube <- repo.Name
    }
    close(tube)
  }()
  return tube
}

func LaunchSubscription(coll *mgo.Collection, batch, reindexTime int, out chan string) {
  iter := ReindexRepos(coll, batch)
  source := SubscribeMongo(iter)
  go func() {
    for val := range source  {
      out <- val
    }
    time.AfterFunc(time.Minute * time.Duration(reindexTime), func() {
      LaunchSubscription(coll, batch, reindexTime, out)
    })
  }()
}

func CreateSubscriptionMongo(conf MongoConf) <-chan string {
  coll := CreateMongoClient(conf)
  out := make(chan string)
  LaunchSubscription(coll, conf.Batch, conf.Time, out)
  return out
}