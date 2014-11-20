package libkb

import (
	"fmt"
)

type Warning interface {
	Warning() string
}

type StringWarning string

func (s StringWarning) Warning() string {
	return string(s)
}

func Warningf(format string, a ...interface{}) Warning {
	return StringWarning(fmt.Sprintf(format, a...))
}

func ErrorToWarning(e error) Warning {
	if e == nil {
		return nil
	} else {
		return StringWarning(e.Error())
	}
}

type Warnings struct {
	w []Warning
}

func (w Warnings) Warnings() []Warning {
	return w.w
}

func (w Warnings) IsEmpty() bool {
	return w.w == nil || len(w.w) == 0
}

func (w *Warnings) Push(e Warning) {
	w.w = append(w.w, e)
}

func (w Warnings) Warn() {
	for _, e := range w.w {
		G.Log.Warning(e.Warning())
	}
}
