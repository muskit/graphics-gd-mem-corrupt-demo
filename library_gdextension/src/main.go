// This file is all you need to start a project.
// Save it somewhere, install the `gd` command and use `gd run` to get started.
package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/Node"
	"graphics.gd/startup"
)

type TestClass struct {
  classdb.Extension[TestClass, Node.Instance] `gd:"TestClass"`
}

func (t *TestClass) Process(delta float32){
  Engine.Println("hello from testclass!!")
}

// Say hello!
func (t *TestClass) SayHello() {
  Engine.Println("Hello!")
}

func main() {
  classdb.Register[TestClass]()
  classdb.Register[Wordbag](RandomSeed)
  startup.Scene()
}
