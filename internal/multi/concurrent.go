package multi

import "cloudprobe/internal/g"

type ConCurTasks struct {
	num int
	c chan struct{}
}

func NewConCurTasks() *ConCurTasks {
	return &ConCurTasks{
		num: *g.ConcurrTasks,
		c: make(chan struct{}, *g.ConcurrTasks),
	}
}

func (g *ConCurTasks) Run(f func()) {
	g.c <- struct{}{}
	go func() {
		f()
		<-g.c
	}()
}
