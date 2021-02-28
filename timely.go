package main

import (
    "fmt"
    "os"
    "time"
    "github.com/dzon2000/timely/io"
    "github.com/dzon2000/timely/data"
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


func stop() {
    jobs := io.Read()
    anyRunning := false
    for i, job := range jobs {
        if job.IsRunning {
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
        fmt.Println("Nothing is started...")
        return
    }
    io.Write(jobs)
}

func start(tag, desc string) {
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
        fmt.Println("Listing......")
        list()
    case "start":
        fmt.Printf("Adding %s to %s", args[2], args[1])
        start(args[1], args[2])
    case "stop":
        stop()
    }
}
