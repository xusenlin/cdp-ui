package js

import (
	"cdpui/config"
	"cdpui/helper"
	"cdpui/tasks"
	"github.com/zserge/lorca"
	"io/ioutil"
	"os"
	"strconv"
)
var UI lorca.UI
var Tasks = make(map[int]*tasks.Task) //正在运行的任务

func Bind(ui *lorca.UI) error {
	UI = *ui
	if err := UI.Bind("RunTask", runTask); err != nil {
		return err
	}
	if err := UI.Bind("StopTask", stopTask); err != nil {
		return err
	}
	if err := UI.Bind("OpenFile", openFile); err != nil {
		return err
	}
	if err := UI.Bind("SaveJsonFile", saveJsonFile); err != nil {
		return err
	}
	if err := UI.Bind("GetJsonTask", getJsonTask); err != nil {
		return err
	}
	return nil
}

func runTask(t string) {
	task, err := tasks.New([]byte(t))
	if err != nil {
		UI.Eval(notifyError(err.Error()))
		return
	}
	_, exist := Tasks[task.Index]
	if exist {
		UI.Eval(notifyError("之前的任务还在运行"))
		return
	}

	task.BindCallBack(func(progress int) {
		UI.Eval(updateProgress(task.Index, progress))
	})
	Tasks[task.Index] = task

	go func() {
		UI.Eval(notifyInfo("开始执行任务" + task.Title))
		err = Tasks[task.Index].Run()
		if err != nil {
			UI.Eval(notifyError(err.Error()))
			UI.Eval(updateProgress(task.Index, -1))
			delete(Tasks, task.Index)
			return
		}
		UI.Eval(notifySuccess("任务执行完毕"))
		UI.Eval(updateProgress(task.Index, -1))
		delete(Tasks, task.Index)
	}()
}

func stopTask(i string)  {
	index, err := strconv.Atoi(i)
	if err != nil {
		UI.Eval(notifyError(err.Error()))
		return
	}
	task, exist := Tasks[index]

	if !exist {
		UI.Eval(notifyError("此任务不存在"))
		return
	}
	select {
	case <-task.Done():
		UI.Eval(updateProgress(task.Index, -1))
		UI.Eval(notifySuccess("此任务已经完成"))
	default:
		task.Cancel()
		UI.Eval(updateProgress(task.Index, -1))
		UI.Eval(notifySuccess("此任务已经暂停"))
	}
}

func openFile(i string)  {
	index, err := strconv.Atoi(i)

	if err != nil {
		UI.Eval(notifyError(err.Error()))
		return
	}

	if !helper.IsDir(helper.GetWorkDir(index)) {
		UI.Eval(notifyWarning("此任务还没有产生文件"))
		return
	}

	if err := helper.OpenFileManager(index); err != nil {
		UI.Eval(notifyError(err.Error()))
		return
	}
}

func saveJsonFile(t string)  {
	if err := helper.RemoveFile(config.TaskFile);err != nil {
		UI.Eval(notifyError(err.Error()))
		return
	}
	err := ioutil.WriteFile(config.TaskFile, []byte(t), os.ModeAppend)
	if err != nil {
		UI.Eval(notifyError(err.Error()))
		return
	}
	UI.Eval(notifySuccess("任务保存成功"))
}

func getJsonTask() string {
	if helper.FileIsExisted(config.TaskFile){
		b, err := ioutil.ReadFile(config.TaskFile)
		if err != nil {
			UI.Eval(notifyError(err.Error()))
			return "[]"
		}
		return string(b)
	}
	return "[]"
}

