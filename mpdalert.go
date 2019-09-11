package main

import (
	"fmt"
	"os"

	"github.com/fhs/gompd/mpd"
)

func main() {
	if len(os.Args) == 1 {
		panic("Usage mpdalert <PLAYLIST> <HOST> <PORT>")
	}
	list := os.Args[1]
	addr := "127.0.0.1"
	if len(os.Args) >= 3 {
		addr = os.Args[2]
	}
	port := "6600"
	if len(os.Args) >= 4 {
		port = os.Args[3]
	}
	cli, err := mpd.Dial("tcp", fmt.Sprintf("%s:%s", addr, port))
	if err != nil {
		panic(err)
	}

	err = cli.PlaylistSave("current")
	if err != nil {
		err = cli.PlaylistRemove("current")
		if err != nil {
			panic(err)
		}
		err = cli.PlaylistSave("current")
		if err != nil {
			panic(err)
		}
	}
	err = cli.Clear()
	if err != nil {
		panic(err)
	}
	err = cli.PlaylistLoad(list, -1, -1)
	if err != nil {
		panic(err)
	}
	err = cli.Play(0)
	if err != nil {
		panic(err)
	}
}
