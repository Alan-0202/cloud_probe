package readfile

import (
	"bufio"
	"healthcheck/log"
	"io"
	"os"
)

type args struct {
	path string
	Res []string
}

func NewArgs(path string) *args {
	return &args{
		path: path,
	}
}

func (hn *args) Read() *args{
	f, err := os.Open(hn.path)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	defer f.Close()

	rds := bufio.NewReader(f)
	for {
		line, err := rds.ReadString('\n')
		if err != nil || io.EOF == err {
			log.Debug(err)
			break
		}
		hn.Res = append(hn.Res, line)
	}
	return hn
}
