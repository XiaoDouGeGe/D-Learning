package utils

import (
	iris "github.com/kataras/iris/v12"
)

func LoggerInfo(msg interface{}) {
	iris.New().Logger().Info(msg)
}
