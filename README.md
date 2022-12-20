# Stash Video Hashes Generator

A small tool that generates Stash and StashBox compatible PHash and OSHash hashes.  
Generation uses the [stashapp/stash](https://github.com/stashapp/stash) implementation.

## [Download latest from releases](https://github.com/peolic/videohashes/releases/latest)

## Usage
```
$ ./videohashes.exe '/path/to/video file.mp4'

Duration: 00:10:34 (634)
PHash:    de83a3120f3eb2ac
OSHash:   bd3a5dba480425b9

```

### Calculate MD5 as well:
```
$ ./videohashes.exe -md5 '/path/to/video file.mp4'

Duration: 00:10:34 (634)
PHash:    de83a3120f3eb2ac
OSHash:   bd3a5dba480425b9
MD5:      3b69ed4e44e73435665b51b7dd989d16

```

### JSON Output:
```
$ ./videohashes.exe -json '/path/to/video file.mp4'
{
  "duration": 634,
  "phash": "de83a3120f3eb2ac",
  "oshash": "bd3a5dba480425b9",
}
```

## Build
```sh
# Windows
make windows-amd64
# Linux
make linux-amd64
# Mac OS
make macos-amd64
```
