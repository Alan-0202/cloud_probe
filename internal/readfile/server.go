package readfile

import "cloudprobe/internal/g"

type server struct {

}

func NewServer() *server {
	return &server{}
}

func (s *server) Do() {

	// for icmp
	g.IcmpScripts = NewArgs(*g.ReadIcmpScriptListPath).Read().Res
	g.IcmpArgs    = NewArgs(*g.ReadIcmpArgsListPath).Read().Res

	// for nc
	g.NcArgs      = NewArgs(*g.ReadNcArgsListPath).Read().Res
	g.NcScripts   = NewArgs(*g.ReadNcScriptListPath).Read().Res

	//for pos
	g.PosScripts  = NewArgs(*g.ReadPosScriptListPath).Read().Res

}
