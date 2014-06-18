package main

import (
  "labix.org/v2/mgo"
  "time"
  "github.com/golang/glog"
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
  session, error := mgo.Dial(conf.Url)
  if error != nil {
    glog.Fatalln("Couldn't establish connection with MongoDB: ", error)
  }
  glog.Info("Established connection with MongoDB")
  return session.DB(conf.Db).C(conf.Collection)
}

func SubscribeMongo(iter *mgo.Iter) <-chan string {
  tube := make(chan string)
  var repo StarRepo
  go func() {
    glog.Info("Start iterating over existing repos")
    for iter.Next(&repo) == true {
      tube <- repo.Name
    }
    close(tube)
    glog.Info("Iterated over all repos, closing channel")
  }()
  return tube
}

func LaunchSubscription(coll *mgo.Collection, batch, reindexTime int, out chan string) {
  glog.Info("Start reindexing all repositories")
  iter := ReindexRepos(coll, batch)
  source := SubscribeMongo(iter)
  go func() {
    for val := range source  {
      out <- val
    }
    glog.Info("Indexation finished")
    time.AfterFunc(time.Minute * time.Duration(reindexTime), func() {
      LaunchSubscription(coll, batch, reindexTime, out)
    })
  }()
}

func CreateSubscriptionMongo(conf MongoConf) <-chan string {
  coll := CreateMongoClient(conf)
  out := make(chan string)
  time.AfterFunc(time.Minute * time.Duration(conf.Time), func() {
    LaunchSubscription(coll, conf.Batch, conf.Time, out)
  })
  return out
}