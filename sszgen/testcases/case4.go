package testcases

//go:generate go run ../main.go --path case4.go --exclude-objs Version,Root

// Version is a fork version.
type Version [4]byte

// Root is a merkle root.
type Root [32]byte

type ForkData struct {
	// Current version is the current fork version.
	CurrentVersion Version `ssz-size:"4"`

	// GenesisValidatorsRoot is the hash tree root of the validators at genesis.
	GenesisValidatorsRoot Root `ssz-size:"32"`
}
