package section

import (
	"fmt"
	"runtime"
)

// Go supports Parallelism. GOMAXPROCS keyword is used to tell Go that this program is running in parallelism.
// GOMAXPROCS is used to set the maximum number of operating system threads that can execute user-level Go code simultaneously.

func FourthSection() {
	// Get GOMAXPROCS
	// By default, Go runs program with GOMAXPROCS which sets to the number of cores available
	threads := runtime.GOMAXPROCS(-1)
	fmt.Println(threads)

	// Set GOMAXPROCS
	runtime.GOMAXPROCS(30)
	threads = runtime.GOMAXPROCS(-1)
	fmt.Println(threads)

	// Get maximum number of logical CPUs usable by the current process
	cpus := runtime.NumCPU()
	fmt.Println(cpus)
}
