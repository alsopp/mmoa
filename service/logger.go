package service

// Monolithic tools.Message-Oriented Application (MMOA)
// Logger
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"log"
)

// NewLogger - create a new Logger
func NewLogger() *Logger {
	l := &Logger{}
	return l
}

// Logger structure
type Logger struct {
	codError map[int]string
}

func (l *Logger) Message() *Msg {
	em := &Msg{}
	return em
}

// Msg - message struct
type Msg struct {
	buf []interface{}
}

func (e *Msg) Field(key string, value interface{}) *Msg {
	e.buf = append(e.buf, key, `="`, value, `" `)
	return e
}

func (e *Msg) Send() {
	log.Print(e.buf...)
}
