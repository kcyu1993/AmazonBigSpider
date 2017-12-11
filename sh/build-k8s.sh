#!/bin/sh
go build -ldflags "-s -w" -x -o /workspace/spiders/usa/UIP /workspace/spiders/usa/ippool.go
go build -ldflags "-s -w" -x -o /workspace/spiders/uk/UIP /workspace/spiders/uk/ippool.go
go build -ldflags "-s -w" -x -o /workspace/spiders/jp/UIP /workspace/spiders/jp/ippool.go
go build -ldflags "-s -w" -x -o /workspace/spiders/de/UIP /workspace/spiders/de/ippool.go

go build -ldflags "-s -w" -x -o /workspace/spiders/usa/UURL /workspace/spiders/usa/urlpool.go
go build -ldflags "-s -w" -x -o /workspace/spiders/uk/UURL /workspace/spiders/uk/urlpool.go
go build -ldflags "-s -w" -x -o /workspace/spiders/jp/UURL /workspace/spiders/jp/urlpool.go
go build -ldflags "-s -w" -x -o /workspace/spiders/de/UURL /workspace/spiders/de/urlpool.go

go build -ldflags "-s -w" -x -o /workspace/spiders/usa/ULIST /workspace/spiders/usa/listmain.go
go build -ldflags "-s -w" -x -o /workspace/spiders/uk/ULIST /workspace/spiders/uk/listmain.go
go build -ldflags "-s -w" -x -o /workspace/spiders/jp/ULIST /workspace/spiders/jp/listmain.go
go build -ldflags "-s -w" -x -o /workspace/spiders/de/ULIST /workspace/spiders/de/listmain.go

go build -ldflags "-s -w" -x -o /workspace/spiders/usa/UASIN /workspace/spiders/usa/asinmain.go
go build -ldflags "-s -w" -x -o /workspace/spiders/uk/UASIN /workspace/spiders/uk/asinmain.go
go build -ldflags "-s -w" -x -o /workspace/spiders/jp/UASIN /workspace/spiders/jp/asinmain.go
go build -ldflags "-s -w" -x -o /workspace/spiders/de/UASIN /workspace/spiders/de/asinmain.go


go build -ldflags "-s -w" -x -o /workspace/spiders/usa/USQL /workspace/spiders/usa/initsql.go
go build -ldflags "-s -w" -x -o /workspace/spiders/uk/USQL /workspace/spiders/uk/initsql.go
go build -ldflags "-s -w" -x -o /workspace/spiders/jp/USQL /workspace/spiders/jp/initsql.go
go build -ldflags "-s -w" -x -o /workspace/spiders/de/USQL /workspace/spiders/de/initsql.go