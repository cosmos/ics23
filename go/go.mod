module github.com/cosmos/ics23/go

go 1.19

require (
	github.com/cosmos/gogoproto v1.4.3
	golang.org/x/crypto v0.2.0
)

// subject to the dragonberry vulnerability
retract [v0.0.0, v0.7.0]
