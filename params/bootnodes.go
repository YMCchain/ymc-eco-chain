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

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://ac124ff65f4ffa8c489185bd883ebeae465f8c2cfdbbd7aebca7a4a60462e46ea1c05d0c7b98e5c592443f169830bc876d3aa984801f613090f1de129bf03005@18.167.115.215:32668",
	"enode://d65f47de6c00424904a48e732cca611e2294e14e0235287914deb654748b78f600f6ea25f0317d82e4b124deb59236227ab80fc1e8b120d3d0d2a3e25196b16c@18.167.87.140:32668",
	"enode://4890c3c2d474872fd00115036b6999ac92eeff3bd94e14b57a585052f8327bd7910088499deb7295488623de1a1d936be7a9b54d9e10d4d3861754879835245a@18.166.47.128:32668",
	"enode://dd0dbb89b92f86009f8bd9c9cee5a9f830a135b57980905ab0a93f04d892f8cf189e9e14414c49e11706ac69684ef7c87dc40fdb7bb736cc3296c111e4af0c69@18.167.87.121:32668",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
var TestnetBootnodes = []string{
	"enode://924543a43d18bc5759a8bdcd17fa9c7c35df63968e9333640b80b58dab94b17a012371c9d46bed10ce7508a607cac76828ca04685893958eee44ade83b856dc2@47.242.237.63:32668",
	"enode://ebad898d980b520ef6adb54ffb6a68117686e7332f1ea01f7551b7a296a34dd945445a078d7cad019d864c5ef0e0b7f2b5777d94f93adf7dc59f798af72609ac@47.242.235.121:32668",
}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
