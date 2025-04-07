module github.com/cosmos/ics23/go

go 1.22
toolchain go1.24.1

require (
	github.com/cosmos/gogoproto v1.7.0
	golang.org/x/crypto v0.37.0
)

require (
	github.com/google/go-cmp v0.6.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)

// subject to the dragonberry vulnerability
retract [v0.0.0, v0.7.0]
