/*
 * MIT License
 *
 * Copyright (c) 2021 Anthony Quéré
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package plugins

import (
	"log"
	"plugin"
)

type Plugin plugin.Plugin

type Event struct {
	Name string
}

type TaskContext struct {
	TaskName string
	Args     []string
}

func B(v string) {
	log.Printf("Here ! %s\n", v)
}

func A(v string) interface{} {
	B(v)
	return nil

}

func RunPlugin() string {
	p, err := plugin.Open("plugins/TestPlugin.so")
	if err != nil {
		return "Fail to launch plugin plugins/TaskRunner_TestPlugin.so : [" + err.Error() + "]"
	}
	log.Print("plugin launched\n")

	onEvent, err := p.Lookup("TestCallback")

	onEvent.(func(callback func(v string) interface{}))(A)

	return "porte"
	//return onEvent.(func(e string) interface{})("hey").(string)

}