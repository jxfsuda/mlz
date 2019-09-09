package timeUtils

import (
	"strconv"
	"syscall"
	"time"
)

func TimespecToInt(ts syscall.Timespec)  int {
	t:=time.Unix(int64(ts.Sec), int64(ts.Nsec))
	str := t.Format(Format_LocalDateTimeMillionSecond)
	s,err:= strconv.Atoi(str)
	if err!=nil{
		panic(err)
	}
	return s
}