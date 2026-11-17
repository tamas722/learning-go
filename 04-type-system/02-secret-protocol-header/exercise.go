package secretprotocolheader

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// createPublishFixHeader constructs an octet (8-bit long byte) based on its three arguments and the fix QoS setting.
func createPublishFixHeader(isFirstAttempt bool, isBroadcasted bool, isSecure bool) byte {
	// INSERT YOUR CODE HERE

	var header byte = (2 << 5) | (2 << 2)

	if isFirstAttempt {
		header |= 1 << 4
	}
	if isBroadcasted {
		header |= 1 << 1
	}
	if isSecure {
		header |= 1 << 0
	}

	return header
}
