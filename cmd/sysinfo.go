/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cobra-Azure-tools/exportexcel"
	"cobra-Azure-tools/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	host, pwd, username string
	port                int
	command1            string
	excel               bool
)

// sysinfoCmd represents the sysinfo command
var sysinfoCmd = &cobra.Command{
	Use:   "sysinfo",
	Short: "查看系统信息的命令",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(host) == 0 || len(pwd) == 0 {
			cmd.Help()
			return
		}
		fmt.Println("执行命令结果 如下：")

		utils.Sysinfo(host, pwd, username, port, command1, excel)

		if excel {
			//如果 输入的命令是 uptime,以及  --excel true，则会输出一份excel在当前目录
			if command1 == "uptime" {
				exportexcel.ExportVmCpu()
			}

		}

	},
}

func init() {
	rootCmd.AddCommand(sysinfoCmd)
	sysinfoCmd.Flags().StringVarP(&host, "host", "i", "", "host ip addr")
	sysinfoCmd.Flags().StringVarP(&username, "username", "u", "", "host username")
	sysinfoCmd.Flags().StringVarP(&command1, "command1", "c", "", "command1")
	sysinfoCmd.Flags().StringVarP(&pwd, "pwd", "p", "", "host password")
	sysinfoCmd.Flags().IntVarP(&port, "port", "P", 0, "host port")
	sysinfoCmd.Flags().BoolVarP(&excel, "excel", "e", false, "请输入是否需要打印EXCEL")

}
