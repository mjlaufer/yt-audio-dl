# yt-audio-dl

`yt-audio-dl` is an application and command-line interface for downloading YouTube videos and
converting them to mp3 files.

## Installation and Usage

### Prerequisites:

-   Go version 1.11 or later
-   FFmpeg https://ffmpeg.org/

To install with `go get`:

```
go get -u github.com/mjlaufer/yt-audio-dl
yt-audio-dl [options] [youtube url]
```

To clone this repository and build:

```
git clone https://github.com/mjlaufer/yt-audio-dl
cd yt-audio-dl
go build -o yt yt-audio-dl.go
./yt [options] [youtube url]
```

### CLI Options:

-   `--verbose, -v`
