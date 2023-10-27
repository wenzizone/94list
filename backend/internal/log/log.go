package log

import (
	"fmt"
	"log/slog"
	"os"
)

var logger *slog.Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func Infof(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	fmt.Println(args...)
	logger.Info(msg)
}

func Errorf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	logger.Error(msg)
}
