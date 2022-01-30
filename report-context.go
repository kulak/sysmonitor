package main

import "time"

type Level string
type Kind string

const errorLvl Level = "error"
const infoLvl Level = "info"

const code Kind = "code"
const p Kind = "p"

type Group struct {
	Title       string
	Description string
	Msgs        []Message
}
type Message struct {
	Level   Level
	Kind    Kind
	Message string
}

type ReportContext struct {
	Now    time.Time
	Groups []Group
}

func Msg(m string, l Level, k Kind) Message {
	return Message{
		Message: m,
		Level:   l,
		Kind:    k,
	}
}
