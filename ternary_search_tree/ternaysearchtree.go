package tst

type (
	TSTNODE struct {
		key             byte
		value           interface{}
		left, eq, right *TSTNODE
	}

	TeranrySearchTree struct {
		length int
		root   *TSTNODE
	}
)
