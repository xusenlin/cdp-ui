package tasks

import (
	"context"
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"math"
)

//截屏，全屏或者截取某个元素
func (t *Task) screenshot(args taskArgs) error {
	var expectArgs struct {
		Sel      string
		FullPage bool
		FileName string `valid:"required"`
	}
	if err := mapstructure.Decode(args, &expectArgs); err != nil {
		return errors.New("screenshot子任务出错:" + err.Error())
	}

	if _, err := govalidator.ValidateStruct(&expectArgs); err != nil {
		return errors.New("screenshot子任务参数出错:" + err.Error())
	}

	var err error
	var file []byte

	if expectArgs.FullPage {
		err = chromedp.Run(*t.ctx, fullPage(100, &file))
	} else {
		err = chromedp.Run(*t.ctx, element(expectArgs.Sel, &file))
	}
	if err != nil {
		return errors.New("screenshot子任务出错:" + err.Error())
	}
	if err := ioutil.WriteFile(t.workDir()+"/"+expectArgs.FileName, file, 0666); err != nil {
		return err
	}
	return nil
}

//截取当前视窗
func (t *Task) captureScreenshot(args taskArgs) error {
	var expectArgs struct {
		FileName string `valid:"required"`
	}
	if err := mapstructure.Decode(args, &expectArgs); err != nil {
		return errors.New("captureScreenshot子任务出错:" + err.Error())
	}

	if _, err := govalidator.ValidateStruct(&expectArgs); err != nil {
		return errors.New("captureScreenshot子任务参数出错:" + err.Error())
	}

	var file []byte
	err := chromedp.Run(*t.ctx, chromedp.CaptureScreenshot(&file))

	if err != nil {
		return errors.New("captureScreenshot子任务出错:" + err.Error())
	}
	if err := ioutil.WriteFile(t.workDir()+"/"+expectArgs.FileName, file, 0666); err != nil {
		return err
	}

	return nil
}

func element(sel string, res *[]byte) chromedp.Tasks {

	return chromedp.Tasks{
		chromedp.WaitVisible(sel),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	}
}

func fullPage(quality int64, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			*res, err = page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}
