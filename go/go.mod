module github.com/cosmos/ics23/go

go 1.22

require (
	github.com/cosmos/gogoproto v1.7.0
	github.com/stretchr/testify v1.8.4
	golang.org/x/crypto v0.31.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// subject to the dragonberry vulnerability
retract [v0.0.0, v0.7.0]
