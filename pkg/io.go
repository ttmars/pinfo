package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IO struct {
	Rchar               int // 传递给read的字节数之和
	Wchar               int // 传递给write的字节数之和
	Syscr               int // read系统调用次数
	Syscw               int // write系统调用次数
	ReadBytes           int // 从存储层读取的字节数之和
	WriteBytes          int // 写入存储层的字节数之和
	CancelledWriteBytes int // 取消写入的字节数
}

func (i IO) String() string {
	return fmt.Sprintf(`IO
Rchar:			%v Mb
Wchar:			%v Mb
Syscr:			%v
Syscw:			%v
ReadBytes:		%v Mb
WriteBytes:		%v Mb
CancelledWriteBytes:	%v Mb
`, i.Rchar/1048576, i.Wchar/1048576, i.Syscr, i.Syscw, i.ReadBytes/1048576, i.WriteBytes/1048576, i.CancelledWriteBytes/1048576)
}

func ParseIO(pid int) (*IO, error) {
	f, err := os.Open(fmt.Sprintf("/proc/%v/io", pid))
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

	var s IO
	s.Rchar, _ = strconv.Atoi(strings.TrimSpace(m["rchar"]))
	s.Wchar, _ = strconv.Atoi(strings.TrimSpace(m["wchar"]))
	s.Syscr, _ = strconv.Atoi(strings.TrimSpace(m["syscr"]))
	s.Syscw, _ = strconv.Atoi(strings.TrimSpace(m["syscw"]))
	s.ReadBytes, _ = strconv.Atoi(strings.TrimSpace(m["read_bytes"]))
	s.WriteBytes, _ = strconv.Atoi(strings.TrimSpace(m["write_bytes"]))
	s.CancelledWriteBytes, _ = strconv.Atoi(strings.TrimSpace(m["cancelled_write_bytes"]))

	return &s, nil
}
