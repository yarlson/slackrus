// Package slackrus provides a Slack hook for the logrus loggin package.
package slackrus

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

// Project version
const (
	Version = "1.0.0"
)

type SlackrusHookConfig struct {
	// Messages with a log level not contained in this array
	// will not be dispatched. If nil, all messages will be dispatched.
	AcceptedLevels []logrus.Level
	Token          string
	IconURL        string
	Channel        string
	IconEmoji      string
	Username       string
	Extra          map[string]interface{}
}

// SlackrusHook is a logrus Hook for dispatching messages to the specified
// channel on Slack.
type SlackrusHook struct {
	config      *SlackrusHookConfig
	slackClient *slack.Client
}

func NewSlackrusHook(config SlackrusHookConfig) *SlackrusHook {
	return &SlackrusHook{
		config:      &config,
		slackClient: slack.New(config.Token),
	}
}

// Levels sets which levels to sent to slack
func (sh *SlackrusHook) Levels() []logrus.Level {
	if sh.config.AcceptedLevels == nil {
		return AllLevels
	}

	return sh.config.AcceptedLevels
}

// Fire -  Sent event to slack
func (sh *SlackrusHook) Fire(e *logrus.Entry) error {
	color := ""
	switch e.Level {
	case logrus.DebugLevel:
		color = "#9B30FF"
	case logrus.InfoLevel:
		color = "good"
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		color = "danger"
	default:
		color = "warning"
	}

	newEntry := sh.newEntry(e)

	attach := slack.Attachment{Text: newEntry.Message}

	// If there are fields we need to render them at attachments
	if len(newEntry.Data) > 0 {

		// Add a header above field data
		attach.Text = "Message fields"

		for k, v := range newEntry.Data {
			slackField := slack.AttachmentField{}

			slackField.Title = k
			slackField.Value = fmt.Sprint(v)
			// If the field is <= 20 then we'll set it to short
			if len(slackField.Value) <= 20 {
				slackField.Short = true
			}

			attach.Fields = append(attach.Fields, slackField)
		}

		attach.Pretext = newEntry.Message
	}

	attach.Fallback = newEntry.Message
	attach.Color = color

	go func() {
		_, _, err := sh.slackClient.PostMessage(
			sh.config.Channel,
			slack.MsgOptionUsername(sh.config.Username),
			slack.MsgOptionIconEmoji(sh.config.IconEmoji),
			slack.MsgOptionAttachments(attach),
		)
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	return nil
}

func (sh *SlackrusHook) newEntry(entry *logrus.Entry) *logrus.Entry {
	data := map[string]interface{}{}

	for k, v := range sh.config.Extra {
		data[k] = v
	}
	for k, v := range entry.Data {
		data[k] = v
	}

	newEntry := &logrus.Entry{
		Logger:  entry.Logger,
		Data:    data,
		Time:    entry.Time,
		Level:   entry.Level,
		Message: entry.Message,
	}

	return newEntry
}
