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
	"enode://0d75e3d5e6017c26fc53934fbaa8305d731bf0574b35022dd4a7b2bea1e7a34f2131439adc3db31cf1a37c8a9be900fbd8e0a4c0a019ff7eaa650963e64c35f1@40.121.105.44:30310", // BOOTNODE PORT Node 1
	"enode://d8f8a7cfe05231426431f489430e880d6fb7f425d49dc1a95abef589c3c6b541d280dcdf1c6ad12b0d343771b909dd021a85d7d47e751269854d57f5cf1be4d1@35.168.54.62:30311",  // Node 2
	// "enode://d5e4956e254c067893394a9b89924275c3696c2e254b8cc74e0694718f4b657bf64427e98776c1675fb40188ebe2426170ba758dd367f87e5a46fc8cdbc2d783@104.42.157.173:30310", // Node 3
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
