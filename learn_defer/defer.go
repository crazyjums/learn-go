package learn_defer

import "fmt"

func cleanup() {
	fmt.Println("Performing cleanup...")
}

func DeferOrder() {
	fmt.Println("Start of main function")

	defer fmt.Println("Deferred statement 1")
	defer fmt.Println("Deferred statement 2")
	defer cleanup()

	panic("A panic occurred")

	fmt.Println("End of main function") // This line will never be reached
}
