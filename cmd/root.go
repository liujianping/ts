package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//RootCmd constructor
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
	(Formats)		$: ts -F 
		ANSIC       = "Mon Jan _2 15:04:05 2006"
		UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
		RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
		RFC822      = "02 Jan 06 15:04 MST"
		RFC822Z     = "02 Jan 06 15:04 -0700" RFC822 with numeric zone
		RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
		RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
		RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" RFC1123 with numeric zone
		RFC3339     = "2006-01-02T15:04:05Z07:00"
		RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
		Kitchen     = "3:04PM"
		Stamp      = "Jan _2 15:04:05"
		StampMilli = "Jan _2 15:04:05.000"
		StampMicro = "Jan _2 15:04:05.000000"
		StampNano  = "Jan _2 15:04:05.000000000"
		TimestampSec	= "time.Unix()"
		TimestampMilli  = "time.UnixNano()/1000000"
		TimestampMicro  = "time.UnixNano()/1000"
		TimestampNano	= "time.UnixNano()"
	`,
		Run: func(cmd *cobra.Command, args []string) {
			exitForErr(Main(cmd, args))
		},
	}
	cmd.Flags().BoolP("version", "v", false, "current version")
	cmd.Flags().StringP("after", "a", "", "after compare")
	cmd.Flags().StringP("before", "b", "", "before compare")
	cmd.Flags().StringP("format", "f", "TimestampMilli", "time format")
	cmd.Flags().StringP("timezone", "z", "", "time zone")
	cmd.Flags().BoolP("Formats", "F", false, "show formats ?")
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
