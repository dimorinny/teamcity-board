package widget

import (
	"fmt"
	"github.com/dimorinny/teamcity-board/view"
	ui "github.com/gizak/termui"
	"time"
)

const (
	logTitle   = "System Log"
	dateFormat = time.RFC822
)

type LogView struct {
	limit    int
	messages []string
}

func NewLogView(limit int) LogView {
	return LogView{
		limit:    limit,
		messages: []string{},
	}
}

func (logView *LogView) AddMessage(message string) {
	newMessages := append(logView.messages, message)
	messagesLength := len(newMessages)
	// Get last messages
	if messagesLength > logView.limit {
		newMessages = newMessages[messagesLength-logView.limit:]
	}

	logView.messages = newMessages
}

func (logView *LogView) GenerateView() *ui.List {
	ls := ui.NewList()
	ls.Border = true
	ls.BorderLabel = logTitle

	for _, message := range logView.messages {
		logMessage := fmt.Sprintf(
			"[%s] %s",
			time.Now().Format(dateFormat),
			message,
		)
		ls.Items = append(ls.Items, logMessage)
	}

	ls.Height = len(logView.messages) + view.BoardHeight
	return ls
}
