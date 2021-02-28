package data

import (
    "fmt"
    "strconv"
)

type Job struct {
    Time int64
    Tag string
    Desc string
    IsRunning bool
}

func (j Job) String() string {
    return fmt.Sprintf("%d, %s, %s, Running? %v", j.Time, j.Tag, j.Desc, j.IsRunning)
}

func (j Job) AsArray() []string {
    job := make([]string, 4)
    job[0] = strconv.FormatInt(j.Time, 10)
    job[1] = j.Tag
    job[2] = j.Desc
    job[3] = strconv.FormatBool(j.IsRunning)
    return job
}
