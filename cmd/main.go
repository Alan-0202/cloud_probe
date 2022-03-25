package main

import (
	"cloudprobe/internal/g"
	"cloudprobe/internal/readfile"
	"cloudprobe/internal/tasks"
	"cloudprobe/web"
)

func main() {

	//init config
	g.ParseConfig()

	// read file context
	readfile.NewServer().Do()

	// work
	tasks.NewServer(g.IcmpScripts,g.IcmpArgs,g.NcScripts,g.NcArgs,g.PosScripts).Handler()

	//web
	web.Start()

}
