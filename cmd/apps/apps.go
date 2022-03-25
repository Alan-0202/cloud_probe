package apps

import (
	"cloudprobe/internal/g"
	"cloudprobe/internal/tasks"
)

func Run() {
	jobs := tasks.NewServer(g.IcmpScripts,g.IcmpArgs,g.NcScripts,g.NcArgs,g.PosScripts)
	ch := make(chan bool)

	go oneJob(ch, jobs.IcmpHandler)
	go oneJob(ch, jobs.NcHandler)
	go oneJob(ch, jobs.PosHandler)

	<- ch
	<- ch
	<- ch
}

func oneJob(ch chan bool, f func()) {

	f()
	ch <- true
}