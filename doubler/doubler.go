package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	path := strings.Split(os.Args[1], "\\")
	name := path[len(path)-1]
	fmt.Print(name)
	finalname := "doubled-"+name
	cmd :=exec.Command("ffmpeg","-i" , "concat:"+name+"|"+name, "-acodec" , "copy" , finalname)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
