package glog

import (
    "flag"
)


type loggerLevel int

const (
    LOGGER_LEVEL_DEBUG      = iota
    LOGGER_LEVEL_INFO
    LOGGER_LEVEL_WARN
    LOGGER_LEVEL_ERROR
    LOGGER_LEVEL_FATAL
    LOGGER_LEVEL_MAX        = LOGGER_LEVEL_FATAL
)

var loggerLevelNames = []string {
    LOGGER_LEVEL_DEBUG:     "DEBUG",
    LOGGER_LEVEL_INFO:      "INFO",
    LOGGER_LEVEL_WARN:      "WARN",
    LOGGER_LEVEL_ERROR:     "ERROR",
    LOGGER_LEVEL_FATAL:     "FATAL",
}

type loggerInterface interface {
    Init(loggerLevel int, logPath string, dailyRolling bool, fileMaxSize uint64)
    Release()

    SetLevel(level int)

    Fatal(args ...interface{})
    Fatalf(format string, args ...interface{})
    Fatalln(args ...interface{})

    Error(args ...interface{})
    Errorf(format string, args ...interface{})
    Errorln(args ...interface{})

    Warn(args ...interface{})
    Warnf(format string, args ...interface{})
    Warnln(args ...interface{})

    Info(args ...interface{})
    Infof(format string, args ...interface{})
    Infoln(args ...interface{})

    Debug(args ...interface{})
    Debugf(format string, args ...interface{})
    Debugln(args ...interface{})
}

type glogLogger struct {
}

var Logger loggerInterface = new(glogLogger)

func (logger *glogLogger) Init(loggerLevel int, logPath string, dailyRolling bool, fileMaxSize uint64) {
    flag.Parse()

    SetStackFramesDepth(4)

    if loggerLevel >= 0 && loggerLevel < len(loggerLevelNames) {
        SetLevelString(loggerLevelNames[loggerLevel])
    } else {
        SetLevelString(loggerLevelNames[len(loggerLevelNames) - 1])
    }

    path := logPath
    SetLogDir(&path)
    SetDailyRolling(dailyRolling)
    SetFileMaxSize(fileMaxSize)
}

func (logger *glogLogger) Release() {
    Flush()
}

func (logger *glogLogger) SetLevel(level int)  {
    if level >= 0 && level < len(loggerLevelNames) {
        SetLevelString(loggerLevelNames[level])
    } else {
        SetLevelString(loggerLevelNames[len(loggerLevelNames) - 1])
    }
}

func (logger *glogLogger) Fatal(args ...interface{}) {
    Fatal(args...)
}

func (logger *glogLogger) Fatalf(format string, args ...interface{}) {
    Fatalf(format, args...)
}

func (logger *glogLogger) Fatalln(args ...interface{}) {
    Fatalln(args...)
}

func (logger *glogLogger) Error(args ...interface{}) {
    Error(args...)
}

func (logger *glogLogger) Errorf(format string, args ...interface{}) {
    Errorf(format, args...)
}

func (logger *glogLogger) Errorln(args ...interface{}) {
    Errorln(args...)
}

func (logger *glogLogger) Warn(args ...interface{}) {
    Warning(args...)
}

func (logger *glogLogger) Warnf(format string, args ...interface{}) {
    Warningf(format, args...)
}

func (logger *glogLogger) Warnln(args ...interface{}) {
    Warningln(args...)
}

func (logger *glogLogger) Info(args ...interface{}) {
    Info(args...)
}
func (logger *glogLogger) Infof(format string, args ...interface{}) {
    Infof(format, args...)
}

func (logger *glogLogger) Infoln(args ...interface{}) {
    Infoln(args...)
}

func (logger *glogLogger) Debug(args ...interface{}) {
    Debug(args...)
}

func (logger *glogLogger) Debugf(format string, args ...interface{}) {
    Debugf(format, args...)
}

func (logger *glogLogger) Debugln(args ...interface{}) {
    Debugln(args...)
}
