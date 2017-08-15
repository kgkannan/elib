// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=socket -id event -d Type=event github.com/platinasystems/go/elib/elog/event.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package socket

import (
	"github.com/platinasystems/go/elib/elog"
)

var eventType = &elog.EventType{
	Name: "socket.event",
}

func init() {
	t := eventType
	t.Strings = stringer_event
	t.Encode = encode_event
	t.Decode = decode_event
	elog.RegisterType(eventType)
}

func stringer_event(t *elog.EventType, e *elog.Event) []string {
	var x event
	x.Decode(e.Data[:])
	return x.Strings()
}

func encode_event(b []byte, e *elog.Event) int {
	var x event
	x.Decode(e.Data[:])
	return x.Encode(b)
}

func decode_event(b []byte, e *elog.Event) int {
	var x event
	x.Decode(b)
	return x.Encode(e.Data[:])
}

func (x event) Log() { x.Logb(elog.DefaultBuffer) }

func (x event) Logb(b *elog.Buffer) {
	e := b.Add(eventType)
	x.Encode(e.Data[:])
}
