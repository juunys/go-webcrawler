package common

import (
	"fmt"
	"time"
)

func SleepBar() {
	count := 0
	percentProgress := 0
	sleepProgress := make([]string, 20)
	for k := 0; k < 20; k++ {
		sleepProgress[k] = "_"
	}

	fmt.Printf("\r%v ==============> %d%%", sleepProgress, percentProgress)

	for {
		if count == 20 {
			break
		}
		time.Sleep(3 * time.Minute)
		percentProgress += 5
		sleepProgress[count] = "#"
		fmt.Printf("\r%v ==============> %d%%", sleepProgress, percentProgress)
		count++
	}
}
