package tasks

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/chromedp/chromedp"
	"github.com/mitchellh/mapstructure"
	"time"
)

//导航
func (t *Task) navigate(args taskArgs) error {

	var expectArgs struct {
		Host string `valid:"requrl"`
	}

	if err := mapstructure.Decode(args, &expectArgs); err != nil {
		return errors.New("navigate子任务出错:" + err.Error())
	}

	if _, err := govalidator.ValidateStruct(&expectArgs); err != nil {
		return errors.New("navigate子任务参数出错:" + err.Error())
	}
	return chromedp.Run(*t.ctx, chromedp.Navigate(expectArgs.Host))
}

//调整浏览器大小
func (t *Task) emulateViewport(args taskArgs) error {

	var expectArgs struct {
		Width  int64 `valid:"required"`
		Height int64 `valid:"required"`
	}

	if err := mapstructure.Decode(args, &expectArgs); err != nil {
		return errors.New("emulateViewport子任务出错:" + err.Error())
	}

	if _, err := govalidator.ValidateStruct(&expectArgs); err != nil {
		return errors.New("emulateViewport子任务参数出错:" + err.Error())
	}
	return chromedp.Run(*t.ctx, chromedp.EmulateViewport(expectArgs.Width, expectArgs.Height))
}

//等待元素加载
func (t *Task) waitVisible(args taskArgs) error {
	var expectArgs struct {
		Sel string `valid:"required"`
	}

	if err := mapstructure.Decode(args, &expectArgs); err != nil {
		return errors.New("waitVisible子任务出错:" + err.Error())
	}

	if _, err := govalidator.ValidateStruct(&expectArgs); err != nil {
		return errors.New("waitVisible子任务参数出错:" + err.Error())
	}
	return chromedp.Run(*t.ctx, chromedp.WaitVisible(expectArgs.Sel))
}

//输入内容,覆盖之前的
func (t *Task) setValue (args taskArgs) error {
	var expectArgs struct {
		Sel string `valid:"required"`
		Val string `valid:"required"`
	}

	if err := mapstructure.Decode(args, &expectArgs); err != nil {
		return errors.New("sendKeys子任务出错:" + err.Error())
	}

	if _, err := govalidator.ValidateStruct(&expectArgs); err != nil {
		return errors.New("sendKeys子任务参数出错:" + err.Error())
	}
	return chromedp.Run(*t.ctx, chromedp.SetValue(expectArgs.Sel, expectArgs.Val))
}

//点击元素
func (t *Task) click(args taskArgs) error {
	var expectArgs struct {
		Sel string `valid:"required"`
	}

	if err := mapstructure.Decode(args, &expectArgs); err != nil {
		return errors.New("click子任务出错:" + err.Error())
	}

	if _, err := govalidator.ValidateStruct(&expectArgs); err != nil {
		return errors.New("click子任务参数出错:" + err.Error())
	}
	return chromedp.Run(*t.ctx, chromedp.Click(expectArgs.Sel, chromedp.NodeVisible))
}
//休眠
func (t *Task) sleep(args taskArgs) error {
	var expectArgs struct {
		Second time.Duration `valid:"required"`
	}

	if err := mapstructure.Decode(args, &expectArgs); err != nil {
		return errors.New("sleep子任务出错:" + err.Error())
	}

	if _, err := govalidator.ValidateStruct(&expectArgs); err != nil {
		return errors.New("sleep子任务参数出错:" + err.Error())
	}
	return chromedp.Run(*t.ctx, chromedp.Sleep(expectArgs.Second * time.Second))
}
