package command

import (
	"fmt"
	"math/rand"
	"regexp"
	"sync"
	"time"

	"github.com/eure/bobo/command"
)

var _ command.CommandTemplate = &ReactEmojiCommand{}

var reReactEmoji = regexp.MustCompile("^")

const defaultProbability = 15
const defaultHighProbability = 40

type ReactEmojiCommand struct {
	HighProbList []string
	Blacklist    []string

	HighProbability    int
	DefaultProbability int

	listOnce     sync.Once
	highproblist map[string]struct{}
	blacklist    map[string]struct{}
}

func (ReactEmojiCommand) GetMentionCommand() string {
	return ""
}

func (ReactEmojiCommand) GetHelp() string {
	return "Add reaction to message"
}

func (ReactEmojiCommand) HasHelp() bool {
	return true
}

func (ReactEmojiCommand) GetRegexp() *regexp.Regexp {
	return reReactEmoji
}

func (r *ReactEmojiCommand) Exec(d command.CommandData) {
	r.init()
	c := r.runReactEmoji(d)
	c.Exec()
}

func (r *ReactEmojiCommand) init() {
	r.listOnce.Do(func() {
		r.highproblist = make(map[string]struct{})
		for _, s := range r.HighProbList {
			r.highproblist[s] = struct{}{}
		}
		r.blacklist = make(map[string]struct{})
		for _, s := range r.Blacklist {
			r.blacklist[s] = struct{}{}
		}
	})
}

func (r ReactEmojiCommand) runReactEmoji(d command.CommandData) command.Command {
	c := command.Command{}
	if r.isInBlacklist(d.SenderName) {
		return c
	}

	prob := r.getProbability(d.SenderName)
	if !isRandValid(prob) {
		return c
	}

	emoji, err := d.Engine.GetEmojiByRandom()
	if err != nil {
		errMessage := fmt.Sprintf("[ERROR]\t[GetEmojiByRandom]\t`%s`", err.Error())
		task := command.NewReplyEngineTask(d.Engine, d.Channel, errMessage)
		c.Add(task)
		return c
	}

	task := command.NewReactionEmojiEngineTask(d.Engine, d.Channel, emoji, d.ThreadTimestamp)
	c.Add(task)
	return c
}

func (r *ReactEmojiCommand) getProbability(name string) int {
	switch {
	case r.isInHighProbList(name):
		if r.HighProbability > 0 {
			return r.HighProbability
		}
		return defaultHighProbability
	default:
		if r.DefaultProbability > 0 {
			return r.DefaultProbability
		}
		return defaultProbability
	}
}

func (r *ReactEmojiCommand) isInBlacklist(name string) bool {
	_, ok := r.blacklist[name]
	return ok
}

func (r *ReactEmojiCommand) isInHighProbList(name string) bool {
	_, ok := r.highproblist[name]
	return ok
}

func isRandValid(percent int) bool {
	rand.Seed(time.Now().UTC().UnixNano())
	return percent > rand.Intn(100)
}
