#onionscan

onionscan is a scanner that will take a list of onions and check each onion to see if it's reachable and report back the result.

Note that this is a work in progress.

#### Install

1. `go get -u github.com/norwack/onionscan`
  + If $GOPATH/bin is in your path, you can run it by executing `onionscan myonionlist.txt`
  +  Will pull in [go-socks as a dependency](https://github.com/samuel/go-socks/).
