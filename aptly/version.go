package aptly

import "github.com/aptly-dev/aptly/internal"

// Version of aptly (filled in at link time)
var Version = internal.Version()

// EnableDebug triggers some debugging features
const EnableDebug = false
