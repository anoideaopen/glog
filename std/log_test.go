package std_test

import (
	"log"
	"testing"

	"github.com/newity/glog"
	"github.com/newity/glog/std"
)

func TestLogger(t *testing.T) {
	t.Parallel()

	var l glog.Logger = std.New(log.Default(), std.LevelError).SetLevel(std.LevelTrace)

	l.Set(glog.Field{K: "KEY", V: "VALUE"})

	l.Trace("1", "2", "3")
	l.Tracef("%s,%s", "1", "2")

	l = l.With(glog.Field{K: "KEY1", V: "VALUE1"}, glog.Field{K: "KEY2", V: "VALUE2"})
	l.Debug("1", "2", "3")
	l.Debugf("%s,%s", "1", "2")
	l.Info("1", "2", "3")
	l.Infof("%s,%s", "1", "2")
	l.Warning("1", "2", "3")
	l.Warningf("%s,%s", "1", "2")
	l.Error("1", "2", "3")
	l.Errorf("%s,%s", "1", "2")

	l.(*std.Log).SetLevel(std.LevelInfo)
	l.Debug("1", "2", "3")
	l.Debugf("%s,%s", "1", "2")

	defer func() {
		_ = recover()

		l.Trace("1", "2", "3")
		l.Tracef("%s,%s", "1", "2")
	}()

	defer func() {
		_ = recover()

		l.Panicf("%s,%s", "1", "2")
	}()

	l.Panic("1", "2", "3")
}
