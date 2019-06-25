package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd *cobra.Command

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ts",
		Short: "timestamp convert & compare tool",
		Example: `
	(timestamp)	$: ts 
	(format)	$: ts -f "2019/06/25 23:30:10"
	(before)	$: ts -b "2019/06/25 23:30:10" ; echo $?
	(after)		$: ts -a "2019/06/25 23:30:10" ; echo $?
	(timezone)	$: ts -f "2019/06/25 23:30:10" -z "Asia/Shanghai"	
	`,
		Run: func(cmd *cobra.Command, args []string) {
			//pipe stdin
			if len(args) == 0 {
				info, err := os.Stdin.Stat()
				exitForErr(err)

				if info.Mode()&os.ModeNamedPipe != 0 {
					d, err := ioutil.ReadAll(stdin)
					exitForErr(err)
					args = append(args, string(d))
				}
			}
			exitForErr(Main(cmd, args))
		},
	}
	cmd.Flags().StringP("after", "a", "", "after compare")
	cmd.Flags().StringP("before", "b", "", "before compare")
	cmd.Flags().StringP("format", "f", "", "time format")
	cmd.Flags().StringP("timezone", "z", "", "time zone")
	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd = RootCmd()
	viper.BindPFlags(rootCmd.Flags())
	rootCmd.HelpFunc()
}
