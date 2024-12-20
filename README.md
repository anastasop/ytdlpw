# ytdlpw

`ytdlpw` is a wrapper around [yt-dlp](https://github.com/yt-dlp/yt-dlp) with some reasonable defaults,
to set formats, names etc. It is a personalized tool to create a media library, with files playable
by all your media players (_your_ phone, _your_ media center). Initially this was a shell script
but converted to go to make things easier for multiple environments.

## Examples

```
$ ytdlpw -a -q https://www.youtube.com/watch?v=eVli-tstM5E
$ ytdlpw -i -q -w 1024 https://www.youtube.com/watch?v=ubJ7iqLkXfs
$ ls
Sabrina_Carpenter_-_Espresso_Official_Video.m4a
We_Need_to_Talk_About_Physics_-_with_Helen_Czerski.info.json
We_Need_to_Talk_About_Physics_-_with_Helen_Czerski.mp4
```

## Installation

Needs [`go`](https://golang.org/dl), [`yt-dlp`](https://github.com/yt-dlp/yt-dlp) and [`ffmpeg`](https://ffmpeg.org/download.html).
Install them using your package manager and then:

`go install github.com/anastasop/ytdlpw@latest`

The executable `ytdlpw` should be at `${GOBIN}` which should be in your `PATH` alongside `yt-dlp` and `ffmpeg`.

## Bugs

- pass `yt-dlp` options and make this a proper wrapper
- distinguish overwrite errors
