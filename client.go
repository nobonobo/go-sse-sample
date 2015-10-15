// +build js

package main

import (
	"fmt"

	"github.com/gopherjs/eventsource"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	es := eventsource.New("http://localhost:3000/sse")
	es.AddEventListener("open", false, func(ev *js.Object) {
		fmt.Println("open", ev)
	})
	es.AddEventListener("message", false, func(ev *js.Object) {
		js.Global.Get("document").Call("writeln", ev.Get("data").String()+"<br/>")
	})
	es.AddEventListener("error", false, func(ev *js.Object) {
		fmt.Println("error", ev)
	})
}
