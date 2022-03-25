package nc

import (
	"cloudprobe/internal/dopy"
	"cloudprobe/internal/g"
	"cloudprobe/internal/tasks"
	"fmt"
)

type Nc struct {
	fPath string
	sName string
	args  string
}

func NewNc(sName,arg string) tasks.Job{
	return &Nc{
		fPath: *g.NcScriptFilePath,
		sName: sName,
		args: arg,
	}
}

func(nc *Nc) Get() {
	res, err := dopy.HandlerPyWithArgs(nc.fPath,nc.sName,nc.args)
	//metricFlag := utils.SplitPy(nc.sName)  // nc.py.py ==> nc.
	if err != nil {
		g.NcResMap.Store(nc.args, map[string]string{
			"value": "-1001",
			"status": fmt.Sprintf("Failed: %v", err),
		})
	}

	g.NcResMap.Store(nc.args, map[string]string{
		"value": res,
		"status": "ok",
	})
	return
}


