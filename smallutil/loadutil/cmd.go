package loadutil

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var TopN int
var LoadPsCmd = &cobra.Command{
	Use:   "loadps",
	Short: "从高到低展示活跃的进程，默认展示10个",
	Long:  `long introduce`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		ret, err := LoadPs()

		if err != nil {
			logrus.Error(err)
			fmt.Println(err)
		}

		cnt := TopN
		if cnt > len(ret) {
			cnt = len(ret)
		}

		fmt.Println("***************** load ps start *******************")
		logrus.Infof("***************** load ps start *******************")
		for i := 0; i < cnt; i++ {
			logrus.Infof(ret[i].String())
			fmt.Println(ret[i].String())
		}
		fmt.Println("***************** load ps end *******************")
		logrus.Infof("***************** load ps end *******************")
	},
}

func init() {
	LoadPsCmd.PersistentFlags().IntVarP(&TopN, "topN", "n", 10, "load ps top nn")
}
