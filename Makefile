HEAD=$(shell git describe --tags 2> /dev/null  || git rev-parse --short HEAD)

all: build

build: darwin64 linux64 windows64

clean:
	-rm -f gifdump
	-rm -rf release

darwin64:
	env GOOS=darwin GOARCH=amd64 go clean -i
	env GOOS=darwin GOARCH=amd64 go build -o release/darwin64/gifdump .

linux64:
	env GOOS=linux GOARCH=amd64 go clean -i
	env GOOS=linux GOARCH=amd64 go build -o release/linux64/gifdump .

windows64:
	env GOOS=windows GOARCH=amd64 go clean -i
	env GOOS=windows GOARCH=amd64 go build -o release/windows64/gifdump.exe .

.PHONY: release
release: clean build
	zip -9 release/gifdump.darwin_amd64.$(HEAD).zip release/darwin64/gifdump
	zip -9 release/gifdump.linux_amd64.$(HEAD).zip release/linux64/gifdump
	zip -9 release/gifdump.windows_amd64.$(HEAD).zip release/windows64/gifdump.exe
