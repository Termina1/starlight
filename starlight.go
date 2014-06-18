package main

import(
  "github.com/stathat/jconfig"
  "github.com/golang/glog"
  "flag"
  "os"
  "os/signal"
  "syscall"
)

func main() {
  flag.Parse()
  config := jconfig.LoadConfig("config.json")
  flag.Set("log_dir", config.GetString("logDir"))
  flag.Set("stderrthreshold", "ERROR")
  glog.Infoln("Config loaded successgully")
  mconf := MongoConf{
    config.GetString("localhost"),
    config.GetString("mongoDb"),
    config.GetString("mongoCollection"),
    config.GetInt("mongoBatch"),
    config.GetInt("mongoReindex"),
  }
  rconf := RedisConf{
    config.GetString("redisHost"),
    uint(config.GetInt("redisPort")),
    config.GetString("redisQueue"),
  }
  token := config.GetString("token")
  density := config.GetInt("beamDensity")
  beam := CreateStarBeam(StarExtractor(mconf, token), mconf, rconf, density)
  beam.launch()
  glog.Infoln("Starlight initialized")
  termination := make(chan os.Signal, 1)
  signal.Notify(termination, os.Interrupt)
  signal.Notify(termination, syscall.SIGTERM)
  <-termination
  glog.Flush()
}