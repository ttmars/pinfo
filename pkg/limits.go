package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 进程的资源限制
type Limits struct {
	MaxOpenFiles LimitOptions
	MaxFileLocks LimitOptions
}

type LimitOptions struct {
	SoftLimit string
	HardLimit string
	Units     string
}

func (p Limits) String() string {
	a := p.MaxOpenFiles
	b := p.MaxFileLocks
	return fmt.Sprintf(`Limits
MaxOpenFiles:		%v %v %v
MaxFileLocks:		%v %v %v
`, a.SoftLimit, a.HardLimit, a.Units, b.SoftLimit, b.HardLimit, b.Units)
}

func ParseLimits(pid int) (*Limits, error) {
	f, err := os.Open(fmt.Sprintf("/proc/%v/limits", pid))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var l Limits
	scan := bufio.NewScanner(f)
	var sli []string
	for scan.Scan() {
		sli = strings.Fields(scan.Text())
		if len(sli) > 3 {
			switch sli[1] {
			case "open":
				l.MaxOpenFiles = LimitOptions{Units: sli[len(sli)-1], HardLimit: sli[len(sli)-2], SoftLimit: sli[len(sli)-3]}
			case "file":
				l.MaxFileLocks = LimitOptions{Units: sli[len(sli)-1], HardLimit: sli[len(sli)-2], SoftLimit: sli[len(sli)-3]}
			}
		}
	}

	return &l, nil
}
