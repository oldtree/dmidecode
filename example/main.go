package main

import (
	"fmt"
	"os"

	"github.com/yumaojun03/dmidecode"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	dmi, err := dmidecode.New()
	checkError(err)

	infos, err := dmi.BIOS()
	dmi.ALL()
	checkError(err)

	for i := range infos {
		fmt.Println(infos[i])
	}
}
