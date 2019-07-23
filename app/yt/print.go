package yt

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// VerbosePrint prints messages when the CLI is run with the "verbose" flag.
func VerbosePrint(message string) {
	if Verbose {
		log.Println(message)
	}
}

// PrintDownloadStats prints the download duration and speed.
func PrintDownloadStats(start time.Time, size int64) {
	duration := time.Now().Sub(start)
	speed := float64(size) / float64(duration/time.Second)
	if duration > time.Second {
		duration -= duration % time.Second
	} else {
		speed = float64(size)
	}

	fmt.Printf("\n\nDownload duration: %s\n", duration)
	fmt.Printf("Average speed: %s/s\n", abbr(int64(speed)))
}

// PrintProgress prints the current download progress.
func PrintProgress(out *os.File, offset, size int64, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	tail := offset
	var err error

	for range ticker.C {
		offset, err = out.Seek(0, os.SEEK_CUR)
		if err != nil {
			return
		}

		percent := 100 * offset / size
		speed := offset - tail
		progress := fmt.Sprintf(
			"%s/%s (%d%%) | %s/s",
			abbr(offset), abbr(size), percent, abbr(speed))
		fmt.Printf("\rProgress: %s", progress)

		tail = offset
		if tail >= size {
			break
		}
	}
}

func abbr(byteSize int64) string {
	const (
		KB float64 = 1 << (10 * (iota + 1))
		MB
		GB
	)

	size := float64(byteSize)

	switch {
	case size > GB:
		return fmt.Sprintf("%.1fGB", size/GB)
	case size > MB:
		return fmt.Sprintf("%.1fMB", size/MB)
	case size > KB:
		return fmt.Sprintf("%.1fKB", size/KB)
	}
	return fmt.Sprintf("%d", byteSize)
}
