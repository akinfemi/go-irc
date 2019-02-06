// package main

// type Input struct {
// 	name      string
// 	x, y      int
// 	w         int
// 	maxLength int
// }

// func NewInput(name string, x, y, w, maxLength int) *Input {
// 	return &Input{name: name, x: x, y: y, w: w, maxLength: maxLength}
// }

// func (i *Input) Layout(g *gocui.Gui) error {
// 	v, err := g.SetView(i.name, i.x, i.y, i.x+i.w, i.y+2)
// 	if err != nil {
// 		if err != gocui.ErrUnknownView {
// 			return err
// 		}
// 		v.Editor = i
// 		v.Editable = true
// 	}
// 	return nil
// }

// func (i *Input) Edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
// 	switch {
// 	case ch != 0 && mod == 0:
// 		v.EditWrite(ch)
// 	case key == gocui.KeySpace:
// 		v.EditWrite(' ')
// 	case key == gocui.KeyBackspace || key == gocui.KeyBackspace2:
// 		v.EditDelete(true)
// 	}
// }

// func SetFocus(name string) func(g *gocui.Gui) error {
// 	return func(g *gocui.Gui) error {
// 		_, err := g.SetCurrentView(name)
// 		return err
// 	}
// }

// func main() {
// 	g, err := gocui.NewGui(gocui.OutputNormal)
// 	if err != nil {
// 		log.Panicln(err)
// 	}
// 	defer g.Close()

// 	g.Cursor = true
// 	maxX, maxY := g.Size()
// 	input := NewInput("input", maxX/4+2, maxY-4, (maxX*3/4)-5, maxY-2)
// 	focus := gocui.ManagerFunc(SetFocus("input"))
// 	g.SetManager(input, focus)

// 	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
// 		log.Panicln(err)
// 	}

// 	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
// 		log.Panicln(err)
// 	}
// }

// func quit(g *gocui.Gui, v *gocui.View) error {
// 	return gocui.ErrQuit
// }
