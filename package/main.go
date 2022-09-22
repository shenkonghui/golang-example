package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gobuffalo/packr/v2/file"

	"github.com/gobuffalo/packr/v2"
)

type File = file.File

func main() {
	PreStr := "tmp/"
	os.Mkdir(PreStr, 0777)
	box := packr.Folder("./..")

	box.Walk(func(path string, info File) error {
		if info == nil {
			return nil
		}
		finfo, _ := info.FileInfo()
		finfo.Name()
		if !finfo.IsDir() {
			index := strings.LastIndexAny(finfo.Name(), "/")
			if index > 1 {
				dirName := finfo.Name()[0:index]
				fmt.Println(PreStr + dirName)
				os.Mkdir(PreStr+dirName, 0777)
			}

			f, err := os.Create(PreStr + finfo.Name())
			defer f.Close()
			if err == nil {
				p, _ := box.Find(finfo.Name())
				_, err = f.Write(p)
			}
		}
		return nil
	})
	fmt.Println("finish")

}
