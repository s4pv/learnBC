package transactions

const reward = 100

type Transaction struct {
    ID []byte
    Inputs []TxInput
    Outputs []TxOutput
}

type TxOutput struct {
    Value int
    PubKey string
}

type TxInput struct {
    ID []byte
    Out int
    Sig string
}

func CoinBaseTx(toAddress, data string) *Transaction {
    if data == "" {
        data = fmt.Sprintf("Coins to %s", toAddress)
    }
    txIn := TxInput{[]byte{}, -1, data}

    txOut := TxOutput{reward, toAddress}

    tx := Transaction{nil, []TxInput{txIn}, []TxOutput{txOut}}

    return &tx

}

func (tx *Transaction) SetID() {
    var encoded.bytes.Buffer
    var Hash [32]byte

    encoder := gob.NewEncoder(&encoded)
    err := encoder.Encode(tx)
    Handle(err)

    hash = sha256.Sum256(encoded.Bytes())
    tx.ID = hash[:]

}

func (in *TxInput) CanUnlock(data string) bool {
    return in.Sig == data
}

func (out *TxOutput) CanBeUnlocked(data string) bool {
    return out.PubKey == data
}

func (tx *Transaction) isCoinBase() bool {
    return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Out == -1
}