package clog

import (
    "fmt"
    "log"

    "github.com/daviddengcn/go-colortext"
)

type (
    LogPriority int
)

const (
    LOG_EMERG   LogPriority = iota
    LOG_ALERT
    LOG_CRIT
    LOG_ERR
    LOG_WARNING
    LOG_NOTICE
    LOG_INFO
    LOG_DEBUG
)

var Color = struct {
    Emerg  ct.Color
    Alert  ct.Color
    Crit   ct.Color
    Error  ct.Color
    Warn   ct.Color
    Notice ct.Color
    Info   ct.Color
    Debug  ct.Color
    Output ct.Color
}{
    ct.Red,
    ct.Red,
    ct.Red,
    ct.Red,
    ct.Yellow,
    ct.Magenta,
    ct.Green,
    ct.Cyan,
    ct.Cyan,
}

var Prefix = struct {
    Emerg  string
    Alert  string
    Crit   string
    Error  string
    Warn   string
    Notice string
    Info   string
    Debug  string
    Output string
}{
    "[EMERG] ",
    "[ALERT] ",
    "[CRIT] ",
    "[ERR] ",
    "[WARN] ",
    "[NOTICE] ",
    "[INFO] ",
    "[DEBUG] ",
    ">> ",
}

var PrintToLogger = false
var PrintPriority = LOG_DEBUG

func PrintLog(priority LogPriority, ft string, args ...interface{}) {
    if PrintPriority < priority {
        return
    }

    switch priority {
    case LOG_EMERG:
        PrintEmerg(ft, args...)
    case LOG_ALERT:
        PrintAlert(ft, args...)
    case LOG_CRIT:
        PrintCrit(ft, args...)
    case LOG_ERR:
        PrintError(ft, args...)
    case LOG_WARNING:
        PrintWarn(ft, args...)
    case LOG_NOTICE:
        PrintNotice(ft, args...)
    case LOG_INFO:
        PrintInfo(ft, args...)
    default:
        PrintDebug(ft, args...)
    }
}

func PrintEmerg(ft string, args ...interface{}) {
    if PrintPriority < LOG_EMERG {
        return
    }
    Println(Color.Emerg, Prefix.Emerg, ft, args...)
}

func PrintAlert(ft string, args ...interface{}) {
    if PrintPriority < LOG_ALERT {
        return
    }
    Println(Color.Alert, Prefix.Alert, ft, args...)
}

func PrintCrit(ft string, args ...interface{}) {
    if PrintPriority < LOG_CRIT {
        return
    }
    Println(Color.Crit, Prefix.Crit, ft, args...)
}

func PrintError(ft string, args ...interface{}) {
    if PrintPriority < LOG_ERR {
        return
    }
    Println(Color.Error, Prefix.Error, ft, args...)
}

func PrintWarn(ft string, args ...interface{}) {
    if PrintPriority < LOG_WARNING {
        return
    }
    Println(Color.Warn, Prefix.Warn, ft, args...)
}

func PrintNotice(ft string, args ...interface{}) {
    if PrintPriority < LOG_NOTICE {
        return
    }
    Println(Color.Notice, Prefix.Notice, ft, args...)
}

func PrintInfo(ft string, args ...interface{}) {
    if PrintPriority < LOG_INFO {
        return
    }
    Println(Color.Info, Prefix.Info, ft, args...)
}

func PrintDebug(ft string, args ...interface{}) {
    if PrintPriority < LOG_DEBUG {
        return
    }
    Println(Color.Debug, Prefix.Debug, ft, args...)
}

func PrintOutput(ft string, args ...interface{}) {
    Println(Color.Output, Prefix.Output, ft, args...)
}

func Println(color ct.Color, prefix string, ft string, args ...interface{}) {
    if PrintToLogger {
        log.Println(prefix + fmt.Sprintf(ft, args...))
    } else {
        ct.Foreground(color, false)
        fmt.Print(prefix)
        ct.ResetColor()
        fmt.Printf(ft, args...)
        fmt.Println()
    }
}
