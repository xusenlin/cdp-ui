package js


import "fmt"

func notice(s string, t string) string {
	return fmt.Sprintf("window.Notice(`%s`,`%s`)", s, t)
}

func notifySuccess(s string) string {
	return notice(s, "success")
}
func notifyWarning(s string) string {
	return notice(s, "warning")
}
func notifyInfo(s string) string {
	return notice(s, "info")
}
func notifyError(s string) string {
	return notice(s, "error")
}

func updateProgress(taskIndex int, progress int) string {
	return fmt.Sprintf(`window.UpdateProgress("%v","%v")`, taskIndex, progress)
}