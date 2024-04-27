package log

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

type Log struct {
	infoLogger  *zerolog.Logger
	errorLogger *zerolog.Logger
}

func (log *Log) Info(str string) {
	log.infoLogger.Info().Msg(str)
}

func (log *Log) Error(str string) {
	log.infoLogger.Error().Msg(str)
}

func UnitFormatter() {
	zerolog.TimestampFunc = func() time.Time {
		format := "2006-01-02 15:04:05"
		timeString := time.Now().Format(format)
		timeFormatted, _ := time.Parse(format, timeString)
		return timeFormatted
	}

}

func InitLogger() (*Log, *os.File, *os.File) {
	UnitFormatter()
	loggerInfoFile, err := os.OpenFile("log/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		panic("Error opening info log file")
	}

	loggerErrorFile, err := os.OpenFile("log/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		panic("Error opening error log file")
	}

	infoLogger := zerolog.New(loggerInfoFile).With().Timestamp().Caller().Logger()
	errorLogger := zerolog.New(loggerErrorFile).With().Timestamp().Caller().Logger()

	log := &Log{
		infoLogger:  &infoLogger,
		errorLogger: &errorLogger,
	}

	return log, loggerInfoFile, loggerErrorFile
}
