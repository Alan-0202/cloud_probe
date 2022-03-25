package pos

import (
	"cloudprobe/internal/dopy"
	"cloudprobe/internal/g"
	"cloudprobe/internal/tasks"
	"cloudprobe/utils"
	"fmt"
)

type Pos struct {
	fPath string
	sName string
}

func NewPos(sName string) tasks.Job {
	return &Pos{
		fPath: *g.PosScriptFilePath,
		sName: sName,
	}
}


// exec : python xxxx.py
func (p *Pos) Get() {
	res, err := dopy.HandlerPyNoArgs(p.fPath, p.sName)
	name := utils.SplitPy(p.sName)
	if err != nil {
		g.PosResMap.Store(name, map[string]string{
			"value": "-1001",
			"status": fmt.Sprintf("Failed: %v", err),
		})
	}
	//if ok
	g.PosResMap.Store(name, map[string]string{
		"value": res,
		"status": "ok",
	})

	return
}
