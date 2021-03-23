package tasks

import (
	"cdpui/helper"
	"context"
	"encoding/json"
	"errors"
	"github.com/chromedp/chromedp"
	"time"
)

type taskArgs map[string]interface{}
type flagArgs map[string]interface{}

type Task struct {
	Index       int           `json:"index"`
	Title       string        `json:"title"`
	Tasks       []taskDesc    `json:"tasks"`
	Flags       []flagArgs    `json:"flags"`
	Width       int           `json:"width"`
	Height      int           `json:"height"`
	TimeOut     time.Duration `json:"timeOut"`
	UserAgent   string        `json:"userAgent"`
	ShowBrowser bool          `json:"showBrowser"`
	//////
	progress int
	callBack func(int)
	ctx      *context.Context
	cancel   context.CancelFunc
}

type taskDesc struct {
	Name string   `json:"name"`
	Args taskArgs `json:"args"`
}

func New(tasksDesc []byte) (*Task, error) {
	var task Task

	if err := json.Unmarshal(tasksDesc, &task); err != nil {
		return nil, errors.New("任务Json解析出错：" + err.Error())
	}

	if task.Width == 0 {
		task.Width = 1366
	}
	if task.Height == 0 {
		task.Height = 768
	}
	if task.UserAgent == "" {
		task.UserAgent = `Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`
	}

	return &task, nil
}

//目前还不支持操作打开新的TAB

func (t *Task) Run() error {

	if err := t.initDir(); err != nil {
		return err
	}

	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", !t.ShowBrowser),
		chromedp.WindowSize(t.Width, t.Width),
		chromedp.UserAgent(t.UserAgent),

	}
	//chromedp.Flag("hide-scrollbars", false),
	//chromedp.Flag("mute-audio", false),
	//chromedp.Flag("mute-audio", false),
	//chromedp.Flag("ignore-certificate-errors", "1"),
	for k, v := range t.Flags {
		options = append(options, chromedp.Flag(string(k), v))
	}

	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, cc := chromedp.NewExecAllocator(context.Background(), options...)
	defer cc()

	ctx, cancel := chromedp.NewContext(c,chromedp.WithLogf(func(s string, i ...interface{}) {

	}))
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, t.TimeOut*time.Second)
	defer cancel()

	t.ctx = &ctx
	t.cancel = cc

	return t.parserRun()
}

func (t *Task) BindCallBack(f func(int)) {
	t.callBack = f
}

func (t *Task) Done() <-chan struct{} {
	return (*t.ctx).Done()
}
func (t *Task) Cancel() {
	t.cancel()
}
func (t *Task) workDir() string {
	return helper.GetWorkDir(t.Index)
}

func (t *Task) initDir() error {

	dir := t.workDir()

	if err := helper.RemoveDir(dir); err != nil {
		return err
	}

	if err := helper.MakeDir(dir); err != nil {
		return err
	}

	return nil
}
