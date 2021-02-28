package data

import (
    "fmt"
    "strconv"
    "github.com/dzon2000/timely/color"
    "github.com/dzon2000/timely/format"
    "time"
)

type Job struct {
    Time int64
    Tag string
    Desc string
    IsRunning bool
}

func (j Job) String() string {
    if j.IsRunning {
        return fmt.Sprintf("[%s>%s]%s %-10s \"%s\"", color.GREEN, color.RESET, format.FormatSec((time.Now().Unix() - j.Time)), j.Tag, j.Desc)
    }
    return fmt.Sprintf("%11s %-10s \"%s\"", format.FormatSec(j.Time), j.Tag, j.Desc)
}

func (j Job) AsArray() []string {
    job := make([]string, 4)
    job[0] = strconv.FormatInt(j.Time, 10)
    job[1] = j.Tag
    job[2] = j.Desc
    job[3] = strconv.FormatBool(j.IsRunning)
    return job
}
