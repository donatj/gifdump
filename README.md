# gifdump

Dump a gif to it's frames!

## Usage:

```bash
gifdump <filename.gif>
```

The "visible" frame of a gif often contains transparency allowing colors from previous frames to show as a naive form of image compression.
By default gifdump composes the frames for you so you get the output you would expect.

To get the true uncomposed frames, you may use the `-u` option.

## Installation

Binaries are availble on the release page:  https://github.com/donatj/gifdump/releases

### From Source:

```bash
$ go get -u -v github.com/donatj/gifdump
```
