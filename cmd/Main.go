package cmd

import (
	"fmt"
	"os"
)

func main()  {

}

func help()  {
	fmt.Fprintln(os.Stderr, `
	  Usage: 
	  nyumbapoa start		- start the server 
   `)
}
