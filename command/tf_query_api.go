package command

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"sync"

	"github.com/eure/bobo/command"
	"github.com/evalphobia/httpwrapper/request"
)

type TFQueryCommand struct {
	Endpoint string
	Command  string
}

func (t TFQueryCommand) GetMentionCommand() string {
	return t.Command
}

func (TFQueryCommand) GetHelp() string {
	return "Send TF Query and get result text"
}

func (TFQueryCommand) HasHelp() bool {
	return true
}

func (TFQueryCommand) GetRegexp() *regexp.Regexp {
	return nil
}

func (t *TFQueryCommand) Exec(d command.CommandData) {
	c := t.runQuery(d)
	c.Exec()
}

func (t *TFQueryCommand) runQuery(d command.CommandData) command.Command {
	c := command.Command{}

	text := strings.TrimSpace(d.TextOther)
	texts := strings.Split(text, " ")

	command.NewReplyEngineTask(d.Engine, d.Channel, "考え中...").Run()

	wg := &sync.WaitGroup{}
	result := &TFQueryResult{}
	for _, v := range texts {
		wg.Add(1)
		go func(keyword string) {
			defer wg.Done()
			input := fmt.Sprintf("%s%s", keyword, getRandKey())
			uri := fmt.Sprintf("%s?input=%s", t.Endpoint, input)
			resp, err := request.GET(uri, request.Option{})
			switch {
			case err != nil,
				!resp.Ok:
				return
			}
			outputs := TFQueryOutputs{}
			if err := resp.JSON(&outputs); err != nil {
				return
			}
			if outputs.HasValue() {
				result.Add(outputs.GetRandomVal())
			}
		}(v)
	}
	wg.Wait()
	if len(result.list) == 0 {
		task := command.NewReplyEngineTask(d.Engine, d.Channel, "何も考えられません")
		c.Add(task)
		return c
	}

	task := command.NewReplyEngineTask(d.Engine, d.Channel, result.Show())
	c.Add(task)
	return c
}

type TFQueryResult struct {
	mu   sync.Mutex
	list []string
}

func (r *TFQueryResult) Add(text string) {
	text = addDesu(text)
	r.mu.Lock()
	r.list = append(r.list, text)
	r.mu.Unlock()
}

func hasDesu(text string) bool {
	switch {
	case strings.HasSuffix(text, "です。"),
		strings.HasSuffix(text, "した。"),
		strings.HasSuffix(text, "ます。"):
		return true
	}
	return false
}

func addDesu(text string) string {
	switch {
	case strings.HasSuffix(text, "。"),
		strings.HasSuffix(text, "、"):
	default:
		text = text + "。"
	}
	if hasDesu(text) {
		return text
	}

	return strings.TrimSuffix(text, "。") + "です。"
}

func (r TFQueryResult) Show() string {
	return fmt.Sprintf("```\n%s\n```", strings.Join(r.list, ""))
}

type TFQueryOutputs struct {
	Outputs []TFQueryOutput `json:"outputs"`
}

func (o TFQueryOutputs) HasValue() bool {
	return len(o.Outputs) != 0 && len(o.Outputs[0].Val) != 0
}

func (o TFQueryOutputs) GetRandomVal() string {
	return o.Outputs[0].GetRandomVal()
}

type TFQueryOutput struct {
	Score []float64 `json:"score"`
	Val   []string  `json:"val"`
}

func (o TFQueryOutput) GetRandomVal() string {
	return o.Val[rand.Intn(len(o.Val)-1)]
}

func getRandKey() string {
	const randKeys = " 　abcdefghijklmnopqrstuvwxyz./<>';:[]`~%^*_-+="
	return string(randKeys[rand.Intn(len(randKeys))])
}
