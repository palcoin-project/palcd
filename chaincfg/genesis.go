// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"time"

	"github.com/palcoin-project/palcd/chaincfg/chainhash"
	"github.com/palcoin-project/palcd/wire"
)

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network, regression test network, and test network (version 3).
var genesisCoinbaseTx = wire.MsgTx{
	Version: 1,
	TxIn: []*wire.TxIn{
		{
			PreviousOutPoint: wire.OutPoint{
				Hash:  chainhash.Hash{},
				Index: 0xffffffff,
			},
			SignatureScript: []byte{
				// بالكوين، لمقاومة الاحتلال الاقتصادي والاستقلال المالي للشعب الفلسطيني. من الشعب الى الشعب
				0xd8, 0xa8, 0xd8, 0xa7, 0xd9, 0x84, 0xd9, 0x83,
				0xd9, 0x88, 0xd9, 0x8a, 0xd9, 0x86, 0xd8, 0x8c,
				0x20, 0xd9, 0x84, 0xd9, 0x85, 0xd9, 0x82, 0xd8,
				0xa7, 0xd9, 0x88, 0xd9, 0x85, 0xd8, 0xa9, 0x20,
				0xd8, 0xa7, 0xd9, 0x84, 0xd8, 0xa7, 0xd8, 0xad,
				0xd8, 0xaa, 0xd9, 0x84, 0xd8, 0xa7, 0xd9, 0x84,
				0x20, 0xd8, 0xa7, 0xd9, 0x84, 0xd8, 0xa7, 0xd9,
				0x82, 0xd8, 0xaa, 0xd8, 0xb5, 0xd8, 0xa7, 0xd8,
				0xaf, 0xd9, 0x8a, 0x20, 0xd9, 0x88, 0xd8, 0xa7,
				0xd9, 0x84, 0xd8, 0xa7, 0xd8, 0xb3, 0xd8, 0xaa,
				0xd9, 0x82, 0xd9, 0x84, 0xd8, 0xa7, 0xd9, 0x84,
				0x20, 0xd8, 0xa7, 0xd9, 0x84, 0xd9, 0x85, 0xd8,
				0xa7, 0xd9, 0x84, 0xd9, 0x8a, 0x20, 0xd9, 0x84,
				0xd9, 0x84, 0xd8, 0xb4, 0xd8, 0xb9, 0xd8, 0xa8,
				0x20, 0xd8, 0xa7, 0xd9, 0x84, 0xd9, 0x81, 0xd9,
				0x84, 0xd8, 0xb3, 0xd8, 0xb7, 0xd9, 0x8a, 0xd9,
				0x86, 0xd9, 0x8a, 0x2e, 0x20, 0xd9, 0x85, 0xd9,
				0x86, 0x20, 0xd8, 0xa7, 0xd9, 0x84, 0xd8, 0xb4,
				0xd8, 0xb9, 0xd8, 0xa8, 0x20, 0xd8, 0xa7, 0xd9,
				0x84, 0xd9, 0x89, 0x20, 0xd8, 0xa7, 0xd9, 0x84,
				0xd8, 0xb4, 0xd8, 0xb9, 0xd8, 0xa8,
			},
			Sequence: 0xffffffff,
		},
	},
	TxOut: []*wire.TxOut{
		{
			Value: 0x3B9ACA00, // 10PALC
			PkScript: []byte{
				0x41, 0x04, 0x67, 0x8a, 0xfd, 0xb0, 0xfe, 0x55, /* |A.g....U| */
				0x48, 0x27, 0x19, 0x67, 0xf1, 0xa6, 0x71, 0x30, /* |H'.g..q0| */
				0xb7, 0x10, 0x5c, 0xd6, 0xa8, 0x28, 0xe0, 0x39, /* |..\..(.9| */
				0x09, 0xa6, 0x79, 0x62, 0xe0, 0xea, 0x1f, 0x61, /* |..yb...a| */
				0xde, 0xb6, 0x49, 0xf6, 0xbc, 0x3f, 0x4c, 0xef, /* |..I..?L.| */
				0x38, 0xc4, 0xf3, 0x55, 0x04, 0xe5, 0x1e, 0xc1, /* |8..U....| */
				0x12, 0xde, 0x5c, 0x38, 0x4d, 0xf7, 0xba, 0x0b, /* |..\8M...| */
				0x8d, 0x57, 0x8a, 0x4c, 0x70, 0x2b, 0x6b, 0xf1, /* |.W.Lp+k.| */
				0x1d, 0x5f, 0xac, /* |._.| */
			},
		},
	},
	LockTime: 0,
}

// 0000000042a08e73efeed391aa13677dfaf5f7d066ca2248b02482bf37b1cf8e
var genesisHash = chainhash.Hash([chainhash.HashSize]byte{
	142, 207, 177, 55, 191, 130, 36, 176, 72, 34, 202, 102, 208,
	247, 245, 250, 125, 103, 19, 170, 145, 211, 238, 239, 115,
	142, 160, 66, 0, 0, 0, 0,
})

