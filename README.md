# blockreward

[![Go Reference](https://pkg.go.dev/badge/github.com/RaghavSood/blockreward.svg)](https://pkg.go.dev/github.com/RaghavSood/blockreward)

Convenience library to calculate block rewards for various bitcoin-like cryptocurrencies.

## Usage

```go
package main

import (
    "fmt"
    "github.com/RaghavSood/blockreward"
)

func main() {
    blockHeight := int64(210000)
    bitcoinReward := blockreward.SubsidyAtHeight(blockreward.BitcoinMainnet, blockHeight)
    litecoinReward := blockreward.SubsidyAtHeight(blockreward.LitecoinMainnet, blockHeight)

    fmt.Printf("Bitcoin reward at height %d: %d satoshis\n", blockHeight, bitcoinReward)
    fmt.Printf("Litecoin reward at height %d: %d satoshis\n", blockHeight, litecoinReward)
}
```
