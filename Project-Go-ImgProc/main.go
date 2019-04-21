package main

import (
	"flag"
	"fmt"
	"time"

	"Training.go/Project-Go-ImgProc/filter"
	"Training.go/Project-Go-ImgProc/task"
)

func main() {
	var srcDir = flag.String("src", "", "Input Directory")
	var dstDir = flag.String("dst", "", "Output Directory")
	var filterType = flag.String("filter", "grayscale", "grayscale/blur")
	var taskType = flag.String("task", "waitgrp", "waitgrp/channel")
	var poolSize = flag.Int("poolsize", 4, "Pool size")
	flag.Parse()

	var f filter.Filter = filter.Grayscale{}
	switch *filterType {
	case "grayscale":
		f = filter.Grayscale{}
	case "blur":
		f = filter.Blur{}
	}

	var t task.Tasker
	switch *taskType {
	case "waitgrp":
		t = task.NewWaitGrpTask(*srcDir, *dstDir, f)
	case "channel":
		t = task.NewChanTask(*srcDir, *dstDir, f, *poolSize)

	}

	start := time.Now()
	t.Process()
	elapsed := time.Since(start)
	fmt.Printf("Image processing took %s", elapsed)
}
