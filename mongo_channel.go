package main

import "labix.org/v2/mgo"

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
  }()
  return tube
}

func CreateSubscriptionMongo(url string, batch int) <-chan string {
  sess := CreateMongoClient(url)
  iter := ReindexRepos(sess, batch)
  return SubscribeMongo(iter)
}