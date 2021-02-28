package io

import (
    "encoding/csv"
    "os"
    "fmt"
    "strconv"
    "time"
    "github.com/dzon2000/timely/data"
)

const FILE = ".jobs_list"
var HOME_DIR, _ = os.UserHomeDir()

func Init() {
    _, err := os.Create(HOME_DIR + "/" + FILE)
    if err != nil {
        fmt.Println(err)
        return
    }
}

func Read() []data.Job {
    f, err := os.Open(HOME_DIR + "/" + FILE)
    if err != nil {
        fmt.Println(err)
        return []data.Job{}
    }
    defer f.Close()
    lines, err := csv.NewReader(f).ReadAll()
    if err != nil {
        fmt.Println(err)
    }
    jobs := make([]data.Job, len(lines))
    for i, line := range lines {
        time, _ := strconv.ParseInt(line[0], 10, 64)
        isRunning, _ := strconv.ParseBool(line[3])
        jobs[i] = data.Job{
            Time: time,
            Tag: line[1],
            Desc: line[2],
            IsRunning: isRunning,
        }
    }
    return jobs
}

func Write(jobs []data.Job) {
    f, err := os.OpenFile(HOME_DIR + "/" + FILE, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(err)
        return
    }
    writer := csv.NewWriter(f)
    count := len(jobs)
    content := make([][]string, count)
    for i, job := range jobs {
        content[i] = job.AsArray()
    }
    fmt.Println(content)
    writer.WriteAll(content)
    writer.Flush()
    f.Close()
}

func Append(job data.Job) {
    f, err := os.OpenFile(HOME_DIR + "/" + FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(err)
        return
    }
    writer := csv.NewWriter(f)
    line := []string{fmt.Sprintf("%d", time.Now().Unix()), job.Tag, job.Desc, fmt.Sprintf("%d", 1)}
    writer.Write(line)
    if err != nil {
        fmt.Println(err)
    }
    writer.Flush()
    f.Close()
}
