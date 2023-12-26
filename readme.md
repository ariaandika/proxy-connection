# Proxy Connection

- if the backend server only support HTTP/1.1, should the proxy also use  
  HTTP/1.1 through the client ?

in nginx, thats the case

- should we even use HTTP/2.0 for server to server inside private network ?

tls and http2 can be unecessary overhead for localhost call

so the goal is to make go serve http2 to the client while having http/1.1 for the backend

opening two connection and having protocol translation is fine, this is called http level proxy

there also a tcp level proxy, where proxy does not care what protocol it uses

## running

```bash
CERT=/cert.pem KEY=/key.pem PORT=3000 TARGET_1=deuzo.me:8000 go run main.go
bun run src/index.ts
cargo run
```

For current configuration, TARGET_1 env is proxy target. The number does not
matter as long as its start with TARGET, this is to provide multiple host

## Subject

go, the easiest reverse proxy server and tls
rust, backend that support HTTP/2.0 via actix
bun, backend with only HTTP/1.1

