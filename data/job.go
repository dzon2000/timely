package data

import (
    "fmt"
    "strconv"
    "github.com/dzon2000/timely/color"
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
        return fmt.Sprintf("[%s>%s]     %-10s \"%s\" running for %d seconds", color.GREEN, color.RESET, j.Tag, j.Desc, (time.Now().Unix() - j.Time))
    }
    return fmt.Sprintf("%7d %-10s \"%s\"", j.Time, j.Tag, j.Desc)
}

func (j Job) AsArray() []string {
    job := make([]string, 4)
    job[0] = strconv.FormatInt(j.Time, 10)
    job[1] = j.Tag
    job[2] = j.Desc
    job[3] = strconv.FormatBool(j.IsRunning)
    return job
}
