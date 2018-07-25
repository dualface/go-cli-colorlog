// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cli

import (
    "fmt"
    "log"
    "strings"

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

func PrintLog(priority LogPriority, lines string) {
    if PrintPriority < priority {
        return
    }

    switch priority {
    case LOG_EMERG:
        PrintEmerg(lines)
    case LOG_ALERT:
        PrintAlert(lines)
    case LOG_CRIT:
        PrintCrit(lines)
    case LOG_ERR:
        PrintError(lines)
    case LOG_WARNING:
        PrintWarn(lines)
    case LOG_NOTICE:
        PrintNotice(lines)
    case LOG_INFO:
        PrintInfo(lines)
    default:
        PrintDebug(lines)
    }
}

func PrintEmerg(lines string) {
    if PrintPriority < LOG_EMERG {
        return
    }
    Println(Color.Emerg, Prefix.Emerg, lines, false)
}

func PrintAlert(lines string) {
    if PrintPriority < LOG_ALERT {
        return
    }
    Println(Color.Alert, Prefix.Alert, lines, false)
}

func PrintCrit(lines string) {
    if PrintPriority < LOG_CRIT {
        return
    }
    Println(Color.Crit, Prefix.Crit, lines, false)
}

func PrintError(lines string) {
    if PrintPriority < LOG_ERR {
        return
    }
    Println(Color.Error, Prefix.Error, lines, true)
}

func PrintWarn(lines string) {
    if PrintPriority < LOG_WARNING {
        return
    }
    Println(Color.Warn, Prefix.Warn, lines, true)
}

func PrintNotice(lines string) {
    if PrintPriority < LOG_NOTICE {
        return
    }
    Println(Color.Notice, Prefix.Notice, lines, true)
}

func PrintInfo(lines string) {
    if PrintPriority < LOG_INFO {
        return
    }
    Println(Color.Info, Prefix.Info, lines, true)
}

func PrintDebug(lines string) {
    if PrintPriority < LOG_DEBUG {
        return
    }
    Println(Color.Debug, Prefix.Debug, lines, true)
}

func PrintOutput(lines string) {
    Println(Color.Output, Prefix.Output, lines, true)
}

func Println(color ct.Color, prefix string, lines string, onlyColorPrefix bool) {
    if PrintToLogger {
        log.Printf("%s%s", prefix, lines)
    } else {
        for _, line := range strings.Split(lines, "\n") {
            ct.Foreground(color, false)
            fmt.Print(prefix)
            if onlyColorPrefix {
                ct.ResetColor()
            }
            fmt.Println(line)
            ct.ResetColor()
        }
    }
}
