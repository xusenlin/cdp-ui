package tasks

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/chromedp/chromedp"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
)

func (t *Task) evaluate(args taskArgs) error {
	var expectArgs struct {
		Script   string `valid:"required"`
		FileName string `valid:"required"`
	}
	if err := mapstructure.Decode(args, &expectArgs); err != nil {
		return errors.New("evaluate子任务出错:" + err.Error())
	}

	if _, err := govalidator.ValidateStruct(&expectArgs); err != nil {
		return errors.New("evaluate子任务参数出错:" + err.Error())
	}

	var file []byte
	err := chromedp.Run(*t.ctx, chromedp.Evaluate(expectArgs.Script, &file))

	if err != nil {
		return errors.New("evaluate子任务出错:" + err.Error())
	}
	if err := ioutil.WriteFile(t.workDir()+"/"+expectArgs.FileName, file, 0666); err != nil {
		return err
	}

	return nil
}
