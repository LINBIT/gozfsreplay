all: zfsreplay_gen.go

zfsreplay_gen.go: zfsreplay.go
	go tool cgo -godefs $< > $@
