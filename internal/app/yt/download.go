package yt

import (
	"fmt"
)

// Download downloads the YouTube video from the provided URL and converts it to audio.
func Download(url string) error {
	fmt.Printf("downloading %s\n", url)

	vid, err := NewVideo(url)
	if err != nil {
		return fmt.Errorf("Error: %s", err)
	}
	fmt.Printf("Video: %+v", vid)

	return err
}
