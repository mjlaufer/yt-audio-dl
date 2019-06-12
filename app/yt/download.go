package yt

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

// Verbose stores the "verbose" CLI flag value.
var Verbose bool

// Options contains CLI flag values.
type Options struct {
	Verbose bool
}

// Download is the main application function.
func Download(url string, options *Options) error {
	if options.Verbose {
		Verbose = true
	}

	fmt.Printf("Downloading from %s\n", url)

	// Instantiate a Video struct.
	vid, err := NewVideo(url)
	if err != nil {
		return fmt.Errorf("Error: %s", err)
	}

	usr, _ := user.Current()
	outputDir := filepath.Join(usr.HomeDir, "Downloads")
	outputFile := vid.ID + ".mp4"
	destination := filepath.Join(outputDir, outputFile)

	// Loop through video streams in descending order of quality.
	for _, stream := range vid.Streams {
		VerbosePrint(fmt.Sprintf("Download url: %s\n", stream.URL))
		VerbosePrint(fmt.Sprintf("Downloading to %s\n", destination))

		err = vid.makeRequest(destination, stream.URL)
		if err == nil {
			break
		}
	}

	return err
}

func (vid *Video) makeRequest(destination string, url string) error {
	var (
		out    *os.File
		err    error
		length int64
	)

	// Get download start time
	start := time.Now()

	// GET request to video stream url
	resp, err := http.Get(url)
	if err != nil {
		VerbosePrint(fmt.Sprintf("http.Get\nerror: %s\nURL: %s\n", err, url))
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		VerbosePrint(fmt.Sprintf("Status code: '%v'\n", resp.StatusCode))
		return errors.New("fetch video failed")
	}

	// Copy data to destination file
	out, err = os.Create(destination)
	if err != nil {
		return err
	}

	if length, err = io.Copy(out, resp.Body); err != nil {
		return err
	}

	PrintDownloadStats(start, length)

	return err
}