//07615f15061cf0691618997604aa39d0263befe810d9503b425c5bc274148caa
var genesisMerkleRoot = chainhash.Hash([chainhash.HashSize]byte{
	170, 140, 20, 116, 194, 91, 92, 66, 59, 80, 217, 16, 232, 239,
	59, 38, 208, 57, 170, 4, 118, 153, 24, 22, 105, 240, 28, 6, 21,
	95, 97, 7,
})

// genesisBlock defines the genesis block of the block chain which serves as the
// public transaction ledger for the main network.
var genesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},
		MerkleRoot: genesisMerkleRoot,
		Timestamp:  time.Unix(0x61180882, 0),
		Bits:       0x1f000010,
		Nonce:      268363379,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// regTestGenesisHash is the hash of the first block in the block chain for the
// regression test network (genesis block).
var regTestGenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x06, 0x22, 0x6e, 0x46, 0x11, 0x1a, 0x0b, 0x59,
	0xca, 0xaf, 0x12, 0x60, 0x43, 0xeb, 0x5b, 0xbf,
	0x28, 0xc3, 0x4f, 0x3a, 0x5e, 0x33, 0x2a, 0x1f,
	0xc7, 0xb2, 0xb7, 0x3c, 0xf1, 0x88, 0x91, 0x0f,
})

// regTestGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the regression test network.  It is the same as the merkle root for
// the main network.
var regTestGenesisMerkleRoot = genesisMerkleRoot

// regTestGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the regression test network.
var regTestGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: regTestGenesisMerkleRoot, // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1296688602, 0), // 2011-02-02 23:16:42 +0000 UTC
		Bits:       0x1d00ffff,               // 545259519 [7fffff0000000000000000000000000000000000000000000000000000000000]
		Nonce:      2,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// testNet3GenesisHash is the hash of the first block in the block chain for the
// test network (version 3).
var testNet3GenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0x43, 0x49, 0x7f, 0xd7, 0xf8, 0x26, 0x95, 0x71,
	0x08, 0xf4, 0xa3, 0x0f, 0xd9, 0xce, 0xc3, 0xae,
	0xba, 0x79, 0x97, 0x20, 0x84, 0xe9, 0x0e, 0xad,
	0x01, 0xea, 0x33, 0x09, 0x00, 0x00, 0x00, 0x00,
})

// testNet3GenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the test network (version 3).  It is the same as the merkle root
// for the main network.
var testNet3GenesisMerkleRoot = genesisMerkleRoot

// testNet3GenesisBlock defines the genesis block of the block chain which
// serves as the public transaction ledger for the test network (version 3).
var testNet3GenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},          // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: testNet3GenesisMerkleRoot, // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1296688602, 0),  // 2011-02-02 23:16:42 +0000 UTC
		Bits:       0x1d00ffff,                // 486604799 [00000000ffff0000000000000000000000000000000000000000000000000000]
		Nonce:      0x18aea41a,                // 414098458
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// simNetGenesisHash is the hash of the first block in the block chain for the
// simulation test network.
var simNetGenesisHash = chainhash.Hash([chainhash.HashSize]byte{ // Make go vet happy.
	0xf6, 0x7a, 0xd7, 0x69, 0x5d, 0x9b, 0x66, 0x2a,
	0x72, 0xff, 0x3d, 0x8e, 0xdb, 0xbb, 0x2d, 0xe0,
	0xbf, 0xa6, 0x7b, 0x13, 0x97, 0x4b, 0xb9, 0x91,
	0x0d, 0x11, 0x6d, 0x5c, 0xbd, 0x86, 0x3e, 0x68,
})

// simNetGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the simulation test network.  It is the same as the merkle root for
// the main network.
var simNetGenesisMerkleRoot = genesisMerkleRoot

// simNetGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the simulation test network.
var simNetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: simNetGenesisMerkleRoot,  // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1401292357, 0), // 2014-05-28 15:52:37 +0000 UTC
		Bits:       0x207fffff,               // 545259519 [7fffff0000000000000000000000000000000000000000000000000000000000]
		Nonce:      2,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// sigNetGenesisHash is the hash of the first block in the block chain for the
// signet test network.
var sigNetGenesisHash = chainhash.Hash{
	0xf6, 0x1e, 0xee, 0x3b, 0x63, 0xa3, 0x80, 0xa4,
	0x77, 0xa0, 0x63, 0xaf, 0x32, 0xb2, 0xbb, 0xc9,
	0x7c, 0x9f, 0xf9, 0xf0, 0x1f, 0x2c, 0x42, 0x25,
	0xe9, 0x73, 0x98, 0x81, 0x08, 0x00, 0x00, 0x00,
}

// sigNetGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the signet test network. It is the same as the merkle root for
// the main network.
var sigNetGenesisMerkleRoot = genesisMerkleRoot

// sigNetGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the signet test network.
var sigNetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  chainhash.Hash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: sigNetGenesisMerkleRoot,  // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1598918400, 0), // 2020-09-01 00:00:00 +0000 UTC
		Bits:       0x1e0377ae,               // 503543726 [00000377ae000000000000000000000000000000000000000000000000000000]
		Nonce:      52613770,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}
