package tasks

import (
	"cloudprobe/internal/icmp"
	"cloudprobe/internal/log"
	"cloudprobe/internal/multi"
	"cloudprobe/internal/nc"
	"cloudprobe/internal/pos"
	"sync"
)

type Job interface {
	Get()
}

type Task interface {


	Handler()
	IcmpHandler()
	NcHandler()
	PosHandler()
}

type server struct {
	// for icmp
	icmpScriptList    []string
	icmpArgsList      []string

	// for nc
	ncScList          []string
	ncArgsList        []string

	//for pos
	posScList         []string
}

func NewServer(icmpSl, icmpAl, ncSL, ncAl, posSl []string) Task {
	return &server{
		icmpScriptList: icmpSl,
		icmpArgsList:   icmpAl,
		ncScList: ncSL,
		ncArgsList: ncAl,
		posScList: posSl,
	}
}

func (s *server) Handler() {
	ch := make(chan bool)

	go oneJob(ch, s.IcmpHandler)
	go oneJob(ch, s.NcHandler)
	go oneJob(ch, s.PosHandler)

	<- ch
	<- ch
	<- ch
	log.Info("job goruntine finished !")
}

func (s *server) IcmpHandler() {
	multiGroup := multi.NewConCurTasks()
	wg := &sync.WaitGroup{}

	for i:=0; i<len(s.icmpScriptList); i++ {
		for m:=0; m<len(s.icmpArgsList); m++ {
			wg.Add(1)
			script := s.icmpScriptList[i]
			arg := s.icmpArgsList[m]

			taskFunc := func() {
				icmp.NewIcmp(script, arg).Get()
				wg.Done()
			}
			multiGroup.Run(taskFunc)
		}
	}
	wg.Wait()
}

func (s *server) NcHandler() {
	multiGroup := multi.NewConCurTasks()
	wg := &sync.WaitGroup{}

	for i:=0; i<len(s.ncScList); i++ {
		for m:=0; m<len(s.ncArgsList); m++ {
			wg.Add(1)

			sc  := s.ncScList[i]
			arg := s.ncArgsList[m]

			taskFunc := func() {
				nc.NewNc(sc, arg).Get()
				wg.Done()
			}
			multiGroup.Run(taskFunc)
		}
	}
	wg.Wait()
}

func (s *server) PosHandler() {
	multiGroup := multi.NewConCurTasks()
	wg := &sync.WaitGroup{}

	for i:=0; i<len(s.posScList); i++ {

		sc := s.posScList[i]

		taskFunc := func() {
			pos.NewPos(sc)
			wg.Done()
		}
		multiGroup.Run(taskFunc)
	}

	wg.Wait()
}


func oneJob(ch chan bool, f func()) {

	f()
	ch <- true
}

