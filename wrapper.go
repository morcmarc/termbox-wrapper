package tbw

import (
	termbox "github.com/nsf/termbox-go"
)

type TermboxApi interface {
	Init() error
	Interrupt()
	Close()
	Flush()
	SetCursor(x, y int)
	HideCursor()
	SetCell(x, y int, ch rune, fg, bg termbox.Attribute)
	CellBuffer() []termbox.Cell
	ParseEvent(data []byte) termbox.Event
	PollRawEvent(data []byte) termbox.Event
	PollEvent() termbox.Event
	Size() (int, int)
	Clear(fg, bg termbox.Attribute) error
	SetInputMode(mode termbox.InputMode) termbox.InputMode
	SetOutputMode(mode termbox.OutputMode) termbox.OutputMode
	Sync() error
}

type TermboxWrapper struct {
	TermboxApi
}

func NewTermboxWrapper() *TermboxWrapper {
	return &TermboxWrapper{}
}

func (t TermboxWrapper) Init() error {
	return termbox.Init()
}

func (t TermboxWrapper) Interrupt() {
	termbox.Interrupt()
}

func (t TermboxWrapper) Close() {
	termbox.Close()
}

func (t TermboxWrapper) Flush() {
	termbox.Flush()
}

func (t TermboxWrapper) SetCursor(x, y int) {
	termbox.SetCursor(x, y)
}

func (t TermboxWrapper) HideCursor() {
	termbox.HideCursor()
}

func (t TermboxWrapper) SetCell(x, y int, ch rune, fg, bg termbox.Attribute) {
	termbox.SetCell(x, y, ch, fg, bg)
}

func (t TermboxWrapper) CellBuffer() []termbox.Cell {
	return termbox.CellBuffer()
}

func (t TermboxWrapper) ParseEvent(data []byte) termbox.Event {
	return termbox.ParseEvent(data)
}

func (t TermboxWrapper) PollRawEvent(data []byte) termbox.Event {
	return termbox.PollRawEvent(data)
}

func (t TermboxWrapper) PollEvent() termbox.Event {
	return termbox.PollEvent()
}

func (t TermboxWrapper) Size() (int, int) {
	return termbox.Size()
}

func (t TermboxWrapper) Clear(fg, bg termbox.Attribute) error {
	return termbox.Clear(fg, bg)
}

func (t TermboxWrapper) SetInputMode(mode termbox.InputMode) termbox.InputMode {
	return termbox.SetInputMode(mode)
}

func (t TermboxWrapper) SetOutputMode(mode termbox.OutputMode) termbox.OutputMode {
	return termbox.SetOutputMode(mode)
}

func (t TermboxWrapper) Sync() error {
	return termbox.Sync()
}
