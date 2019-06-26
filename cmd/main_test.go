package cmd

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/x-mod/errors"
)

func TestVersion(t *testing.T) {
	cmd := RootCmd()
	viper.BindPFlags(cmd.Flags())
	cmd.Flags().Set("version", "true")
	assert.Nil(t, Main(cmd, []string{}))
}

func TestFormats(t *testing.T) {
	cmd := RootCmd()
	viper.BindPFlags(cmd.Flags())
	cmd.Flags().Set("Formats", "true")
	assert.Nil(t, Main(cmd, []string{}))
}

func TestArgs(t *testing.T) {
	cmd := RootCmd()
	viper.BindPFlags(cmd.Flags())
	assert.Nil(t, Main(cmd, []string{}))
}

func TestTimezone(t *testing.T) {
	cmd := RootCmd()
	viper.BindPFlags(cmd.Flags())
	cmd.Flags().Set("timezone", "asdfasdf")
	assert.NotNil(t, Main(cmd, []string{}))

	cmd.Flags().Set("timezone", "Asia/Shanghai")
	assert.Nil(t, Main(cmd, []string{}))
}

func TestAddSub(t *testing.T) {
	cmd := RootCmd()
	viper.BindPFlags(cmd.Flags())
	cmd.Flags().Set("add", "1h")
	assert.Nil(t, Main(cmd, []string{}))
	cmd.Flags().Set("sub", "1h")
	assert.Nil(t, Main(cmd, []string{}))
}

func TestBefore(t *testing.T) {
	cmd := RootCmd()
	viper.BindPFlags(cmd.Flags())
	cmd.Flags().Set("before", "2039adfasdf")
	assert.NotNil(t, Main(cmd, []string{"2019/06/12"}))
	cmd.Flags().Set("before", "2029/12/12")
	e1 := Main(cmd, []string{"2019/06/12"})
	assert.NotNil(t, e1)
	assert.Equal(t, int32(1), errors.ValueFrom(e1), e1.Error())
	e2 := Main(cmd, []string{"2039/06/12"})
	assert.NotNil(t, e2)
	assert.Equal(t, int32(0), errors.ValueFrom(e2), e2.Error())
}

func TestAfter(t *testing.T) {
	cmd := RootCmd()
	viper.BindPFlags(cmd.Flags())
	cmd.Flags().Set("after", "2009fasdfa")
	assert.NotNil(t, Main(cmd, []string{"2019/06/12"}))
	cmd.Flags().Set("after", "2009/01/01")
	e1 := Main(cmd, []string{"2008/06/12"})
	assert.NotNil(t, e1)
	assert.Equal(t, int32(0), errors.ValueFrom(e1), e1.Error())
	e2 := Main(cmd, []string{"2019/06/12"})
	assert.NotNil(t, e2)
	assert.Equal(t, int32(1), errors.ValueFrom(e2), e2.Error())
}

func TestFormat(t *testing.T) {
	cmd := RootCmd()
	viper.BindPFlags(cmd.Flags())
	cmd.Flags().Set("format", "2039sadfadsfa")
	assert.NotNil(t, Main(cmd, []string{"2019/06/12"}))
	cmd.Flags().Set("format", "2039/12/12")
	assert.Nil(t, Main(cmd, []string{"2019/06/12"}))
	cmd.Flags().Set("format", "ANSIC")
	assert.Nil(t, Main(cmd, []string{"2019/06/12"}))
	cmd.Flags().Set("format", "UnixDate")
	assert.Nil(t, Main(cmd, []string{"2019/06/12"}))
	cmd.Flags().Set("format", "RubyDate")
	assert.Nil(t, Main(cmd, []string{"2019/06/12"}))
	cmd.Flags().Set("format", "RFC822")
	assert.Nil(t, Main(cmd, []string{"2019/06/12"}))
	cmd.Flags().Set("format", "RFC822Z")
	assert.Nil(t, Main(cmd, []string{"2019/06/12"}))
	cmd.Flags().Set("format", "RFC850")
	assert.Nil(t, Main(cmd, []string{"2019/06/12"}))
	cmd.Flags().Set("format", "RFC1123")
	assert.Nil(t, Main(cmd, []string{"2019/06/12"}))
	cmd.Flags().Set("format", "RFC1123Z")
	assert.Nil(t, Main(cmd, []string{"2019/06/12"}))
	cmd.Flags().Set("format", "RFC3339")
	assert.Nil(t, Main(cmd, []string{"2019/06/12"}))
	cmd.Flags().Set("format", "RFC3339Nano")
	assert.Nil(t, Main(cmd, []string{"2019/06/12"}))
	cmd.Flags().Set("format", "Kitchen")
	assert.Nil(t, Main(cmd, []string{"2019/06/12"}))
	cmd.Flags().Set("format", "Stamp")
	assert.Nil(t, Main(cmd, []string{"2019/06/12"}))
	cmd.Flags().Set("format", "StampMilli")
	assert.Nil(t, Main(cmd, []string{"2019/06/12"}))
	cmd.Flags().Set("format", "StampMicro")
	assert.Nil(t, Main(cmd, []string{"2019/06/12"}))
	cmd.Flags().Set("format", "StampNano")
	assert.Nil(t, Main(cmd, []string{"2019/06/12"}))
}

func TestCovert(t *testing.T) {
	cmd := RootCmd()
	viper.BindPFlags(cmd.Flags())
	assert.NotNil(t, Main(cmd, []string{"2019asdfafd"}))
}
