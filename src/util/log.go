package util

import (
    "context"
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "sync"
    "path/filepath"
    "strings"
    "time"

    logrus "github.com/sirupsen/logrus"
)

var (
    maxSizeMB        int64  = 10
    maxFiles         int    = 10
    defaultLogLevel  string = "INFO"
)


type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
    timestamp := entry.Time.Format("2006-01-02 15:04:05.000000000")
    logLevel := strings.ToUpper(entry.Level.String())
    var file string
    var line int
    if entry.HasCaller() {
        file = filepath.Base(entry.Caller.File)
        line = entry.Caller.Line
    }
    msg := fmt.Sprintf("%s | %s | [%s:%d] %s\n", timestamp, logLevel, file, line, entry.Message)
    return []byte(msg), nil
}


func NewWriterHook(writer io.Writer, logLevels []logrus.Level, file *os.File, level logrus.Level) *WriterHook {
    return &WriterHook{
        Writer:    writer,
        LogLevels: logLevels,
        File:      file,
        Level:     level,
        mu:        &sync.Mutex{},
    }
}

type WriterHook struct {
    Writer    io.Writer
    LogLevels []logrus.Level
    File      *os.File
    Level    logrus.Level
    mu        *sync.Mutex 
}

func (hook *WriterHook) Fire(entry *logrus.Entry) error {
    hook.mu.Lock()
    defer hook.mu.Unlock()

    for _, level := range hook.LogLevels {
        if entry.Level == level {
            line, err := entry.String()
            if err != nil {
                return err
            }
            _, err = hook.Writer.Write([]byte(line))
            return err
        }
    }
    return nil
}

func (hook *WriterHook) Levels() []logrus.Level {
    return hook.LogLevels
}

func (hook *WriterHook) UpdateWriter(newWriter io.Writer, newFile *os.File) {
    hook.mu.Lock()
    defer hook.mu.Unlock()

    if hook.File != nil {
        hook.File.Close()
    }
    hook.Writer = newWriter
    hook.File = newFile
}



func SetupLogging(ctx context.Context, logPath, model, name, logLevel string) *logrus.Logger {
    logLevel = strings.ToUpper(logLevel)

    if logPath == "" {
        logPath = "/tmp/log"
    }
    logDir := filepath.Join(logPath, name)

    if err := os.MkdirAll(logDir, 0750); err != nil {
        logrus.Fatalf("Failed to create log directory: %v", err)
    }

    logger := logrus.New()
    logger.SetFormatter(new(CustomFormatter))
    logger.SetReportCaller(true)
    logger.Out = ioutil.Discard
    
    setlogLevel := setLogLevel(logLevel)
    defaultlogLevel := setLogLevel(defaultLogLevel)
    logger.SetLevel(setlogLevel)

    initializeHookForLevel(logger, logDir, name, setlogLevel)
    initializeHookForLevel(logger, logDir, name, defaultlogLevel)

    go rotateLogFile(ctx, logger, logDir, name, maxSizeMB, maxFiles,setlogLevel)
    go rotateLogFile(ctx, logger, logDir, name, maxSizeMB, maxFiles,defaultlogLevel)

    return logger
}

func setLogLevel(logLevel string) logrus.Level {
    switch strings.ToUpper(logLevel) {
    case "DEBUG":
        return logrus.DebugLevel
    case "INFO":
        return logrus.InfoLevel
    case "WARNING":
        return logrus.WarnLevel
    case "ERROR":
        return logrus.ErrorLevel
    case "CRITICAL":
        return logrus.FatalLevel
    default:
        return logrus.InfoLevel
    }
}


func rotateLogFile(ctx context.Context, logger *logrus.Logger, logDir, name string, maxSizeMB int64, maxFiles int, level logrus.Level) {
    checkLogDuration := 1 * time.Second
    checkSizeTicker := time.NewTicker(checkLogDuration)
    defer checkSizeTicker.Stop()

    for {
        select {
        case <-ctx.Done():
            logger.Info("Log rotation stopped due to context cancellation")
            return
        case <-checkSizeTicker.C:
            logFilePath := filepath.Join(logDir, fmt.Sprintf("%s_%s_go.log", name, strings.ToUpper(level.String())))
            fileInfo, err := os.Stat(logFilePath)
            if err != nil && !os.IsNotExist(err) {
                logger.Errorf("Failed to get log file info: %v", err)
                continue
            }
    
            if fileInfo != nil && fileInfo.Size() >= maxSizeMB*1024*1024 {
                for i := maxFiles - 1; i >= 0; i-- {
                    oldPath := fmt.Sprintf("%s.%d", logFilePath, i)
                    newPath := fmt.Sprintf("%s.%d", logFilePath, i+1)
                    if i == 0 {
                        oldPath = logFilePath
                    }
                    if _, err := os.Stat(oldPath); err == nil {
                        os.Rename(oldPath, newPath)
                    }
                }
    
                newFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
                if err != nil {
                    logger.Errorf("Failed to open new log file: %v", err)
                    continue
                }
    
                updateLevelHook(logger, level, newFile)
            }

        }
    }
}

func initializeHookForLevel(logger *logrus.Logger, logDir, name string, level logrus.Level) {
    levelStr := strings.ToUpper(level.String())
    filePath := filepath.Join(logDir, fmt.Sprintf("%s_%s_go.log", name, levelStr))
    file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        logrus.Fatalf("Failed to open log file: %v", err)
    }

    var levels []logrus.Level
    for _, l := range logrus.AllLevels {
        if l <= level {
            levels = append(levels, l)
        }
    }

    hook := NewWriterHook(file, levels, file, level)
    logger.AddHook(hook)
}


func updateLevelHook(logger *logrus.Logger, level logrus.Level, newFile *os.File) {
    for _, hook := range logger.Hooks[level] {
        if writerHook, ok := hook.(*WriterHook); ok {
            if writerHook.Level == level { 
                writerHook.UpdateWriter(newFile, newFile)
            }
        }
    }
}
