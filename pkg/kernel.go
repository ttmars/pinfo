package pkg

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Kernel struct {
	Version string
	Uptime  string
}

func (k Kernel) String() string {
	return fmt.Sprintf(`Kernel
Version:		%v
Uptime:			%v
`, k.Version, k.Uptime)
}

func ParseKernel() (*Kernel, error) {
	var k Kernel
	var b []byte
	var sli []string
	b, _ = os.ReadFile("/proc/version")
	k.Version = strings.TrimSpace(string(b))

	b, _ = os.ReadFile("/proc/uptime")
	sli = strings.Fields(string(b))
	if len(sli) == 2 {
		n1, _ := strconv.ParseFloat(sli[0], 10)
		n2, _ := strconv.ParseFloat(sli[1], 10)
		k.Uptime = fmt.Sprintf("%vd%vh %vd%vh", int(n1)/86400, int(n1)%86400/3600, int(n2)/86400, int(n2)%86400/3600)
	}

	return &k, nil
}
