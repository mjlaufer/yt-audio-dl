package yt

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Video struct
type Video struct {
	ID      string
	Formats []format
}

type format struct {
	Itag    int
	Type    string
	Quality string
	URL     string
}

// NewVideo instantiates a new video struct
func NewVideo(url string) (vid *Video, err error) {
	var (
		id      string
		info    string
		formats []format
	)
	if id, err = getVideoID(url); err != nil {
		return
	}
	if info, err = fetchVideoInfo(id); err != nil {
		return
	}
	if formats, err = getVideoFormats(info); err != nil {
		return
	}

	vid = &Video{id, formats}
	return
}

func getVideoID(url string) (id string, err error) {
	id = strings.TrimPrefix(url, "https://www.youtube.com/watch?v=")

	if strings.ContainsAny(id, "?&/<%=") {
		err = errors.New("invalid characters in video id")
	}
	if len(id) < 10 {
		err = errors.New("the video id must be at least 10 characters long")
	}

	return
}

func getVideoFormats(videoInfo string) ([]format, error) {
	var formats []format

	values, err := url.ParseQuery(videoInfo)
	if err != nil {
		return formats, err
	}

	if values.Get("errorcode") != "" || values.Get("status") == "fail" {
		return formats, errors.New(values.Get("reason"))
	}

	formatParams := strings.Split(values.Get("url_encoded_fmt_stream_map"), ",")

	for _, f := range formatParams {
		fURL, _ := url.Parse("?" + f)
		fQueryParams := fURL.Query()
		itag, _ := strconv.Atoi(fQueryParams.Get("itag"))

		formats = append(formats, format{
			Itag:    itag,
			Type:    fQueryParams.Get("type"),
			Quality: fQueryParams.Get("quality"),
			URL:     fQueryParams.Get("url") + "&signature=" + fQueryParams.Get("sig"),
		})
	}

	return formats, nil
}

func fetchVideoInfo(id string) (string, error) {
	url := "http://youtube.com/get_video_info?video_id=" + id

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), nil
}
