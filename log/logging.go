package log

import (
	"strings"
	// "fmt"
	log "github.com/sirupsen/logrus"
	"github.com/t-tomalak/logrus-easy-formatter"
	"io"
	"os"
)

/*
PRINT :
Input:  content -> string
		level -> int 	+--> Debug
						+--> Info
						+--> Warn
						+--> Error
						+--> Fatal
						+--> Panic
*/
func PRINT(content string, level int) {
	typeName := "TEST LOG"
	targetFileName := "PLS_GO.log"
	file, err := os.OpenFile(targetFileName,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
	// log.SetFormatter(&log.JSONFormatter{})
	// log.SetLevel(log.InfoLevel)
	logger := &log.Logger{
		Out:   mw,
		Level: log.InfoLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "%time% | %lvl% | %type% | %msg%",
		},
	}
	logger2 := &log.Logger{
		Out:   os.Stdout,
		Level: log.DebugLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "%time% | %lvl% | %type% | %msg%",
		},
	}
	switch level {
	case 1:
		logger2.WithFields(log.Fields{
			"type": typeName,
		}).Debugln(content)
	case 2:
		logger.WithFields(log.Fields{
			"type": typeName,
		}).Infoln(content)
	case 3:
		logger.WithFields(log.Fields{
			"type": typeName,
		}).Warnln(content)
	case 4:
		logger.WithFields(log.Fields{
			"type": typeName,
		}).Errorln(content)
	case 5:
		logger.WithFields(log.Fields{
			"type": typeName,
		}).Fatalln(content)
	case 6:
		logger.WithFields(log.Fields{
			"type": typeName,
		}).Panicln(content)
	}

}

/*
PRINTF :
Input:  content -> string
		level -> int 	+--> Debug
						+--> Info
						+--> Warn
						+--> Error
						+--> Fatal
						+--> Panic
		targetFileName -> string
*/
func PRINTF(content string, level int, targetFileName string) {
	var typeName string
	if strings.HasSuffix(targetFileName, ".log") {
		temp := strings.Split(targetFileName, ".")
		typeName = temp[0]
	} else {
		log.Fatal("Target File name must end with .log\n")
	}
	// targetFileName := "PLS_GO.log"
	file, err := os.OpenFile(targetFileName,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
	// log.SetFormatter(&log.JSONFormatter{})
	// log.SetLevel(log.InfoLevel)
	logger := &log.Logger{
		Out:   mw,
		Level: log.InfoLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "%time% | %lvl% | %type% | %msg%",
		},
	}
	logger2 := &log.Logger{
		Out:   os.Stdout,
		Level: log.DebugLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "%time% | %lvl% | %type% | %msg%",
		},
	}
	switch level {
	case 1:
		logger2.WithFields(log.Fields{
			"type": typeName,
		}).Debugln(content)
	case 2:
		logger.WithFields(log.Fields{
			"type": typeName,
		}).Infoln(content)
	case 3:
		logger.WithFields(log.Fields{
			"type": typeName,
		}).Warnln(content)
	case 4:
		logger.WithFields(log.Fields{
			"type": typeName,
		}).Errorln(content)
	case 5:
		logger.WithFields(log.Fields{
			"type": typeName,
		}).Fatalln(content)
	case 6:
		logger.WithFields(log.Fields{
			"type": typeName,
		}).Panicln(content)
	}

}

// Debug :
func Debug(content string) {
	PRINT(content, 1)
}

// Info :
func Info(content string) {
	PRINT(content, 2)
}

// Warn :
func Warn(content string) {
	PRINT(content, 3)
}

// Error :
func Error(content string) {
	PRINT(content, 4)
}

// Fatal :
func Fatal(content string) {
	PRINT(content, 5)
}

// Panic :
func Panic(content string) {
	PRINT(content, 6)
}
