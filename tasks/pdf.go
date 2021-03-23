package tasks

import (
	"context"
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
)

func (t *Task) pdf(args taskArgs) error {

	if t.ShowBrowser {
		return errors.New("打印网页PDF只支持在无头浏览器模式下运行")
	}

	var expectArgs struct {
		FileName string `valid:"required"`
	}
	if err := mapstructure.Decode(args, &expectArgs); err != nil {
		return errors.New("pdf子任务出错:" + err.Error())
	}

	if _, err := govalidator.ValidateStruct(&expectArgs); err != nil {
		return errors.New("pdf子任务参数出错:" + err.Error())
	}

	var res []byte
	err := chromedp.Run(*t.ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(false).Do(ctx)
			if err != nil {
				return err
			}
			res = buf
			return nil
		}),
	)
	if err != nil {
		return errors.New("pdf子任务出错(必须是无头浏览器才能正常运行):" + err.Error())
	}
	if err := ioutil.WriteFile(t.workDir()+"/"+expectArgs.FileName, res, 0666); err != nil {
		return err
	}

	return nil
}
