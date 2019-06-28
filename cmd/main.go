package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/x-mod/build"
	"github.com/x-mod/errors"
)

type cmpCode struct {
	val int32
}

func (cc *cmpCode) Value() int32 {
	if cc.val != 0 {
		return 1
	}
	return cc.val
}

func (cc *cmpCode) String() string {
	if cc.val == 1 {
		return "True"
	}
	return "False"
}

func exitForErr(err error) {
	if err != nil {
		os.Stdout.Write([]byte(err.Error() + "\n"))
		os.Exit(int(errors.ValueFrom(err)))
	}
}

func printFormats() {
	fmt.Println(`
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
		`)
}

func Main(cmd *cobra.Command, args []string) error {
	//version
	if viper.GetBool("version") {
		fmt.Println(build.String())
		return nil
	}
	//formats
	if viper.GetBool("Formats") {
		printFormats()
		return nil
	}
	//pipe stdin
	if len(args) == 0 {
		info, err := os.Stdin.Stat()
		if err != nil {
			return errors.Annotate(err, "stdin stat failed")
		}
		//OR ModeCharDevice & Size check
		// if (info.Mode()&os.ModeCharDevice) == os.ModeCharDevice &&
		// 	info.Size() > 0 {
		if (info.Mode() & os.ModeNamedPipe) == os.ModeNamedPipe {
			d, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return errors.Annotate(err, "stdin read failed")
			}
			args = append(args, string(d))
		}
	}
	//timezone
	if len(viper.GetString("timezone")) > 0 {
		loc, err := time.LoadLocation(viper.GetString("timezone"))
		if err != nil {
			return err
		}
		time.Local = loc
	}
	//times
	times := make([]time.Time, 0, len(args)+1)
	if len(args) == 0 {
		t := time.Now()
		t = t.Add(viper.GetDuration("add"))
		t = t.Add(-viper.GetDuration("sub"))
		times = append(times, t)
	}
	for _, arg := range args {
		t, err := dateparse.ParseIn(strings.TrimSpace(arg), time.Local)
		if err != nil {
			return errors.Annotate(err, "parse strict")
		}
		t = t.Add(viper.GetDuration("add"))
		t = t.Add(-viper.GetDuration("sub"))
		times = append(times, t)
	}

	//before compare
	if len(viper.GetString("before")) > 0 {
		t, err := dateparse.ParseIn(viper.GetString("before"), time.Local)
		//t, err := dateparse.ParseStrict(viper.GetString("before"))
		if err != nil {
			return errors.Annotate(err, "parse strict")
		}
		if t.After(times[0]) {
			return errors.CodeError(&cmpCode{val: 1})
		}
		return errors.CodeError(&cmpCode{})
	}

	//after compare
	if len(viper.GetString("after")) > 0 {
		t, err := dateparse.ParseIn(viper.GetString("after"), time.Local)
		if err != nil {
			return errors.Annotate(err, "parse strict")
		}
		if t.Before(times[0]) {
			return errors.CodeError(&cmpCode{val: 1})
		}
		return errors.CodeError(&cmpCode{})
	}

	//convert
	for _, tm := range times {
		if len(viper.GetString("format")) > 0 {
			dest := ""
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
			case "TimestampSec":
				fmt.Fprintln(os.Stdout, tm.Unix())
				continue
			case "TimestampMilli":
				fmt.Fprintln(os.Stdout, tm.UnixNano()/1000000)
				continue
			case "TimestampMicro":
				fmt.Fprintln(os.Stdout, tm.UnixNano()/1000)
				continue
			case "TimestampNano":
				fmt.Fprintln(os.Stdout, tm.UnixNano())
				continue
			default:
				d, err := dateparse.ParseFormat(viper.GetString("format"))
				if err != nil {
					return errors.Annotate(err, "parse format")
				}
				dest = d
			}
			fmt.Fprintln(os.Stdout, tm.Format(dest))
			continue
		}
		fmt.Fprintln(os.Stdout, tm.UnixNano()/1000000)
	}
	return nil
}
