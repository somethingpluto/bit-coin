package cli

import (
	"bitcoin/block"
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
	blockChain *block.BlockChain
}

const USAGE = `
	Usage:
 		addblock -data BLOCK_data  add a block to blockchain(向区块链中添加一个区块 简称挖矿 eg: addblock -data 11)
		printChain 				   print all block of the chain(打印区块链中所有区块)
`

func NewCli(bc *block.BlockChain) *CLI {
	return &CLI{blockChain: bc}
}

func (cli *CLI) printUsage() {
	fmt.Println(USAGE)
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	addBlockData := addBlockCmd.String("data", "", "Block data")
	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printChain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}
	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.blockChain.AddBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}
