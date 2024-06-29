# NetConfig
A client server service that can :
1. Update host name
2. Add DNS-server
3. Remove DNS-server
# Usage
1. Start a server with `go build internal/server.go`
2. Build client by using `go build main.go`

## Supported commands:
1. `update 'hostname'`
2. `add 'dns'`
3. `delete 'dns'`
