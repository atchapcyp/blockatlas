# Block Atlas by Trust Wallet

[![Build Status](https://dev.azure.com/TrustWallet/Trust%20BlockAtlas/_apis/build/status/TrustWallet.blockatlas?branchName=master)](https://dev.azure.com/TrustWallet/Trust%20BlockAtlas/_build/latest?definitionId=27&branchName=master)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/43834b0c94ad4f6088629aa3e3bb5e94)](https://www.codacy.com/app/TrustWallet/blockatlas?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=TrustWallet/blockatlas&amp;utm_campaign=Badge_Grade)
[![Go Report Card](https://goreportcard.com/badge/TrustWallet/blockatlas)](https://goreportcard.com/report/TrustWallet/blockatlas)

Clean explorer API and events observer for crypto currencies.

#### Supported Coins

<a href="https://binance.com" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/714.png" width="32" /></a>
<a href="https://nimiq.com" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/242.png" width="32" /></a>
<a href="https://ripple.com" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/144.png" width="32" /></a>
<a href="https://stellar.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/148.png" width="32" /></a>
<a href="https://kin.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/2017.png" width="32" /></a>
<a href="https://tezos.com" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/1729.png" width="32" /></a>
<a href="https://aion.network" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/425.png" width="32" /></a>
<a href="https://ethereum.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/60.png" width="32" /></a>
<a href="https://ethereumclassic.github.io" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/61.png" width="32" /></a>
<a href="https://poa.network" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/178.png" width="32" /></a>
<a href="https://callisto.network" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/820.png" width="32" /></a>
<a href="https://gochain.io" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/6060.png" width="32" /></a>
<a href="https://wanchain.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/5718350.png" width="32" /></a>
<a href="https://thundercore.com" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/1001.png" width="32" /></a>
<a href="https://icon.foundation" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/74.png" width="32" /></a>
<a href="https://tron.network" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/195.png" width="32" /></a>
<a href="https://vechain.org/" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/818.png" width="32" /></a>
<a href="https://www.thetatoken.org/" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/500.png" width="32" /></a>
<a href="https://cosmos.network/" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/118.png" width="32" /></a>
<a href="https://semux.org/" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/7562605.png" width="32" /></a>
<a href="https://bitcoin.org/" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/0.png" width="32" /></a>

## Setup

### Quick start

Deploy it in less than 30 seconds!

### Prerequisite
* [GO](https://golang.org/doc/install) `1.12+`
* Locally running [Redis](https://redis.io/topics/quickstart) or url to remote instance (required for Observer only)

#### From Source (Go Toolchain required)

```shell
go get -u github.com/trustwallet/blockatlas
cd blockatlas

// Start API server
go build -o blockatlas ./cmd && ./blockatlas api

//Start Observer
go build -o blockatlas ./cmd && ./blockatlas observer
```

#### Docker

Using Docker Hub:

`docker run -it -p 8420:8420 trustwallet/blockatlas`

Build and run from local Dockerfile:
`docker build -t blockatlas .`
`docker run -p 8420:8420 blockatlas`


#### Tools

-   Setup Redis

```shell
brew install redis // Install Redis using Homebrew
```

```shell
ln -sfv /usr/local/opt/redis/*.plist ~/Library/LaunchAgents  // Enable Redis autostart
```

-   Running in the IDE ( GoLand )

1.  Run
2.  Edit configuration
3.  New Go build configuration
4.  Select `directory` as configuration type
5.  Set `api` as program argument and `-i` as Go tools argument 

## Deploy

#### Supported platforms

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://www.heroku.com/deploy/?template=https://github.com/TrustWallet/blockatlas)

[![Docker](https://img.shields.io/docker/cloud/build/trustwallet/blockatlas.svg?style=for-the-badge)](https://hub.docker.com/r/trustwallet/blockatlas)

Block Atlas can run just fine without configuration.

If you want to use custom RPC endpoints, or enable coins without public RPC (like Nimiq),
you can configure Block Atlas over `config.yml` or environment variables.

#### Config File

By default, `config.yml` is loaded from the working directory.

Example (`config.yml`):

```yaml
nimiq:
  api: http://localhost:8648
#...
```

#### Environment

The rest gets loaded from the environment variables.
Every config option is available under the `ATLAS_` prefix.

Example:

```shell
ATLAS_NIMIQ_API=http://localhost:8648 \
blockatlas
```

## Tests

### Unit
To run the unit tests: `make test`

### Integration
All integration tests are generated automatically. You only need to set the environment to your coin in the config file.
The tests use a different build constraint, named `integration`.

To run the integration tests: `make integration` 

or you can run manually: `TEST_CONFIG=$(TEST_CONFIG) TEST_COINS=$(TEST_COINS) go test -tags=integration -v ./pkg/integration`

##### Fixtures

- If you need to change the parameters used in our tests, update the file `pkg/integration/testdata/fixtures.json`

- To exclude an API from integration tests, you need to add the route inside the file `pkg/integration/testdata/exclude.json`

    E.g.:
```
[
  "/v2/ethereum/collections/:owner",
  "/v2/ethereum/collections/:owner/collection/:collection_id"
]
```


## Authors

-   [Richard Patel](https://github.com/terorie)

## Contributing

If you'd like to add support for a new blockchain, feel free to file a pull request.
Note that most tokens that run on top of other chains are already supported and
don't require code changes (e.g. ERC-20).

The best way to submit feedback and report bugs is to open a GitHub issue.
Please be sure to include your operating system, version number, and
[steps](https://gist.github.com/nrollr/eb24336b8fb8e7ba5630) to reproduce reported bugs.
