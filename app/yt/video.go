package yt

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Video contains information essential to downloading a video from YouTube.
type Video struct {
	ID      string
	Streams []stream
}

type stream struct {
	Itag    int
	Type    string
	Quality string
	URL     string
}

// NewVideo instantiates a new Video struct
func NewVideo(url string) (vid *Video, err error) {
	var (
		id      string
		info    string
		streams []stream
	)
	if id, err = getVideoID(url); err != nil {
		return
	}
	if info, err = fetchVideoInfo(id); err != nil {
		return
	}
	if streams, err = getVideoStreams(info); err != nil {
		return
	}

	vid = &Video{id, streams}
	return
}

func getVideoID(url string) (id string, err error) {
	id = strings.TrimPrefix(url, "https://www.youtube.com/watch?v=")

	if strings.ContainsAny(id, "?&/<%=") {
		err = errors.New("Invalid characters in video id")
	}
	if len(id) < 10 {
		err = errors.New("Video ID must be at least 10 characters long")
	}

	return
}

func fetchVideoInfo(id string) (string, error) {
	url := "https://youtube.com/get_video_info?video_id=" + id

	VerbosePrint(fmt.Sprintf("video info URL: %s\n", url))

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), nil
}

// getVideoStreams parses the video_info API response and returns an array of video stream structs.
func getVideoStreams(videoInfo string) ([]stream, error) {
	var streams []stream

	videoInfoParams, err := url.ParseQuery(videoInfo)
	if err != nil {
		return streams, err
	}

	if videoInfoParams.Get("errorcode") != "" || videoInfoParams.Get("status") == "fail" {
		return streams, errors.New(videoInfoParams.Get("reason"))
	}

	streamMetadata := strings.Split(videoInfoParams.Get("url_encoded_fmt_stream_map"), ",")

	for i, streamURL := range streamMetadata {
		streamParams, err := url.ParseQuery(streamURL)
		if err != nil {
			log.Printf("An error occured while decoding one of the video's stream urls: stream %d: %s\n", i, err)
		}

		itag, _ := strconv.Atoi(streamParams.Get("itag"))

		streams = append(streams, stream{
			Itag:    itag,
			Type:    streamParams.Get("type"),
			Quality: streamParams.Get("quality"),
			URL:     streamParams.Get("url"),
		})

		VerbosePrint(fmt.Sprintf("Stream found: quality '%s', format '%s'\n", streamParams.Get("quality"), streamParams.Get("type")))
	}

	return streams, nil
}
