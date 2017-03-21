package pac

import "fmt"

func init() {
	fmt.Println("Call one init()")
}

func TheExportableOne() {
	fmt.Println("I am Exportalbe!")
}
