package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/yarlson/slackrus/v1"
)

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.SetOutput(os.Stderr)

	logrus.SetLevel(logrus.DebugLevel)

	logrus.AddHook(slackrus.NewSlackrusHook(slackrus.SlackrusHookConfig{
		Token:          "12345",
		AcceptedLevels: slackrus.LevelThreshold(logrus.DebugLevel),
		Channel:        "#slack-testing",
		IconEmoji:      ":ghost:",
		Username:       "foobot",
	}))

	logrus.WithFields(logrus.Fields{"foo": "bar", "foo2": 42}).Warn("this is a warn level message")
	logrus.Info("this is an info level message")
	logrus.Debug("this is a debug level message")
}
