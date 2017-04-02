package main

import (
  log "github.com/Sirupsen/logrus"
  _ "github.com/go-sql-driver/mysql"
  "os"
  "time"
)

func main() {
  initLogging()

  db := mustSetupDB()

  minId := 0
  limit := 500000
  for {
    then := time.Now()
    ts, err := db.fetch(minId, limit)
    elapsed := time.Now().Sub(then)
    if err != nil {
      log.WithError(err).Error("unexpected error")
      os.Exit(1)
    }
    log.WithField("min-id", minId).WithField("duration", elapsed).Infof("fetched %d transactions", len(ts))

    if len(ts) < limit {
      minId = 0
    } else {
      minId = ts[len(ts)-1].Id
    }
  }
}

func initLogging() {
  log.SetLevel(log.InfoLevel)
}
