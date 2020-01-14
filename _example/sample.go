package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"

	zfsreplay "github.com/LINBIT/gozfsreplay"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal(os.Args[0], "file.zstream")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

Loop:
	for {
		obj, typ, err := zfsreplay.DRRRead(f, binary.LittleEndian)
		if err != nil {
			log.Fatal(err)
		}

		switch typ {
		case zfsreplay.DRR_BEGIN:
			_ = obj.(zfsreplay.DRR_begin)
			fmt.Println("begin")
		case zfsreplay.DRR_FREEOBJECTS:
			_ = obj.(zfsreplay.DRR_freeobjects)
			fmt.Println("freeobjects")
		case zfsreplay.DRR_OBJECT:
			_ = obj.(zfsreplay.DRR_object)
			fmt.Println("object")
		case zfsreplay.DRR_FREE:
			_ = obj.(zfsreplay.DRR_free)
			fmt.Println("free")
		case zfsreplay.DRR_WRITE:
			_ = obj.(zfsreplay.DRR_write)
			fmt.Println("write")
		case zfsreplay.DRR_END:
			_ = obj.(zfsreplay.DRR_end)
			fmt.Println("end")
			break Loop
		default:
			fmt.Println("DEFAULT main")
			break Loop
		}
	}
}
