package conemu_progressbar_go

// Example for setting progress successfully
func ExampleSetProgress() {
	err := SetProgress(80)
	if err != nil {
		return
	}

	// Output: ]9;4;1;80\
}

// Example for setting progress with error status successfully
func ExampleSetErrorProgress() {
	err := SetErrorProgress(20)
	if err != nil {
		return
	}

	// Output: ]9;4;2;20\
}

// Example for setting progress with warning status successfully
func ExampleSetWarningProgress() {
	err := SetWarningProgress(30)
	if err != nil {
		return
	}

	// Output: ]9;4;4;30\
}

// Example for setting progress to indeterminate
func ExampleSetIndeterminateProgress() {
	SetIndeterminateProgress()

	// Output: ]9;4;3;\
}

// Example for clearing progress
func ExampleClearProgress() {
	ClearProgress()

	// Output: ]9;4;0;\
}
