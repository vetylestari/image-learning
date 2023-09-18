package infrastructure

import (
	"fmt"
	"io"
	"os"

	"github.com/multiplay/go-slack/chat"
	"github.com/multiplay/go-slack/lrhook"
	"github.com/sirupsen/logrus"
)

func InitLog() *logrus.Logger {
	cfg := lrhook.Config{
		MinLevel: logrus.ErrorLevel,
		Message: chat.Message{
			Channel:   fmt.Sprintf("#%s", os.Getenv("SLACK_CHANNEL_NAME")),
			IconEmoji: ":ghost:",
		},
	}
	h := lrhook.New(cfg, os.Getenv("SLACK_WEBHOOK_URL"))
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: false,
	})
	log.SetLevel(logrus.DebugLevel)
	log.SetOutput(io.MultiWriter(os.Stderr))
	log.AddHook(h)

	return log
}
