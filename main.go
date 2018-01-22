package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Fatal(err)
		}
		stamp := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("%s\t%s", stamp, input)
	}
}
