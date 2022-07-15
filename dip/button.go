// Package dip show my understanding of dependency inversion principle.
package dip

type Button interface {
	TurnOn()
	TurnOff()
}

type UI struct {
	button Button
}

func NewUI(b Button) *UI {
	return &UI{button: b}
}

func (u *UI) Poll() {
	u.button.TurnOn()
	u.button.TurnOff()
}
