package cmd

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/araddon/dateparse"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/x-mod/errors"
)

var stdin io.Reader
var stderr io.Writer
var stdout io.Writer

func exitForErr(err error) {
	if err != nil {
		stderr.Write([]byte(err.Error()))
		os.Exit(errors.ValueFrom(err))
	}
}

func Main(cmd *cobra.Command, args []string) error {
	//timezone
	if len(viper.GetString("timezone")) > 0 {
		loc, err := time.LoadLocation(viper.GetString("timezone"))
		if err != nil {
			panic(err.Error())
		}
		time.Local = loc
	}
	//times
	times := make([]time.Time, 0, len(args)+1)
	if len(args) == 0 {
		times = append(times, time.Now())
	}
	for _, arg := range args {
		t, err := dateparse.ParseStrict(arg)
		exitForErr(err)
		times = append(times, t)
	}

	//before compare
	if len(viper.GetString("before")) > 0 {
		t, err := dateparse.ParseStrict(viper.GetString("before"))
		exitForErr(err)
		if t.After(times[0]) {
			os.Exit(1)
		}
		os.Exit(0)
	}

	//after compare
	if len(viper.GetString("after")) > 0 {
		t, err := dateparse.ParseStrict(viper.GetString("after"))
		exitForErr(err)
		if t.Before(times[0]) {
			os.Exit(1)
		}
		os.Exit(0)
	}

	//convert
	for _, tm := range times {
		// ANSIC       = "Mon Jan _2 15:04:05 2006"
		// UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
		// RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
		// RFC822      = "02 Jan 06 15:04 MST"
		// RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
		// RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
		// RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
		// RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
		// RFC3339     = "2006-01-02T15:04:05Z07:00"
		// RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
		// Kitchen     = "3:04PM"
		// // Handy time stamps.
		// Stamp      = "Jan _2 15:04:05"
		// StampMilli = "Jan _2 15:04:05.000"
		// StampMicro = "Jan _2 15:04:05.000000"
		// StampNano  = "Jan _2 15:04:05.000000000"
		dest := fmt.Sprint(time.Now().UnixNano() / 1000000)
		if len(viper.GetString("format")) > 0 {
			switch viper.GetString("format") {
			case "ANSIC":
				dest = time.ANSIC
			case "UnixDate":
				dest = time.UnixDate
			case "RubyDate":
				dest = time.RubyDate
			case "RFC822":
				dest = time.RFC822
			case "RFC822Z":
				dest = time.RFC822Z
			case "RFC850":
				dest = time.RFC850
			case "RFC1123":
				dest = time.RFC1123
			case "RFC1123Z":
				dest = time.RFC1123Z
			case "RFC3339":
				dest = time.RFC3339
			case "RFC3339Nano":
				dest = time.RFC3339Nano
			case "Kitchen":
				dest = time.Kitchen
			case "Stamp":
				dest = time.Stamp
			case "StampMilli":
				dest = time.StampMilli
			case "StampMicro":
				dest = time.StampMicro
			case "StampNano":
				dest = time.StampNano
			default:
				d, err := dateparse.ParseFormat(viper.GetString("format"))
				exitForErr(err)
				dest = d
			}
		}
		fmt.Fprintln(stdout, tm.Format(dest))
	}
	return nil
}

func init() {
	stdin = os.Stdin
	stderr = os.Stderr
	stdout = os.Stdout
}
