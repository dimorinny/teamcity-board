package screen

import (
	"github.com/dimorinny/teamcity-board/data"
	ui "github.com/gizak/termui"
)

const (
	oneSecondTimer = "/timer/1s"
)

type Context struct {
	updateInterval int
	client         *data.Teamcity
}

func NewContext(updateInterval int, client *data.Teamcity) *Context {
	return &Context{
		updateInterval: updateInterval,
		client:         client,
	}
}

func (c *Context) StartScreen(screen Screen) {
	c.clearHandlers()
	c.renderContent(screen)

	counter := 0
	// Only 1 second timer works properly
	c.AddHandler(oneSecondTimer, func(e ui.Event) {
		if counter%c.updateInterval == 0 {
			c.renderContent(screen)
		}
		counter += 1
	})
	screen.StartHandlers()
}

func (c *Context) renderContent(screen Screen) {
	ui.Body.Rows = screen.Content()
	ui.Body.Align()
	ui.Render(ui.Body)
}

func (c *Context) Init() {
	if err := ui.Init(); err != nil {
		panic(err)
	}
}

func (c *Context) StartLoop() {
	ui.Loop()
}

func (c *Context) Exit() {
	ui.StopLoop()
}

func (c *Context) Close() {
	ui.Close()
}

func (c *Context) AddHandler(path string, event func(ui.Event)) {
	ui.Handle(path, event)
}

func (c *Context) clearHandlers() {
	for key := range ui.DefaultEvtStream.Handlers {
		delete(ui.DefaultEvtStream.Handlers, key)
	}
}
