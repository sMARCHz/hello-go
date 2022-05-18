package section

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{})

func SixthSection() {
	// In this section, we'll talk about select statement
	// The select statement lets a goroutine wait on multiple communication operations.
	// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

	go logger()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{}
}

func logger() {
	for {
		// The entire statement is blocked until a message is received on one of the channels that it's listening (one of these cases)
		select {
		case entry := <-logCh:
			// If we received a message from logCh, select statement will bring you in this block
			fmt.Printf("%v - [%v] %v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			// If we received a message from doneCh, select statement will bring you in this block
			break
		}
	}
}

// We can add default case to select statement.
// If there's a message ready on one of the channels that are being monitored, then it's going to execute that code path.
// If not, it'll execute a default block.

// But keep it in mind that default case make the select statement no longer blocking
// It is useful if you want a non-blocking select statement

// select {
// case entry := <-logCh:
// 		fmt.Printf("%v - [%v] %v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
// case <-doneCh:
// 		break
// default:
// 		fmt.Println("Default")
// }
