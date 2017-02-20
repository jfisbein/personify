# personify

## Linux Install

1. Install the OpenCV Developer package. On Ubuntu systems that's `sudo apt install libopencv-dev`

2. `go get github.com/rasmi/personify`

3. `go get github.com/lazywei/go-opencv`

4. `cd $GOPATH/github.com/rasmi/personify && go build && go install`

## Usage


Simplest: `./personify --faces /path/to/faces path/to/input.jpg > output.jpg`

If executed from any location besides the repository, you must tell it where to find the
bundled Haar Cascade face recognition XML file. I tried to bundle it with the binary, but
it must be provided as a file to the OpenCV library, so a file path is necessary.

`personify --haar /path/to/haarcascade_frontalface_alt.xml /path/to/input.jpg > output.jpg`
