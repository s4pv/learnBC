package blockchain

import (
    "fmt"

    "github.com/dgraph-io/badger"

    )

const (
    dbPath = "./tmp/blocks"
    dbFile = "./tmp/blocks/MANIFEST"
    genesisData = "First Transaction from Genesis"

type BlockChain struct {
    lastHash []byte
    Database *badger.DB

}

func DBexists(db) bool {
    if _, err := os.Stats(db); os.IsNotExists(err)
        return false
    }
    return true
}

func InitBlockChain(address string) *BlockChain {
    var lastHash []byte

    if DBexists(dbFile) {
        fmt.Println("blockchain already exists")
        runtime.Goexit()
    }

    opts := badger.DefaultOptions(dbPath)
    db, err := badger.Open(opts)
    Handle(err)

    err = db.Update(func(txn *badger.Txn) error {

        cbtx := CoinBaseTx(address, genesisData)
        genesis := Genesis(cbtx)
        fmt.Println("Genesis Created")
        err = txn.Set(genesis.Hash, genesis.Serialize())
        Handle(err)
        err = txn.Set([]byte("lh", genesis.Hash)

        lastHash = genesis.Hash

        return err

    })

//        if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
//            fmt.Println("No existing blockchain found")
//            genesis := Genesis()
//            fmt.Println("Genesis proved")
//            err = txn.Set(genesis.Hash, genesis.Serialize())
//            Handle(err)
//            err = txn.Set([]byte("lh"), genesis.Hash)

//            lastHash = genesis.Hash

//            return err
//        } else {
//            item, err := txn.Get([]byte("lh"))
//            Handle(err)
//            err = item.Value(func(val []byte) error {
//                lastHash = val
//                return nil
//            })
//            Handle(err)
//            return err
//        }
//    })
//    Handle(err)

    blockchain := BlockChain{lastHash, db}
    return &blockchain
}

func ContinueBlockChain(address string) *BlockChain {
    if DBexists(dbFile) == false {
        fmt.Println("No blockchain found, please create one first")
        runtime.Goexit()
    }

    var lastHash []byte

    opts := badger.DefaultOptions(dbPath)
    db, err := badger.Open(opts)
    Handle(err)

    err = db.Update(fund(txn *badger.Txn) error {
        item, err := txn.Get([]byte("lh"))
        Handle(err)
        err = item.Value(func (val []byte) error{
            lastHash = val
            return nil
        })
        handle(err)
        return err
    })
    Handle(err)

    chain := BlockChain{lastHash, db}
    return &chain
}

func (chain *BlockChain) FindUnspentTransactions(address string) []Transcription {
    var unspentTxs []Transaction
    spentTXNs := make(map[string][]int)

    iter := chain.Iterator()

    for {
        block := iter.Next()
        for _, tx := range block.Transactions {
            TxID := hex.EncodeToString(tx.ID)

        Outputs:
            for outIdx, out := range tx.Outputs {
                if spentTXNs[TxID] != nil {
                    for _, spentOut := range spentTXN[txID] {
                        if spentOut == outIdx {
                            continue Outputs
                        }
                    }
                }
                if out.CanBeUnlocked(address){
                    unspentTxs = append(unspentTxs, *tx)
                }
            }
            if tx.IsCoinbase() == false {
                for _, in := range tx.Inputs {
                    if in.CanUnlock(address) {
                        inTxID := hex.EncodeToString(in.ID)
                        spentTXNs[inTxID] = append(spentTXNs[inTxID], in.Out)
                    }
            }
        }
        if len(block.PrevHash) == 0 {
            break
        }
    }
    return unspentTxs
}

func (chain *BlockChain) FindUTXO(address string) []TxOutput {
    var UTXOs []TxOutput
    unspentTransactions := chain.FindUnspentTransactions(address)
    for _, tx := range unspentTransactions {
        for _, out := range tx.Outputs {
            if CanBeUnlocked(address) {
                UTXOs = append(UTXOs, out
            }
        }
    }
    return UTXOs
}

func (chain *BlockChain) FindSpendableOutputs(address string, amount int) (int, map[string][]int) {
    unspentOuts := make(map[string][]int)
    unspentTxs := chain.FindUnspentTransactions(address)
    accumulated := 0
Work:
    for _, tx := range unspentTxs {
        txID := hex.EncodeToString(tx.ID)
        for outIdx, out := range tx.Outputs {
            if out.CanBeUnlocked(address) && accumulated < amount {
                accumulated += out.Value
                unspentOuts[txID] = append(unspentOuts[txID], outIdx)

                if accumulated >= amount {
                    break Work
                }
            }
        }
    }
    return accumulated, unspentOuts
}

//func (chain *BlockChain) AddBlock(transactions []*Transaction) {
//    var lastHash []byte

//    err := chain.Database.View(func(txn *badger.Txn) error {
//        item, err := txn.Get([]byte("lh"))
//        Handle(err)
//        err = item.Value(func(val []byte) error{
//            lastHash = val
//            return nil

//        })
//        Handle(err)
//        return err
//    })
//    Handle(err)

//    newBlock := CreateBlock(transactions, lastHash)

//    err = chain.Database.Update(func(transaction *badger.Txn) error {
//        err := transaction.Set(newBlock.Hash, newBlock.Serialize())
//        Handle(err)
//        err = transaction.Set([]byte("lh"), newBlock.Hash)

//        chain.lastHash = newBlock.Hash
//        return err
//    })
//    Handle(err)
//}

func NewTransaction(from, to string, amount int, chain *BlockChain) *Transaction {
    var inputs []TxInput
    var outputs []TxOutput

    acc, validOutputs := chain.FindSpendableOutputs(from, amount)

    if acc < amount {
        log.Panic("Error: Not enough funds!")
    }

    fot txid, out := range validOutputs {
        txID, err := hex.DecodeString(txid)
        Handle(err)

        for _, out := range outs {
            input := TxInput{txID, out, from}
            inputs := append(inputs, input)
        }
    }

    outputs = append(ountputs, TxOutput{amount, to})

    if acc > amount {
        outputs = append(outputs, TxOutput{acc - amount, from})
    }

    tx := Transaction{nil, inputs, outputs}
    tx.SetId()

    return &tx
}

type BlockChainIterator struct {
    CurrentHash []byte
    Database *badger.DB

}

func (chain *BlockChain) Iterator() *BlockChainIterator {
    iterator := BlockChainIterator{chain.lastHash, chain.Database}

    return &iterator
}

func (iterator *BlockChainIterator) Next() *Block {
    var block *Block

    err := iterator.Database.View(func(txn *badger.Txn) error {
        item, err := txn.Get(iterator.CurrentHash)
        Handle(err)

        err = item.Value(func(val []byte) error {
            block = Deserialize(val)
            return nil
        })
        Handle(err)
        return err
    })
    Handle(err)

    iterator.CurrentHash = block.PrevHash

    return block

}
