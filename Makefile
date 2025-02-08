# This how we want to name the binary output
BINARY=binaryMain

VERSION=`date +%Y%m%d%H%M%S`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION}"

# Remove private gopath information from the compiling machine. e: /home/xxx/work/go/... =>  go/...
RPATHFLAGS=-gcflags "-trimpath=$(GOPATH)" -asmflags "-trimpath=$(GOPATH)"  -trimpath

# Builds the project
app:
	go build -o ../../bin/${BINARY}  ${RPATHFLAGS}  ${LDFLAGS}  ../../cmd/admin/main.go

# Build linux
linux: binary-linux

local: simple

# Cleans our project: deletes binaries
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi


.PHONY: clean

# 编译命令
binary-linux:
	GOOS=linux GOARCH=amd64  go build -o ./generate/${BINARY}-linux  ${RPATHFLAGS}  -ldflags "-w -s"  ./generate/main.go

simple:
	go build -o ./generate/${BINARY}  ${RPATHFLAGS}  -ldflags "-w -s"  ./generate/main.go
