# saga-agent
This is the LXD agent for creating new LXC

## Dependencies
1. One of glibc, musl libc, uclib or bionic as your C library
2. Linux kernel >= 2.6.32 
3. Go 1.9 : https://golang.org/doc/install
4. LXD 3.0.x : https://linuxcontainers.org/lxd/getting-started-cli/
5. Testify : https://github.com/stretchr/testify
6. LXD client :   https://github.com/lxc/lxd/tree/master/client 
7. gorilla/mux : https://github.com/gorilla/mux

Make sure to set the GOPATH environment variable as stated on Go installation tutorial

## Running the Tests
```
$ go test
```

## Installing the Program
```
$ go install
```
## Running the Program
```
$ saga-agent
```
