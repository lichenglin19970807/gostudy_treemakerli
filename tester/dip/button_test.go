package tester

import (
	"gostudy_treemakerli/concurrency"
	"gostudy_treemakerli/dip"
	"gostudy_treemakerli/language"
	"testing"
)

func TestButton(t *testing.T) {
	runPoll := func(b dip.Button) {
		u := dip.NewUI(b)
		u.Poll()
	}

	runPoll(language.NewLanguageLamp())
	runPoll(concurrency.NewConcurrencyLamp())
}
