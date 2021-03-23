package tasks

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/mitchellh/mapstructure"
)

func (t *Task) addReqHeaders(args taskArgs) error {
	var expectArgs struct {
		Headers taskArgs `valid:"-"`
	}

	if err := mapstructure.Decode(args, &expectArgs); err != nil {
		return errors.New("setReqHeaders子任务出错:" + err.Error())
	}
	if _, err := govalidator.ValidateStruct(&expectArgs); err != nil {
		return errors.New("setReqHeaders子任务参数出错:" + err.Error())
	}
	return chromedp.Run(*t.ctx,
		network.Enable(),
		network.SetExtraHTTPHeaders(network.Headers(expectArgs.Headers)),
	)
}
