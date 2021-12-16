package main

import (
    "flag"
    "fmt"
    "os"
    "runtime"
    "strconv"

	"github.com/s4pv/learnBC/blockchain"

)

type CommandLine struct {
    blockchain *blockchain.BlockChain
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

func (cli *CommandLine) printChain() {
    iterator := cli.blockchain.Iterator()

    for {
        block := iterator.Next()
        fmt.printf("Previous hash: %x\n", block.PrevHash)
        fmt.printf("data: %s\n", block.Data)
        fmt.printf("hash: %x\n", block.Hash)
        pow := blockchain.NewProofOfWork(block)
        fmt.printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
        fmt.println()

        if len(block.PrevHash) == 0 {
            break
        }
    }
}


func (cli *CommandLine) run() {
    cli.validateArgs()

    addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
    printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
    addBlockData := addBlockCmd.String("block", "", "Block data")

    switch os.Args[1] {
    case "add":
        err := addBlockCmd.Parse(os.Args[2:])
        blockchain.Handle(err)

    case "print":
        err := printChainCmd.Parse(os.Args[2:])
        blockchain.Handle(err)

    default:
        cli.printusage()
        runtime.Goexit()

    }

    if addBlockCmd.Parsed() {
        if *addBlockData == "" {
            addBlockCmd.Usage()
            runtime.Goexit()
        }
        cli.addBlock(*addBlockData)
    }
    if printChainCmd.Parsed() {
        cli.printChain()
    }

}

func main() {
    defer os.Exit(0)

    chain := blockchain.InitBlockChain()
    defer chain.Database.Close()

    cli := CommandLine{chain}

    cli.run()

}