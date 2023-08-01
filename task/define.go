package task

const (
	// region delay status
	Wait = iota
	Running
	AfterRunningDone
	AfterRunningErr
	Cancel
	// endregion
)
