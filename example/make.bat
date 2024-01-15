set GOARCH=amd64
set GOOS=linux
set EXT=
set FLAGS=-ldflags "-s -w"

go build cmdline.go
go build configfile.go
