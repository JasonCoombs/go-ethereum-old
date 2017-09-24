// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// VentureCurrency Fork of Ethereum - Go Bootnodes
	"enode://c28a44da72a018183f21124ac38fe73a1947ab41d86bce375b99afc238e6f9299cd0084125d228c6c290a8946ce3ab6e9e28192c4135763f7913175472015330@54.241.36.138:13332", // USA-WEST
	"enode://@127.0.0.1:13332", // USA-EAST
	"enode://@127.0.0.1:13332",  // EUROPE-GERMANY
	"enode://@127.0.0.1:13332", // BREXIT-LONDON
	"enode://@127.0.0.1:13332", // JAPAN
	"enode://@127.0.0.1:13332", // AUSTRALIA
	"enode://@127.0.0.1:13332",  // INDIA
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://@127.0.0.1:13332", // CANADA
	"enode://@127.0.0.1:13332", // IRELAND
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://@127.0.0.1:30303", // IE
	"enode://@127.0.0.1:30303",  // INFURA
}

// RinkebyV5Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network for the experimental RLPx v5 topic-discovery network.
var RinkebyV5Bootnodes = []string{
	"enode://@127.0.0.1:30303?discport=30304", // IE
	"enode://@127.0.0.1:30303?discport=30304",  // INFURA
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://@127.0.0.1:30305",
	"enode://@127.0.0.1:30308",
	"enode://@127.0.0.1:30309",
}
