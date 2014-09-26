// GPL v3 or higher

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func check(path string, finfo os.FileInfo, err error) error {
	var mode = finfo.Mode()
	var exe bool = (mode & 0111) != 0
	var buf = make([]byte, 12)
	if exe && !mode.IsDir() {
		f, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		n, err := f.Read(buf)
    f.Close()
		if err != nil || n < 12 {
			fmt.Println(err)
			return nil
		}
		line := string(buf)
		var shell string
		if strings.HasPrefix(line, "#!/bin/bash") {
			shell = "bash"
		}
		if strings.HasPrefix(line, "#!/bin/sh") {
			shell = "sh  "
		}
		var danger string = "     "
		if mode&(os.ModeSetuid|os.ModeSetgid) != 0 {
			danger = "setid"
		}
    if shell != "" {
	  	fmt.Println(danger, shell, path)
    }
	}
	return nil
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			"fac143: Scan for executable shell or bash scripts.\n"+
				"Usage: %s [dir1] [dir2]...\n", os.Args[0])
	}
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		return
	}
	for _, root := range flag.Args() {
		fmt.Printf("Checking %s\n", root)
		err := filepath.Walk(root, check)
		if err != nil {
			log.Println(err)
		}
	}
}
