// Copyright 2016 Richard Hawkins
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Package events manages events

package events

import (
	"runtime"

	"github.com/go-gl/glfw/v3.1/glfw"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

// Handler for events
type Handler interface {
	Handle(event Event)
}

type Event struct {
	Window   *glfw.Window
	KeyEvent bool
	Key      glfw.Key
	Scancode int
	Action   glfw.Action
	Mods     glfw.ModifierKey
	X        float32
	Y        float32
}

var events []Event

func Get() []Event {
	// TODO: This might cause a lot of garbage collection, which is prob bad.
	var elist []Event
	for i := range events {
		elist = append(elist, Event{
			Window:   events[i].Window,
			KeyEvent: events[i].KeyEvent,
			Key:      events[i].Key,
			Scancode: events[i].Scancode,
			Action:   events[i].Action,
			Mods:     events[i].Mods,
			X:        events[i].X,
			Y:        events[i].Y,
		})
	}
	events = nil
	return elist
}

// CursorPositionCallback TODO doc
func CursorPositionCallback(w *glfw.Window, x, y float64) {
	// TODO: these are from the top/left should be bottom/left to match sprite drawing
	events = append(events, Event{
		Window:   w,
		KeyEvent: false,
		X:        float32(x),
		Y:        float32(y),
	})
}

// KeyCallback TODO doc
func KeyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	events = append(events, Event{
		Window:   w,
		KeyEvent: true,
		Key:      key,
		Scancode: scancode,
		Action:   action,
		Mods:     mods,
	})
}
