package typedeclimport

import subpkg "github.com/cosmos/gogoproto/test/typedeclimport/subpkg"

type SomeMessage struct {
	subpkg.AnotherMessage
	Imported subpkg.AnotherMessage
}
