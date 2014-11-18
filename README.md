Starlight
=========
[![Build Status](https://drone.io/github.com/Termina1/starlight/status.png)](https://drone.io/github.com/Termina1/starlight/latest)
[ ![Codeship Status for Termina1/starlight](https://codeship.com/projects/6b4efa10-5187-0132-d013-6a4a03e15b82/status)](https://codeship.com/projects/48374)

Simple github repo indexer written in Go.

Built on top of [http://www.mongodb.org/](MongoDB) for persistence and [http://redis.io/](Redis) for Pub/Sub.

This whole project was built as my first experiment with Go.

## Build

To build you should run

```
go get labix.org/v2/mgo
go get github.com/golang/glog
go get menteslibres.net/gosexy/redis
go get code.google.com/p/goauth2/oauth
go get github.com/google/go-github/github
go get github.com/stathat/jconfig
go build
```

## Config

Config should be stored in config.json. There is an example-config.json as config example.

```javascript
{
  "mongoUrl": "localhost", //path to MongoDB
  "mongoDb": "db", //db where indexd repos should be stored
  "mongoCollection": "collection", //collection where starred repos should be stored
  "mongoBatch": 100, //batch size to select repos from mongo (while reindexing db)
  "mongoReindex": 1440, //how often app should reindex db
  "redisHost": "localhost", //redis host
  "redisPort": 6379, //redis port
  "redisQueue": "queue", //redis Pub/Sub channel 
  "token": "github token", //your github token for API access
  "beamDensity": 10, //how many workers starlight should use (amount of concurrently indexed repos)
  "logDir": "logs" //where to store logs
}
```
