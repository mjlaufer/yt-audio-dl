# yt-audio-dl

`yt-audio-dl` is an application and command-line interface for downloading YouTube videos and
converting them to mp3 files.

## Installation

Prerequisite: Go version 1.11 or later

To install with `go get`:

```
go get -u github/mjlaufer/yt-audio-dl/...
```

To clone this repository and install:

```
git clone https://github.com/mjlaufer/yt-audio-dl
cd https://github.com/mjlaufer/yt-audio-dl
make
```

## Usage

```
yt [options] [youtube url]
```

### CLI Options

-   `--debug, -d` : output debug logs
-   `--help, -h` : show help
-   `--output, -o` : output to specific path
-   `--start-offset` : offset the start of the video by a duration of time (e.g., 20s or 1 min)
-   `--version, -v` : show yt-audio-dl version

### Examples (coming soon)
