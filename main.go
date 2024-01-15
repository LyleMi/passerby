package passerby

import (
    "os"
	"fmt"
	"time"
    "strings"
)

func DeleteFileAfterDuration(filePath string) {
	go deleteFileAfterDuration(filePath)
}

func deleteFileAfterDuration(filePath string) {
	defaultDuration := 60 * time.Second
	durationStr := os.Getenv("DELETE_DURATION")

	if durationStr != "" {
		duration, err := time.ParseDuration(durationStr)
		if err != nil {
			fmt.Printf("Error parsing duration: %v\n", err)
			return
		}
		defaultDuration = duration
	}

	timer := time.NewTimer(defaultDuration)
	defer timer.Stop()

	<-timer.C
	err := os.Remove(filePath)
	if err != nil {
		fmt.Printf("Error deleting file: %v\n", err)
	} else {
		fmt.Printf("File %s has been deleted\n", filePath)
	}
}

func HookArgs() {
    arg := os.Getenv("ARG")
    if arg == "" {
        return
    }

    newArgs := strings.Split(arg, " ")
    os.Args = append([]string{os.Args[0]}, newArgs...)
}

func init() {
	if os.Getenv("PASSERBY_NO_INIT") != "" {
		return
	}
	HookArgs()
}
