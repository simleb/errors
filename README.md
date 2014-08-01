# Simple errors stacking package for go

Install:

	go get github.com/simleb/errors

Import:

	import "github.com/simleb/errors"

Example use:

	if err != nil {
		return errors.Stack(err, "one more error: %d", 42)
	}

