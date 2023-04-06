package main

import (
	"bitcoin/block"
	"bitcoin/cli"
)

func main() {
	blockChain := block.NewBlockChain()
	defer blockChain.CloseDB()
	CLI := cli.NewCli(blockChain)

	CLI.Run()
}
