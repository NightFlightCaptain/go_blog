package logging

import (
	"fmt"
	"go_blog/pkg/setting"
	"time"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.Config.App.RuntimeRootPath, setting.Config.App.LogSavePath)
}

func getLogFileName() string {
	suffixPath := fmt.Sprintf("%s%s.%s",
		setting.Config.App.LogSaveName,
		time.Now().Format(setting.Config.App.TimeFormat),
		setting.Config.App.LogSaveExt)
	return fmt.Sprintf("%s", suffixPath)
}
