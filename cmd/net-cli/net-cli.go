package main

import (
	"fmt"
	"log"
	"os"
	
	"github.com/urfave/cli"
)

func main(){
	fmt.Println("net-cli v0.0.10")

	err := cli.NewApp().Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}