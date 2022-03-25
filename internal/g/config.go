package g

import (
	"github.com/alecthomas/kingpin"
	"sync"
)

const version = "1.000"


var (
	NcResMap      sync.Map
	IcmpResMap    sync.Map
	PosResMap     sync.Map
)


var (
	// icmp
	IcmpArgs    []string  // max_value
	IcmpScripts []string  // 1.py

	// nc
	NcScripts   []string  // nc.py
	NcArgs      []string  // hostname_port

	// pos
	PosScripts  []string  // 3.py
)


var (

	ListenAndPort = kingpin.Flag("listen_addr", "cloud probe Check listen addr").Default(":9400").String()

	//icmp
	IcmpScriptFilePath     = kingpin.Flag("icmp_script_path", "ICMP cloud Script Path in cloud machine").Default("/no").String()
	ReadIcmpScriptListPath = kingpin.Flag("read_icmp_script_list_path", "read Icmp scritps list path").Default("./no").String()
	ReadIcmpArgsListPath   = kingpin.Flag("read_icmp_args_list_path", "read Icmp scripts list path").Default("./no").String()

	//nc
	NcScriptFilePath       = kingpin.Flag("nc_script_path", "nc script path in cloud machine").Default("./no").String()
	ReadNcScriptListPath   = kingpin.Flag("read_nc_script_list_path", "nc scripts list path").Default("./no").String()
	ReadNcArgsListPath     = kingpin.Flag("read_nc_script_list_path", "nc args list path").Default("./no").String()

	//pos
	PosScriptFilePath      = kingpin.Flag("pos_script_path", "script name list path").Default("./no").String()
	ReadPosScriptListPath  = kingpin.Flag("read_pos_script_list_path", "pos scripts list path").Default("./no").String()


	ConcurrTasks           = kingpin.Flag("concurrent_tasks_number", "concurrent to do tasks").Default("100").Int()
	//NameSpace = kingpin.Flag("prom_namespace", "select CDN type, like tencentCDN; aliCDN; paCDN").Default("paCDN").String()
)

func ParseConfig() {
	kingpin.Version(version)
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

}