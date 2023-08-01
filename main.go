package main

import (
	"bitcoin/block"
	"bitcoin/cli"
	"bitcoin/constant"
	"flag"
)

func main() {
	flag.IntVar(&constant.TARGET_BITES, "n", 16, "计算难度")
	flag.Parse()
	blockChain := block.NewBlockChain()
	defer blockChain.CloseDB()
	CLI := cli.NewCli(blockChain)

	CLI.Run()
}
