# go-waku
#### Waku is a family of robust, censorship-resistant, peer-to-peer communication protocols that enable privacy-focused messaging for Web3 applications, allowing you to integrate decentralised communication features into your dApp without compromising security or privacy
A Go implementation of the [Waku v2 protocol](https://rfc.vac.dev/spec/10).


<p align="left">
  <a href="https://goreportcard.com/report/github.com/waku-org/go-waku"><img src="https://goreportcard.com/badge/github.com/waku-org/go-waku" /></a>
  <a href="https://godoc.org/github.com/waku-org/go-waku"><img src="http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square" /></a>
  <a href=""><img src="https://img.shields.io/badge/golang-%3E%3D1.20.0-orange.svg?style=flat-square" /></a>
  <a href="https://codeclimate.com/github/waku-org/go-waku/maintainability"><img src="https://api.codeclimate.com/v1/badges/426bdff6a339ff4d536b/maintainability" /></a>
  <br>
</p>

## Install

#### Building from source
```
git clone https://github.com/waku-org/go-waku
cd go-waku
make

# See the available command line options with
./build/waku --help
```
#### Nix
You can build Waku v2 node using [Nix](https://nixos.org/) [Flakes](https://nixos.wiki/wiki/Flakes):
```sh
nix build github:waku-org/go-waku
```
Or build the library using:
```
nix build github:waku-org/go-waku#library
```
To start a shell with build dependencies use:
```
nix develop
```

#### Docker
```
docker run -i -t -p 60000:60000 -p 9000:9000/udp \
  wakuorg/go-waku:latest \ 
    --dns-discovery \
    --dns-discovery-url enrtree://ANEDLO25QVUGJOUTQFRYKWX6P4Z4GKVESBMHML7DZ6YK4LGS5FC5O@prod.wakuv2.nodes.status.im \
    --discv5-discovery
```
or use the [image:tag](https://hub.docker.com/r/wakuorg/go-waku/tags) of your choice.

or build and run the image with:

```
docker build -t wakuorg/go-waku:latest .

docker run wakuorg/go-waku:latest --help
```

#### Building on windows

Windows requires the following tools to be installed
- git bash  (which is installed as part of [Git](https://git-scm.com/downloads))
- [chocolatey](https://chocolatey.org/install)
- [make](https://community.chocolatey.org/packages/make)
- [mingw](https://community.chocolatey.org/packages/mingw)
- [go](https://go.dev/doc/install)

## Library
```
go get github.com/waku-org/go-waku
```

## C Bindings
```
make static-library
make dynamic-library
```

## Mobile libraries
Requires [`gomobile`](https://pkg.go.dev/golang.org/x/mobile/cmd/gomobile)
```
make mobile-android
make mobile-ios
```

## Tutorials and documentation
- [Receive and send messages using Waku Relay](docs/api/relay.md)
- [Send messages using Waku Lightpush](docs/api/lightpush.md)
- [Encrypting and decrypting Waku Messages](docs/api/encoding.md)
- [Retrieve message history using Waku Store](docs/api/store.md)
- [C Bindings](library/c/README.md)
- [Waku Specs](https://rfc.vac.dev/spec), has information of [waku topics](https://rfc.vac.dev/spec/23/), wakuv1/[wakuv2](https://rfc.vac.dev/spec/14/) message, [rln relay](https://rfc.vac.dev/spec/58/) etc.
- [Enr](https://eips.ethereum.org/EIPS/eip-778), [Enrtree](https://eips.ethereum.org/EIPS/eip-1459)
- [devp2p](https://github.com/ethereum/go-ethereum/tree/master/cmd/devp2p) tool for playing with enr/entree sync tree. [Tutorial](https://geth.ethereum.org/docs/developers/geth-developer/dns-discovery-setup)

## Examples
Examples of usage of go-waku as a library can be found in the `examples/` folder:

- [**basic2**](examples/basic2) - demonstrates how to send and receive messages
- [**chat2**](examples/chat2) - simple chat client using waku relay / lightpush + filter / store protocol to send/receive messages and retrieve message history
- [**filter2**](examples/filter2) - demonstrates how to use filter protocol
- [**c-bindings**](examples/c-bindings) - simple program to demonstrate how to consume the go-waku library via C FFI
- [**waku-csharp**](examples/waku-csharp) - C# console application that uses the go-waku library via FFI
- [**android-kotlin**](examples/android-kotlin) - android app that uses a .jar generated by gomobile using kotlin


## Contribution
Thank you for considering to help out with the source code! We welcome contributions from anyone on the internet, and are grateful for even the smallest of fixes!

If you'd like to contribute to go-waku, please fork, fix, commit and send a pull request. If you wish to submit more complex changes though, please check up with the core devs first to ensure those changes are in line with the general philosophy of the project and/or get some early feedback which can make both your efforts much lighter as well as our review and merge procedures quick and simple.

To build and test this repository, you need:
  - [Go](https://golang.org/) (version 1.20)
  - [protoc](https://grpc.io/docs/protoc-installation/) 
  - [protoc-gen-go](https://protobuf.dev/getting-started/gotutorial/#compiling-protocol-buffers)

To enable the git hooks:

```bash
git config core.hooksPath hooks
```

## Bugs, Questions & Features

If you encounter any bug or would like to propose new features, feel free to [open an issue](https://github.com/waku-org/go-waku/issues/new/).

For more general discussion, help and latest news,  join **#go-waku** on [Vac Discord](https://discord.com/channels/864066763682218004/865466710924460034) or [Telegram](https://t.me/vacp2p).


## License
Licensed and distributed under either of

* MIT license: [LICENSE-MIT](LICENSE-MIT) or http://opensource.org/licenses/MIT

or

* Apache License, Version 2.0, ([LICENSE-APACHEv2](LICENSE-APACHEv2) or http://www.apache.org/licenses/LICENSE-2.0)

at your option. These files may not be copied, modified, or distributed except according to those terms.
