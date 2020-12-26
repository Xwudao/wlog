package wlog

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"

	"path/filepath"
)

var wLog *WLog

type WLog struct {
	logger   *logrus.Logger
	path     string
	filename string
	leastDay uint
}

func New() *WLog {
	return &WLog{logger: logrus.New()}
}

func NewWLog(logger *logrus.Logger, path string, filename string, leastDay uint) *WLog {
	return &WLog{logger: logger, path: path, filename: filename, leastDay: leastDay}
}

func init() {
	wLog = New()
	wLog.leastDay = 10
	p, err := GetCurrentPath()
	if err != nil {
		p, _ = os.Getwd()
	}
	wLog.path = p
	wLog.filename = "log"
	// add hooks
	wLog.logger.AddHook(wLog.NewLfsHook())
	// caller
	wLog.logger.SetReportCaller(true)
}

func (w *WLog) NewLfsHook() logrus.Hook {
	writer, err := rotatelogs.New(
		// 日志文件
		filepath.Join(w.path, w.filename)+".%Y-%m-%d.log",
		// 日志周期(默认每86400秒/一天旋转一次)
		//rotatelogs.WithRotationTime(rotationTime),
		// 清除历史 (WithMaxAge和WithRotationCount只能选其一)
		//rotatelogs.WithMaxAge(time.Hour*24*7), //默认每7天清除下日志文件
		rotatelogs.WithRotationCount(w.leastDay), //只保留最近的N个日志文件
		//rotatelogs.WithLinkName(filepath.Join(logPath, "current.log")),
	)
	if err != nil {
		panic(err)
	}
	// 可设置按不同level创建不同的文件名
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
		//}, &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	}, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return lfsHook
}
