package main

import (
  "labix.org/v2/mgo"
  "time"
)
var counting int = 0

func CreateMongoClient(url string) *mgo.Session {
  session, _ := mgo.Dial(url)
  return session
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

func LaunchSubscription(sess *mgo.Session, batch int, out chan string) {
  iter := ReindexRepos(sess, batch)
  source := SubscribeMongo(iter)
  go func() {
    for val := range source  {
      out <- val
    }
    time.AfterFunc(time.Minute, func() {
      LaunchSubscription(sess, batch, out)
    })
  }()
}

func CreateSubscriptionMongo(url string, batch int) <-chan string {
  sess := CreateMongoClient(url)
  out := make(chan string)
  LaunchSubscription(sess, batch, out)
  return out
}