package cli

import (
	"bitcoin/block"
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

const shell = ">>"

type CLI struct {
	blockChain *block.BlockChain
}

const USAGE = `
Usage:
	addblock data BLOCK_data  add a block to blockchain(向区块链中添加一个区块 简称挖矿 eg: addblock -data 11)
 	printChain 				  print all block of the chain(打印区块链中所有区块 eg: 直接打印出)
`

// NewCli
// @Description: 根据区块链创建cli
// @param bc
// @return *CLI
func NewCli(bc *block.BlockChain) *CLI {
	return &CLI{blockChain: bc}
}

// printUsage
// @Description: 打印cli使用方法
// @receiver cli
func (cli *CLI) printUsage() {
	fmt.Println(USAGE)
}

// Run
// @Description: 脚手架运行
// @receiver cli
func (cli *CLI) Run() {
	cli.help()
	option := ""
	var args []string
	for true {
		fmt.Printf(color.GreenString("%s", shell))
		reader := bufio.NewReader(os.Stdin)
		option, _ = reader.ReadString('\r')
		option = strings.Replace(option, "\r", "", -1)
		args = strings.Split(option, " ")
		switch args[0] {
		case "printchain":
			cli.printChain()
			break
		case "addblock":
			cli.addBlock(args)
			break
		case "help":
			cli.help()
		case "quit":
			return
		default:
			fmt.Println("invalid input")
			cli.help()
		}
	}
}
