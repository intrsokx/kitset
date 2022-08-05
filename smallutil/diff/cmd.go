package diff

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strings"
)

var oldFile, newFile string

var DiffFileCmd = &cobra.Command{
	Use:   "diff",
	Short: "diff 文件版本工具",
	Long:  `文件比较工具，输入两个版本文件，输出文件改动信息（增加或减少部分）`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("***************** diff file start *******************")
		logrus.Infof("***************** diff file start *******************")

		addLines, delLines := CompareFile(oldFile, newFile)
		fmt.Printf("addLines: \n%v\n", strings.Join(addLines, "\n"))
		fmt.Printf("delLines: \n%v\n", strings.Join(delLines, "\n"))

		logrus.Infof("addLines: \n%v\n", strings.Join(addLines, "\n"))
		logrus.Infof("delLines: \n%v\n", strings.Join(delLines, "\n"))

		fmt.Println("***************** diff file end *******************")
		logrus.Infof("***************** diff file end *******************")

	},
}

func init() {
	DiffFileCmd.PersistentFlags().StringVarP(&oldFile, "oldFile", "o", "", "old version file")
	DiffFileCmd.PersistentFlags().StringVarP(&newFile, "newFile", "n", "", "new version file")
}
