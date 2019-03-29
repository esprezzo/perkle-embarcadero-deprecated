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
	// Esprezzo Foundation Go Bootnodes
	"enode://9a07e80df806bfacd89b4245f8272ced37ff4985e7db98d7c1a4883cf8c49c89e970d3b28bd2f986e97cace01af9a78e117478d4927b2d65c306a4622a2ff4b0@3.90.100.66:30310",   // BOOTNODE/Blocknet2
	"enode://d8f8a7cfe05231426431f489430e880d6fb7f425d49dc1a95abef589c3c6b541d280dcdf1c6ad12b0d343771b909dd021a85d7d47e751269854d57f5cf1be4d1@35.168.54.62:30311",  // Blocknet 1
	"enode://0757ba56ff97496c32dfc586193f917411c0364b5a4892db79e52e1a36b410430d87767d40afbe6e5c94c4a72d1b59db90fa4262bd52f0ea8d0a4e117cc11390@34.232.227.76:30311", // Blocknet 3
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
