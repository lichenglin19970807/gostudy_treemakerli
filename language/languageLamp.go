package language

import "fmt"

type LanguageLamp struct {
}

func NewLanguageLamp() *LanguageLamp {
	return &LanguageLamp{}
}

func (l *LanguageLamp) TurnOn() {
	fmt.Println("LanguageLamp turn on.")
}

func (l *LanguageLamp) TurnOff() {
	fmt.Println("LanguageLamp turn off.")
}
