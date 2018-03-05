package log

import (
    "github.com/op/go-logging"
    "os"
)

var Log = logging.MustGetLogger("example")

var format = logging.MustStringFormatter(
    "%{time:15:04:05} %{longfunc} ▶ [%{level}] %{message}",
)

func init() {
    backend1 := logging.NewLogBackend(os.Stderr, "", 0)
    
    backend1Leveled := logging.AddModuleLevel(backend1)
    backend1Leveled.SetLevel(logging.DEBUG, "")
    
    logging.SetFormatter(format)

    logging.SetBackend(backend1Leveled)
}

