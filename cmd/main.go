package main

import (
	"os"

	"github.com/eure/bobo"
	bobocommand "github.com/eure/bobo/command"
	"github.com/eure/bobo/engine/slack"
	"github.com/eure/bobo/log"
	"github.com/evalphobia/bobo-experiment/experiment/aws"
	"github.com/evalphobia/bobo-experiment/experiment/faceplusplus"
	"github.com/evalphobia/bobo-experiment/experiment/google"

	"github.com/evalphobia/kentykenty-bot/command"
)

func main() {
	bobo.Run(bobo.RunOption{
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
			aws.AWSCostCommand{},
			google.CalendarCommand,
			google.WhereCommand,
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
