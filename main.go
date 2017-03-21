package main

import (
	"fmt"
	"git.sshaman.ru/sshaman/i_am_init/pac"
)

func init() {
	fmt.Println("Call main init()")
}

func main() {
	fmt.Println("Call main main()")
	pac.TheExportableOne()
}
