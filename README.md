# Go Nil Slice Iteration Bug
This example demonstrates a common error in Go when iterating over a nil slice using the `range` keyword.  Attempting to modify elements within the slice will result in a runtime panic.