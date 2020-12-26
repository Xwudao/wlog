package wlog

import (
	"github.com/sirupsen/logrus"
)

func Trace(args ...interface{}) {
	wLog.logger.Trace(args...)
}

func Debug(args ...interface{}) {
	wLog.logger.Debug(args...)
}

func Print(args ...interface{}) {
	wLog.logger.Print(args...)
}

func Info(args ...interface{}) {
	wLog.logger.Info(args...)
}

func Warn(args ...interface{}) {
	wLog.logger.Warn(args...)
}

func Warning(args ...interface{}) {
	wLog.logger.Warning(args...)
}

func Error(args ...interface{}) {
	wLog.logger.Error(args...)
}

func Fatal(args ...interface{}) {
	wLog.logger.Fatal(args...)
}

func Panic(args ...interface{}) {
	wLog.logger.Panic(args...)
}
func WithField(key string, value interface{}) *logrus.Entry {
	return wLog.logger.WithField(key, value)
}
func WithFields(fields logrus.Fields) *logrus.Entry {
	return wLog.logger.WithFields(fields)
}
func WithError(err error) *logrus.Entry {
	return wLog.logger.WithError(err)
}

func Debugf(format string, args ...interface{}) {
	wLog.logger.Debugf(format, args...)

}
func Infof(format string, args ...interface{}) {
	wLog.logger.Infof(format, args...)
}
func Printf(format string, args ...interface{}) {
	wLog.logger.Printf(format, args...)

}
func Warnf(format string, args ...interface{}) {
	wLog.logger.Warnf(format, args...)

}
func Warningf(format string, args ...interface{}) {
	wLog.logger.Warningf(format, args...)

}
func Errorf(format string, args ...interface{}) {
	wLog.logger.Errorf(format, args...)

}
func Fatalf(format string, args ...interface{}) {
	wLog.logger.Fatalf(format, args...)

}
func Panicf(format string, args ...interface{}) {
	wLog.logger.Panicf(format, args...)

}

func Debugln(args ...interface{}) {
	wLog.logger.Debugln(args...)

}
func Infoln(args ...interface{}) {
	wLog.logger.Infoln(args...)

}
func Println(args ...interface{}) {
	wLog.logger.Println(args...)

}
func Warnln(args ...interface{}) {
	wLog.logger.Warnln(args...)

}
func Warningln(args ...interface{}) {
	wLog.logger.Warningln(args...)

}
func Errorln(args ...interface{}) {
	wLog.logger.Errorln(args...)

}
func Fatalln(args ...interface{}) {
	wLog.logger.Fatalln(args...)

}
func Panicln(args ...interface{}) {
	wLog.logger.Panicln(args...)
}
