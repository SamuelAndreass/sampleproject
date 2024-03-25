package handle

import "fmt"

func HandleErr(err error) {
	if err != nil {
		fmt.Print(err)
		return
	}
}