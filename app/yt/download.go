package yt

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
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
		out  *os.File
		err  error
		size int64
	)

	// Create output file
	out, err = os.Create(destination)
	if err != nil {
		return err
	}

	if size, err = fetchVideoContentLength(url); err != nil {
		VerbosePrint(fmt.Sprintf("http.Head\nerror: %s\nURL: %s\n", err, url))
	}

	if size > 0 {
		go PrintProgress(out, 0, size)
	}

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
	if size, err = io.Copy(out, resp.Body); err != nil {
		return err
	}

	PrintDownloadStats(start, size)

	return err
}

func fetchVideoContentLength(url string) (size int64, err error) {
	res, err := http.Head(url)
	if err != nil {
		err = fmt.Errorf("Head request failed: %s", err)
		return
	}
	if res.StatusCode == 403 {
		err = errors.New("Head request failed with status code 403")
		return
	}

	contentLength := res.Header.Get("Content-Length")
	if len(contentLength) == 0 {
		err = errors.New("Content-Length header is missing")
		return
	}

	size, err = strconv.ParseInt(contentLength, 10, 64)
	if err != nil {
		err = fmt.Errorf("Invalid Content-Length: %s", err)
	}
	return
}
