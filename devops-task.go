package main

import (
	"github.com/braior/brtool"
	"github.com/braior/devops-api/cmd"
	"github.com/braior/devops-api/utils"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
	"log"
	"time"
)

type dbBackup struct{}

func (d dbBackup) Run() {

	timeString := time.Now().Format("2006-01-02T15:04:05")
	tableName := viper.GetString("task.dbBackup.tableName")
	fileNamePrefix := viper.GetString("task.dbBackup.fileNamePrefix")
	btb, err := brtool.NewBoltDB(cmd.DBPath, tableName)
	if err != nil {
		utils.Logger.Error(nil, err.Error())
	}
	err = btb.Backup(fileNamePrefix + "_" + timeString)
	if err != nil {
		utils.Logger.Error(nil, err.Error())
	}
	log.Println("database backup succeed")
}

func main() {

	c := cron.New()
	spec := viper.GetString("task.dbBackup.spec")
	err := c.AddFunc(spec, func() {
		log.Println("database backup cron running...")
	})
	if err != nil {
		log.Println(err)
	}

	err = c.AddJob(spec, dbBackup{})
	if err != nil {
		log.Println(err)
	}
	c.Start()
	defer c.Stop()

	select {}
}
