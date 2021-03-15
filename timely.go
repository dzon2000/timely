package main

import (
    "fmt"
    "os"
    "time"
    "github.com/dzon2000/timely/io"
    "github.com/dzon2000/timely/data"
    "github.com/dzon2000/timely/color"
    "github.com/dzon2000/timely/format"
)

const PROG_NAME = "timely"

func printUsage() {
    fmt.Printf("%s [ACTION]", PROG_NAME)
}

func list() {
    jobs := io.Read()
    for _, job := range jobs {
        fmt.Println(job)
    }
}

func listTag(tag string) {
    jobs := io.Read()
    for _, job := range jobs {
        if job.Tag == tag {
            fmt.Println(job)
        }
    }
}

func stop() {
    jobs := io.Read()
    anyRunning := false
    for i, job := range jobs {
        if job.IsRunning {
            fmt.Printf("Stopping \"%s\" for %s\n", job.Desc, job.Tag)
            start := job.Time
            diff := time.Now().Unix() - start
            jobs[i] = data.Job{
                Time: diff,
                Tag: job.Tag,
                Desc: job.Desc,
                IsRunning: false,
            }
            anyRunning = true
            break
        }
    }
    if !anyRunning {
        fmt.Printf("[%s!%s] Nothing is started.", color.RED, color.RESET)
        return
    }
    io.Write(jobs)
}

func aggregate() {
    jobs := io.Read()
    byTag := make(map[string][]data.Job)
    for _, job := range jobs {
        val, ok := byTag[job.Tag]
        if ok {
            val = append(val, job)
            byTag[job.Tag] = val
        } else {
            byTag[job.Tag] = []data.Job{job}
        }
    }
    for k := range byTag {
        jobs := byTag[k]
        tTime := int64(0)
        for _, job := range jobs {
            tTime += job.Time
        }
        fmt.Printf("%s - %s\n", k, format.FormatSec(tTime))
    }
}

func start(tag, desc string) {
    jobs := io.Read()
    for _, job := range jobs {
        if job.IsRunning {
            fmt.Printf("[%s!%s] \"%s\" is already running.\n", color.RED, color.RESET, job.Desc)
            return
        }
    }
    job := data.Job{
        Tag: tag,
        Desc: desc,
    }
    io.Append(job)
}

func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        printUsage()
        return
    }
    switch args[0] {
    case "init":
        fmt.Println("Creating the file...")
        io.Init()
    case "list":
        if len(args) == 1 {
            fmt.Println("Listing all")
            list()
        } else {
            fmt.Printf("Listing for %s\n", args[1])
            listTag(args[1])
        }
    case "start":
        fmt.Printf("Adding %s to %s\n", args[2], args[1])
        start(args[1], args[2])
    case "stop":
        stop()
    case "aggr":
        fmt.Println("Listing aggregated view")
        aggregate()
    }
}
