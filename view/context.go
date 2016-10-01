package view

import (
	"github.com/dimorinny/teamcity-board/data"
	ui "github.com/gizak/termui"
	"strconv"
)

const (
	oneSecondTimer = "/timer/1s"
	refreshKey     = "r"
	backKey        = "b"
)

type Context struct {
	updateInterval int
	backStack      []Screen
	Client         *data.Teamcity
}

func NewContext(updateInterval int, client *data.Teamcity) *Context {
	return &Context{
		updateInterval: updateInterval,
		Client:         client,
	}
}

func (c *Context) StartScreen(screen Screen, addToBackStack bool) {
	c.clearHandlers()

	c.AddTimerHandler(c.updateInterval, func(e ui.Event) {
		c.renderContent(screen)
	})
	c.AddKeyHandler(refreshKey, func(e ui.Event) {
		c.renderContent(screen)
	})
	c.AddKeyHandler(backKey, func(e ui.Event) {
		c.Back()
	})
	screen.StartHandlers()

	c.renderContent(screen)

	if addToBackStack {
		c.backStack = append(c.backStack, screen)
	}
}

func (c *Context) Back() {
	length := len(c.backStack)
	newScreenLength := length - 2

	if length <= 1 {
		c.Exit()
	} else {
		c.backStack = c.backStack[:length-1]
		c.StartScreen(c.backStack[newScreenLength], false)
	}
}

func (c *Context) renderContent(screen Screen) {
	ui.Clear()
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

func (c *Context) AddNumberHandler(event func(int)) {
	ui.Handle("/sys/kbd/", func(e ui.Event) {
		if e, ok := e.Data.(ui.EvtKbd); ok {
			if intKey, err := strconv.Atoi(e.KeyStr); err == nil && intKey > -1 && intKey < 10 {
				event(intKey)
			}
		}
	})
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
