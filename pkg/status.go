package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 进程的状态信息
type Status struct {
	Name    string
	Umask   string
	State   string
	Pid     string
	PPid    string
	FDSize  string
	Threads string
}

func (s Status) String() string {
	return fmt.Sprintf(
		`Status
Name:			%v	
Umask:			%v
State:			%v
Pid:			%v
PPid:			%v
FDSize:			%v
Threads:		%v
`, s.Name, s.Umask, s.State, s.Pid, s.PPid, s.FDSize, s.Threads)
}

func ParseStatus(pid int) (*Status, error) {
	f, err := os.Open(fmt.Sprintf("/proc/%v/status", pid))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scan := bufio.NewScanner(f)
	var sli []string
	m := make(map[string]string)
	for scan.Scan() {
		sli = strings.Split(scan.Text(), ":")
		if len(sli) != 2 {
			continue
		}
		m[sli[0]] = sli[1]
	}

	var s Status
	s.Name = strings.TrimSpace(m["Name"])
	s.State = strings.TrimSpace(m["State"])
	s.Umask = strings.TrimSpace(m["Umask"])
	s.Pid = strings.TrimSpace(m["Pid"])
	s.PPid = strings.TrimSpace(m["PPid"])
	s.FDSize = strings.TrimSpace(m["FDSize"])
	s.Threads = strings.TrimSpace(m["Threads"])
	return &s, nil
}
