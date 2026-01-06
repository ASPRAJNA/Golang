package utils

import "fmt"

func Log(message string) {
	fmt.Println(message)
}

//✔️ Capital L → exported
//❌ lowercase → private to package