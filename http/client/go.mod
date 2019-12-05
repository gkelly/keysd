module github.com/keys-pub/keysd/http/client

go 1.12

require (
	github.com/keys-pub/keys v0.0.0-20191205223248-af81f4ce20b7
	github.com/keys-pub/keysd/http/api v0.0.0-00010101000000-000000000000
	github.com/keys-pub/keysd/http/server v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.4.0
)

replace github.com/keys-pub/keysd/http/api => ../api

replace github.com/keys-pub/keysd/http/server => ../server
