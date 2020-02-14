package main

import (
	"os"
	"time"

	"github.com/eure/bobo"
	bobocommand "github.com/eure/bobo/command"
	"github.com/eure/bobo/engine/slack"
	"github.com/eure/bobo/log"
	"github.com/evalphobia/bobo-experiment/experiment/aws"
	"github.com/evalphobia/bobo-experiment/experiment/faceplusplus"
	"github.com/evalphobia/bobo-experiment/experiment/google"
	"github.com/jpillora/overseer/fetcher"

	"github.com/evalphobia/kentykenty-bot/command"
)

func main() {

	bobo.Run(bobo.RunOption{
		UseUpgrade: true,
		UpgradeFetcher: &bobo.MultiFetcher{
			List: []fetcher.Interface{
				&fetcher.File{
					Path: getLocalBinPath(),
				},
				&fetcher.Github{
					User:     "evalphobia",
					Repo:     "kentykenty-bot",
					Interval: 5 * time.Minute,
				},
			},
		},
		Engine: &slack.SlackEngine{},
		Logger: &log.StdLogger{
			IsDebug: bobo.IsDebug(),
		},
		CommandSet: bobocommand.NewCommandSet(
			bobocommand.PingCommand,
			bobocommand.ParrotCommand,
			bobocommand.GoodMorningCommand,
			bobocommand.ReloadCommand,
			bobocommand.HelpCommand,
			aws.CostCommandByCostExplorer{},
			aws.SQSCommand{},
			aws.DynamoDBCommand{},
			google.CalendarCommand,
			google.WhereCommand,
			&google.RoomCommand{},
			&faceplusplus.MergeCommand{},
			command.MergeYui,
			command.MergeAoi,
			command.MergeRizap,
			command.MergeKanada,
			&command.ReactEmojiCommand{
				HighProbList: []string{
					"takuma.morikawa",
					"hiroki.kojima",
					"kento.yamashita",
				},
				Blacklist: []string{
					"kentaro.takahashi",
				},
			},
			&command.TFQueryCommand{
				Endpoint: os.Getenv("BOBO_TFQ_URL"),
				Command:  os.Getenv("BOBO_TFQ_COMMAND"),
			},
		),
	})
}

func getLocalBinPath() string {
	path, _ := os.Executable()
	return path
}
