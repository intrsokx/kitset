package loadutil

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type ProcessStat struct {
	Pid       string `json:"pid"`
	ThreadCnt int    `json:"thread_cnt"`
	RCnt      int    `json:"R"`
	DCnt      int    `json:"D"`
	SCnt      int    `json:"S"`
	TCnt      int    `json:"T"`
	ZCnt      int    `json:"Z"`
	XCnt      int    `json:"X"`
	ICnt      int    `json:"I"`
}

func (p *ProcessStat) String() string {
	ret := fmt.Sprintf("PID: %s, R: %d, D: %d, S: %d, ThreadCnt: %d", p.Pid, p.RCnt, p.DCnt, p.SCnt, p.ThreadCnt)
	return ret
}

func (p *ProcessStat) ParseStatus(status string) {
	p.ThreadCnt++
	switch status {
	case "R":
		p.RCnt++
	case "S":
		p.SCnt++
	case "D":
		p.DCnt++
	case "T":
		p.TCnt++
	case "Z":
		p.ZCnt++
	case "X":
		p.XCnt++
	case "I":
		p.ICnt++
	default:
		logrus.Errorf("process status illegal, pid: %s, status: %s", p.Pid, status)
	}
}

type ProcessStatSlice []*ProcessStat

func (s ProcessStatSlice) Len() int {
	return len(s)
}
func (s ProcessStatSlice) Less(i, j int) bool {
	if s[i].RCnt != s[j].RCnt {
		return s[i].RCnt > s[j].RCnt
	} else if s[i].DCnt != s[j].DCnt {
		return s[i].DCnt > s[j].DCnt
	} else {
		return s[i].SCnt > s[j].SCnt
	}
}

func (s ProcessStatSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
