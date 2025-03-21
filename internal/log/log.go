package log

import (
	"io"
	"os"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func InitLogger() {
    Logger.SetFormatter(&logrus.TextFormatter{
        TimestampFormat: "15:04:05",
        FullTimestamp: true,
        ForceColors: true,
    })

    logFilePath := getLogFilePath()
    os.MkdirAll(filepath.Dir(logFilePath), os.ModePerm)

    logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err == nil {
        multiWriter := io.MultiWriter(os.Stdout, logFile)
        Logger.SetOutput(multiWriter)
    } else {
        Logger.SetOutput(os.Stdout)
    }
}

func getLogFilePath() string {
    switch runtime.GOOS {
    case "windows":
        return filepath.Join(os.Getenv("APPDATA"), "ripple", "logs", "ripple.log")
    case "darwin":
        return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "ripple", "logs", "ripple.log")
    default:
        return filepath.Join(os.Getenv("HOME"), ".local", "share", "ripple", "logs", "ripple.log")
    }
}
