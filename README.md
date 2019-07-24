# yt-audio-dl

`yt-audio-dl` is an application and command-line interface for downloading YouTube videos and
converting them to mp3 files.

## Installation

Prerequisites:

-   Go version 1.11 or later
-   FFmpeg https://ffmpeg.org/

To install with `go get`:

```
go get -u github.com/mjlaufer/yt-audio-dl
```

To clone this repository and install:

```
git clone https://github.com/mjlaufer/yt-audio-dl
cd yt-audio-dl
make
```

## Usage

```
yt-audio-dl [options] [youtube url]
```

### CLI Options

-   `--verbose, -v`
