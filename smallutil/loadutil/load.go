package loadutil

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"path"
	"regexp"
	"sort"
	"strings"
)

func LoadPs() (ProcessStatSlice, error) {
	procDir := `/proc`

	//匹配数字的正则
	reg, err := regexp.Compile(`\d`)
	if err != nil {
		return nil, err
	}

	dirs, err := ioutil.ReadDir(procDir)
	if err != nil {
		return nil, err
	}

	ret := make(ProcessStatSlice, 0)

	for _, dir := range dirs {
		if dir.IsDir() && reg.MatchString(dir.Name()) {
			logrus.Debugf("scan process: %s", dir.Name())

			pid := dir.Name()
			pStat := &ProcessStat{
				Pid: pid,
			}

			taskPath := path.Join(procDir, dir.Name(), "task")

			taskDirs, err := ioutil.ReadDir(taskPath)
			if err != nil {
				logrus.Warn(err)
				continue
			}
			for _, pidDir := range taskDirs {
				statPath := path.Join(taskPath, pidDir.Name(), "stat")

				buf, err := ioutil.ReadFile(statPath)
				if err != nil {
					logrus.Warn(err)
					continue
				}

				//TODO 特殊case处理
				//470801 (PM2 v4.2.3: God) S 467979 470801 470801 0
				str := string(buf)
				left := strings.Index(str, "(")
				right := strings.Index(str, ")")

				dealStr := strings.ReplaceAll(string(str), str[left:right+1], "")

				stat := strings.Split(dealStr, " ")
				//pid := stat[0]
				status := stat[2]

				pStat.ParseStatus(status)
			}

			ret = append(ret, pStat)
		}
	}

	sort.Sort(ret)

	return ret, nil
}
