// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package impalaservice

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"status"
	"beeswax"
	"cli_service"

)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = status.GoUnusedProtection__
var _ = beeswax.GoUnusedProtection__
var _ = cli_service.GoUnusedProtection__
var DEFAULT_QUERY_OPTIONS map[TImpalaQueryOptions]string

func init() {
DEFAULT_QUERY_OPTIONS = map[TImpalaQueryOptions]string{
    0: "false",
    1: "0",
    2: "false",
    3: "0",
    5: "0",
    6: "0",
    7: "0",
    8: "0",
    9: "false",
    10: "-1",
    11: "",
    4: "0",
}

}

