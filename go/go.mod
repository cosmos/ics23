module github.com/cosmos/ics23/go

go 1.21

require (
	github.com/cosmos/gogoproto v1.4.11
	golang.org/x/crypto v0.16.0
)

require (
	github.com/google/go-cmp v0.5.9 // indirect
	golang.org/x/exp v0.0.0-20230811145659-89c5cff77bcb // indirect
	golang.org/x/sys v0.15.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

// subject to the dragonberry vulnerability
retract [v0.0.0, v0.7.0]
