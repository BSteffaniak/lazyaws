package logs

import (
	"fmt"
	"log"
	"os"

	"github.com/aybabtme/humanlog"
	"github.com/BSteffaniak/lazyaws/pkg/config"
)

// TailLogs lets us run `lazyaws --logs` to print the logs produced by other lazyaws processes.
// This makes for easier debugging.
func TailLogs() {
	logFilePath, err := config.LogPath()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Tailing log file %s\n\n", logFilePath)

	opts := humanlog.DefaultOptions
	opts.Truncates = false

	_, err = os.Stat(logFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("Log file does not exist. Run `lazyaws --debug` first to create the log file")
		}
		log.Fatal(err)
	}

	TailLogsForPlatform(logFilePath, opts)
}
