package parseip

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strings"
)

var str string

var ParseIpCmd = &cobra.Command{
	Use:   "parseip",
	Short: "ip解析工具",
	Long:  `ip解析工具: 输入的字符串s, 输出ips`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("***************** parse ip start *******************")
		logrus.Infof("***************** parse ip start *******************")

		ips := ParseIP(str)
		fmt.Println("parse ip success, length: ", len(ips))
		fmt.Println(strings.Join(ips, "\n"))

		logrus.Info("parse ip success, length: ", len(ips))
		logrus.Infof(strings.Join(ips, "\n"))

		fmt.Println("***************** parse ip end *******************")
		logrus.Infof("***************** parse ip end *******************")

	},
}

func init() {
	ParseIpCmd.PersistentFlags().StringVarP(&str, "str", "s", "", "input str")
}
