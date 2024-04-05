package pkg

import (
	"fmt"
	"os"
	"strings"
)

type Others struct {
	Cmdline string // 完整的启动命令
	Environ string // 环境变量
	Comm    string

	CWD     string // 工作目录
	FDOpens int    // 打开的文件数

	OomScore    string // 分数越大，越容易被kill
	OomAdj      string
	OomScoreAdj string
}

func (o Others) String() string {
	return fmt.Sprintf(`Others
Cmdline:		%v
Environ:		%v
Comm:			%v
CWD:			%v
FDOpens:		%v
OomScore:		%v
OomAdj:			%v
OomScoreAdj:		%v
	`, o.Cmdline, o.Environ, o.Comm, o.CWD, o.FDOpens, o.OomScore, o.OomAdj, o.OomScoreAdj)
}

func ParseOther(pid int) (*Others, error) {
	var s Others

	var b []byte
	b, _ = os.ReadFile(fmt.Sprintf("/proc/%v/cmdline", pid))
	s.Cmdline = strings.TrimSpace(string(b))
	b, _ = os.ReadFile(fmt.Sprintf("/proc/%v/environ", pid))
	s.Environ = strings.TrimSpace(string(b))
	b, _ = os.ReadFile(fmt.Sprintf("/proc/%v/comm", pid))
	s.Comm = strings.TrimSpace(string(b))

	s.CWD, _ = os.Readlink(fmt.Sprintf("/proc/%v/cwd", pid))
	dirEntrys, _ := os.ReadDir(fmt.Sprintf("/proc/%v/fd", pid))
	s.FDOpens = len(dirEntrys)

	// oom info
	b, _ = os.ReadFile(fmt.Sprintf("/proc/%v/oom_score", pid))
	s.OomScore = strings.TrimSpace(string(b))
	b, _ = os.ReadFile(fmt.Sprintf("/proc/%v/oom_adj", pid))
	s.OomAdj = strings.TrimSpace(string(b))
	b, _ = os.ReadFile(fmt.Sprintf("/proc/%v/oom_score_adj", pid))
	s.OomScoreAdj = strings.TrimSpace(string(b))

	return &s, nil
}
