package common

import (
	"fmt"
	"runtime"
	"strings"
	"unicode"
)

func assert(err error) {
	if err != nil {
		panic(err)
	}
}

type exception []string

func (e exception) Error() string {
	return strings.Join(e, "\n")
}

func throw(msg string, args ...interface{}) {
	panic(trace(msg, args...))
}

func trace(msg string, args ...interface{}) (logs exception) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	msg = strings.TrimRightFunc(msg, unicode.IsSpace)
	if len(msg) > 0 {
		logs = exception{msg}
	}
	n := 1
	for {
		n++
		pc, file, line, ok := runtime.Caller(n)
		if !ok {
			break
		}
		f := runtime.FuncForPC(pc)
		name := f.Name()
		if strings.HasPrefix(name, "runtime.") {
			continue
		}
		fn := file[strings.Index(file, "/src/")+5:]
		logs = append(logs, fmt.Sprintf("\t(%s:%d) %s", fn, line, name))
	}
	return
}

func catch(err *error, handler ...func()) {
	if e := recover(); e != nil {
		*err = e.(error)
	}
	for _, h := range handler {
		h()
	}
}
