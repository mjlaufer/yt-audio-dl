package yt

import (
	"fmt"
	"log"
	"os"
	"time"
)

// VerbosePrint prints messages when the CLI is run with the "verbose" flag.
func VerbosePrint(message string) {
	if Verbose {
		log.Println(message)
	}
}

// PrintDownloadStats prints the download duration and speed.
func PrintDownloadStats(start time.Time, length int64) {
	duration := time.Now().Sub(start)
	speed := float64(length) / float64(duration/time.Second)
	if duration > time.Second {
		duration -= duration % time.Second
	} else {
		speed = float64(length)
	}

	fmt.Printf("Download duration: %s\n", duration)
	fmt.Printf("Average speed: %s/s\n", abbr(int64(speed)))
}

// PrintProgress prints the download progress bar.
func PrintProgress(out *os.File, offset, length int64) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	start := time.Now()
	tail := offset

	var err error
	for now := range ticker.C {
		duration := now.Sub(start)
		duration -= duration % time.Second
		offset, err = out.Seek(0, os.SEEK_CUR)
		if err != nil {
			return
		}
		speed := offset - tail
		percent := int(100 * offset / length)
		progress := fmt.Sprintf(
			"%s\t %s/%s\t %d%%\t %s/s",
			duration, abbr(offset), abbr(length), percent, abbr(speed))
		fmt.Println(progress)
		tail = offset
		if tail >= length {
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
