slackrus
========

Slack hook for [Logrus](https://github.com/sirupsen/logrus). 

## Use

```go
package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/yarlson/slackrus"
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

```

### Extra fields
You can also add some extra fields to be sent with every slack message
```go
	extra := map[string]interface{}{
		"hostname": "nyc-server-1",
		"tag": "some-tag",
	}
```

## Parameters

#### Required
  * Token

#### Optional
  * IconEmoji
  * IconURL
  * Username
  * Channel
  * Extra
## Installation

    go get -u github.com/yarlson/slackrus/v1

## Credits

Based on Slackrus by [johntdyer](https://github.com/johntdyer/slackrus)
