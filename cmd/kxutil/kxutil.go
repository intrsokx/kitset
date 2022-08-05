package main

import (
	"errors"
	"fmt"
	"github.com/intrsokx/kitset/smallutil/diff"
	"github.com/intrsokx/kitset/smallutil/loadutil"
	"github.com/intrsokx/kitset/smallutil/parseip"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	logfile := "kxutil.log"
	var f *os.File

	_, err := os.Stat(logfile)
	if errors.Is(err, os.ErrNotExist) {
		f, err = os.Create(logfile)
	} else {
		f, err = os.OpenFile(logfile, os.O_RDWR|os.O_APPEND, 0666)
	}
	if err != nil {
		panic(err)
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	//logrus.SetReportCaller(true)
	logrus.SetOutput(f)
}

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "kx util",
	Long:  `kx util`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		logrus.Info("rootCmd run", args)
		fmt.Println("rootCmd run: ", args)
	},
}

func Execute() {
	//add sub command
	rootCmd.AddCommand(loadutil.LoadPsCmd)
	rootCmd.AddCommand(parseip.ParseIpCmd)
	rootCmd.AddCommand(diff.DiffFileCmd)

	//execute
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
