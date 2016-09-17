package screen

import (
	"github.com/dimorinny/teamcity-board/data"
	ui "github.com/gizak/termui"
)

const (
	oneSecondTimer = "/timer/1s"
	refreshKey     = "r"
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

	c.AddTimerHandler(c.updateInterval, func(e ui.Event) {
		c.renderContent(screen)
	})
	c.AddKeyHandler(refreshKey, func(e ui.Event) {
		c.renderContent(screen)
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

func (c *Context) AddKeyHandler(key string, event func(ui.Event)) {
	ui.Handle("/sys/kbd/"+key, event)
}

func (c *Context) AddTimerHandler(seconds int, event func(ui.Event)) {
	counter := 0
	// Only 1 second timer works properly in termui
	ui.Handle(oneSecondTimer, func(e ui.Event) {
		if counter%seconds == 0 {
			event(e)
			counter = 0
		}
		counter += 1
	})
}

func (c *Context) AddHandler(path string, event func(ui.Event)) {
	ui.Handle(path, event)
}

func (c *Context) clearHandlers() {
	for key := range ui.DefaultEvtStream.Handlers {
		delete(ui.DefaultEvtStream.Handlers, key)
	}
}
