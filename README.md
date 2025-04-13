# BitBadges Module

The BitBadges Module is a Cosmos SDK module that implements a innovative token standard for blockchain applications. It provides a comprehensive system for creating, managing, and transferring digital NFTs and tokens on the blockchain.

## Prerequisites

- Go 1.23 or later
- Cosmos SDK v0.50.11 (Only one tested so far)
- CometBFT v0.38.12

## Installation

To use this module in your Cosmos SDK application:

1. Add the module to your go.mod:

```go
require (
	github.com/bitbadges/badges-module latest
)
```

2. Import the module in your app:

```go
import (
	badgeskeeper "github.com/bitbadges/badges-module/x/badges/keeper"
	badges "github.com/bitbadges/badges-module/x/badges/module"
	badgetypes "github.com/bitbadges/badges-module/x/badges/types"
)
```

## Structure

The module is structured as follows:

- `x/badges`: All core logic is here
- `app`: This is just the most basic little dependency injection setup used for testing.
- `api`: API types for the module

## Usage

This is a standard Cosmos SDK module. Apply it as such in your app.go file.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the terms of the license found in the LICENSE file. Please reach out to us for more information on using this module in your blockchain.
