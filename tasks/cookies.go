package tasks

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/mitchellh/mapstructure"
)

func (t *Task) addCookies(args taskArgs) error {
	var expectArgs struct {
		Cookies []*network.CookieParam `valid:"-"`
	}

	if err := mapstructure.Decode(args, &expectArgs); err != nil {
		return errors.New("addCookies子任务出错:" + err.Error())
	}

	if _, err := govalidator.ValidateStruct(&expectArgs); err != nil {
		return errors.New("addCookies子任务参数出错:" + err.Error())
	}
	return chromedp.Run(*t.ctx, network.SetCookies(expectArgs.Cookies), )
}
