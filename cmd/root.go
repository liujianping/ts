package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ts",
		Short: "timestamp convert & compare tool",
		Example: `
	(now timestamp)	$: ts
	(now add)		$: ts --add 1d
	(now sub)		$: ts --sub 1d
	(convert)		$: ts "2019/06/24 23:30:10"
	(pipe)			$: echo "2019/06/24 23:30:10" | ts		
	(format)		$: ts -f "2019/06/25 23:30:10"
	(before)		$: ts -b "2019/06/25 23:30:10" ; echo $?
	(after)			$: ts -a "2019/06/25 23:30:10" ; echo $?
	(timezone)		$: ts -f "2019/06/25 23:30:10" -z "Asia/Shanghai"	
	`,
		Run: func(cmd *cobra.Command, args []string) {
			exitForErr(Main(cmd, args))
		},
	}
	cmd.Flags().BoolP("version", "v", false, "current version")
	cmd.Flags().StringP("after", "a", "", "after compare")
	cmd.Flags().StringP("before", "b", "", "before compare")
	cmd.Flags().StringP("format", "f", "", "time format")
	cmd.Flags().StringP("timezone", "z", "", "time zone")
	cmd.Flags().DurationP("add", "", 0*time.Second, "add duration")
	cmd.Flags().DurationP("sub", "", 0*time.Second, "sub duration")
	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd := RootCmd()
	exitForErr(viper.BindPFlags(rootCmd.Flags()))
	rootCmd.HelpFunc()
	exitForErr(rootCmd.Execute())
}
