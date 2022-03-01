package main

//go run main.go add -block "YOUR BLOCK DATA HERE"

//go run main.go print


import (
    "flag"
    "fmt"
    "os"
    "runtime"
    "strconv"

	"github.com/s4pv/learnBC/blockchain"

)

type CommandLine struct {
//    blockchain *blockchain.BlockChain //file package.struct
}

func (cli *CommandLine) printusage() {
    fmt.Println("Usage: ")
    fmt.Println(" add -block <BLOCK_DATA> - add a block to the chain")
    fmt.Println(" print - prints the blocks in the chain")
}

func (cli *CommandLine) validateArgs() {
    if len(os.Args) < 2 {
        cli.printusage()
        runtime.Goexit()

    }
}

//func (cli*CommandLine) addBlock(data string) {
//    cli.blockchain.AddBlock(data)
//    fmt.Println("Added Block!")
//}

func (cli \*CommandLine) printUsage() {

 fmt.Println("Usage: ")
 fmt.Println("getbalance -address ADDRESS - get balance for ADDRESS")
 fmt.Println("createblockchain -address ADDRESS creates a blockchain and rewards the mining fee")
 fmt.Println("printchain - Prints the blocks in the chain")
 fmt.Println("send -from FROM -to TO -amount AMOUNT - Send amount of coins from one address to another")

}

func (cli \*CommandLine) printChain() {
    iterator := cli.blockchain.Iterator()

    for {
        block := iterator.Next()
        fmt.Printf("Previous hash: %x\n", block.PrevHash)
        fmt.Printf("hash: %x\n", block.Hash)
        pow := blockchain.NewProofOfWork(block)
        fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
        fmt.Println()

        if len(block.PrevHash) == 0 {
            break
        }
    }
}

func (cli *CommandLine) createBlockChain(address string) {
    newChain := blockchain.InitBlockChain(address)
    newChain.Database.Close()
    fmt.Println("Finished creating chain")
}

func (cli *CommandLine) getBalance(address string) {
    chain := blockchain.ContinueBlockChain(address)
    defer chain.Database.Close()

    balance := 0
    UTXOs := chain.FindUTXO(address)

    for _, out := range UTXOs {
        balance += out.Value
    }

    fmt.Printf("Balance of %s: %d\n", address, balance)
}

func (cli *CommandLine) send(from, to string, amount int) {
    chain := blockchain.ContinueBlockChain(from)
    defer chain.Database.Close()

    tx := blockchain.NewTransaction(from, to, amount, chain)

    chain.AddBlock([]*blockchain.Transaction{tx})
    fmt.Println("Success!")
}


func (cli *CommandLine) run() {
    cli.validateArgs()

    getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
    createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
    sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
    printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

    getBalanceAddress := getBalanceCmd.String("address", "", "The address to get balance for")
    createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to send genesis block reward to")
    sendFrom := sendCmd.String("from", "", "Source wallet address")
    sendTo := sendCmd.String("to", "", "Destination wallet address")
    sendAmount := sendCmd.Int("amount", 0, "Amount to send")

    switch os.Args[1] {
    case "getbalance":
        err := getBalanceCmd.Parse(os.Args[2:])
        if err != nil {
            log.Panic(err)
        }
    case "createblockchain":
        err := createBlockchainCmd.Parse(os.Args[2:])
        if err != nil {
            log.Panic(err)
        }
    case "printchain":
        err := printChainCmd.Parse(os.Args[2:])
        if err != nil {
            log.Panic(err)
        }
    case "send":
        err := sendCmd.Parse(os.Args[2:])
        if err != nil {
            log.Panic(err)
        }
    default:
        cli.printUsage()
        runtime.Goexit()
    }

    if getBalanceCmd.Parsed() {
        if *getBalanceAddress == "" {
            getBalanceCmd.Usage()
            runtime.Goexit()
        }
        cli.getBalance(*getBalanceAddress)
    }

    if createBlockchainCmd.Parsed() {
        if *createBlockchainAddress == "" {
            createBlockchainCmd.Usage()
            runtime.Goexit()
        }
        cli.createBlockChain(*createBlockchainAddress)
    }

    if printChainCmd.Parsed() {
        cli.printChain()
    }

    if sendCmd.Parsed() {
        if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
            sendCmd.Usage()
            runtime.Goexit()
        }

        cli.send(*sendFrom, *sendTo, *sendAmount)
    }
}

func main() {
    defer os.Exit(0)

    chain := blockchain.InitBlockChain()
    cli := CommandLine{chain}
    cli.run()

}