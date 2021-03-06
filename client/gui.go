package main

import (
	"log"
	"net"
	"sync"

	"github.com/akinfemi/go-irc/server"
	"github.com/jroimartin/gocui"
)

//Client : Mother Struct
type Client struct {
	connected bool
	conn      net.Conn
	sv        *server.Server
	mutex     sync.Mutex
}

var client = &Client{connected: false, conn: nil, sv: nil}

type Reply struct {
	command string
	state   string
	body    []string
}

//Input Text box
type Input struct {
	name      string
	x, y      int
	w         int
	maxLength int
}

//BoxView each view
type BoxView struct {
	name      string
	x, y      int
	w, h      int
	maxLength int
}

//NewInput new textbox
func NewInput(name string, x, y, w, maxLength int) *Input {
	return &Input{name: name, x: x, y: y, w: w, maxLength: maxLength}
}

//NewBoxView : new views
func NewBoxView(name string, x, y, w, h int) *BoxView {
	return &BoxView{name: name, x: x, y: y, w: w, h: h}
}

//Layout : Each view must implement a Layout function
func (b *BoxView) Layout(g *gocui.Gui) error {
	v, err := g.SetView(b.name, b.x, b.y, b.w, b.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = b.name
	}
	return nil
}

//Layout : Layout function for textbox
func (i *Input) Layout(g *gocui.Gui) error {
	v, err := g.SetView(i.name, i.x, i.y, i.w, i.y+2)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editor = i
		v.Editable = true
	}
	return nil
}

//Edit : Editable views must implement Edit
func (i *Input) Edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch {
	case ch != 0 && mod == 0:
		v.EditWrite(ch)
	case key == gocui.KeySpace:
		v.EditWrite(' ')
	case key == gocui.KeyBackspace || key == gocui.KeyBackspace2:
		v.EditDelete(true)
	}
}

//SetFocus : Focus on view
func SetFocus(name string) func(g *gocui.Gui) error {
	return func(g *gocui.Gui) error {
		_, err := g.SetCurrentView(name)
		return err
	}
}

func updateViewList(v *gocui.View, users []string) error {
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func must(err error) {
	if err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	must(err)
	defer g.Close()

	g.Cursor = true
	maxX, maxY := g.Size() //get screen size
	chat := NewBoxView("chat", maxX/4+1, 0, maxX-1, maxY-5)
	channels := NewBoxView("channels", 0, 0, maxX/4, maxY/2-1)
	users := NewBoxView("users", 0, maxY/2+1, maxX/4, maxY-2)
	input := NewInput("input", maxX/4+1, maxY-4, maxX-1, maxY-2)
	focus := gocui.ManagerFunc(SetFocus("input"))

	g.SetManager(chat, channels, users, input, focus)

	must(g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit))
	must(g.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, readAndSend))
	must(g.MainLoop())
}
