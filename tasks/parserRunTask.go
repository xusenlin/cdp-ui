package tasks

import "errors"

func (t *Task) parserRun() error {
	for progress, task := range t.Tasks {
		t.progress = progress
		t.callBack(progress)
		if err := t.selectFunc(task.Name, task.Args); err != nil {
			return err
		}
	}
	return nil
}

func (t *Task) selectFunc(taskName string, taskArgs taskArgs) error {
	switch taskName {
	case "pdf":
		return t.pdf(taskArgs)
	case "click":
		return t.click(taskArgs)
	case "sleep":
		return t.sleep(taskArgs)
	case "evaluate":
		return t.evaluate(taskArgs)
	case "navigate":
		return t.navigate(taskArgs)
	case "setValue":
		return t.setValue(taskArgs)
	case "screenshot":
		return t.screenshot(taskArgs)
	case "addCookies":
		return t.addCookies(taskArgs)
	case "waitVisible":
		return t.waitVisible(taskArgs)
	case "addReqHeaders":
		return t.addReqHeaders(taskArgs)
	case "emulateViewport":
		return t.emulateViewport(taskArgs)
	case "captureScreenshot":
		return t.captureScreenshot(taskArgs)
	default:
		return errors.New("未识别的任务名字" + taskName)
	}
}
