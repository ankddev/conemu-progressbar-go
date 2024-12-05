package conemu_progressbar_go

import "fmt"

// ClearProgress clears progress via sending sequence with 0 state
func ClearProgress() {
	fmt.Printf("\x1b]9;4;0;\x1b\\")
}

// SetIndeterminateProgress sets indeterminate progress state
func SetIndeterminateProgress() {
	fmt.Printf("\x1b]9;4;3;\x1b\\")
}

// SetProgress sets progress with default state
// * progress is a number between 0 and 100, returns error if number isn't in this range
func SetProgress(progress int) error {
	if progress > 100 || progress < 0 {
		return fmt.Errorf("progress must be in range 0..100, while it is %d", progress)
	}
	fmt.Printf("\x1b]9;4;1;%d\x1b\\", progress)
	return nil
}

// SetErrorProgress sets progress with error state
// * progress is a number between 0 and 100, returns error if number isn't in this range
func SetErrorProgress(progress int) error {
	if progress > 100 || progress < 0 {
		return fmt.Errorf("progress must be in range 0..100, while it is %d", progress)
	}
	fmt.Printf("\x1b]9;4;2;%d\x1b\\", progress)
	return nil
}

// SetWarningProgress sets progress with warning state
// * progress is a number between 0 and 100, returns error if number isn't in this range
func SetWarningProgress(progress int) error {
	if progress > 100 || progress < 0 {
		return fmt.Errorf("progress must be in range 0..100, while it is %d", progress)
	}
	fmt.Printf("\x1b]9;4;4;%d\x1b\\", progress)
	return nil
}
