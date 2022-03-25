package icmp

import (
	"cloudprobe/internal/dopy"
	"cloudprobe/internal/g"
	"cloudprobe/internal/tasks"
	"cloudprobe/utils"
	"fmt"
)

type Icmp struct {
	fPath  string
	sName  string
	args   string
}

func NewIcmp(sn, args string) tasks.Job {
	return &Icmp{
		fPath: *g.IcmpScriptFilePath,
		sName: sn,
		args: args,
	}
}

// python icmp_bgp_nat.py max_value
func (i *Icmp) Get() {
	res, err := dopy.HandlerPyWithArgs(i.fPath,i.sName, i.args)

	name := utils.SplitPy(i.sName)
	if err != nil {
		g.IcmpResMap.Store(fmt.Sprintf("%v[%v]", name, i.args), map[string]string{
			"value": "-1001",
			"status": fmt.Sprintf("Failed: %v", err),
		})
	}

	// ok
	g.IcmpResMap.Store(fmt.Sprintf("%v[%v]", name, i.args), map[string]string{
		"value": res,
		"status": "ok",
	})
	return
}
