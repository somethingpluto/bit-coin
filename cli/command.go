package cli

import (
	"bitcoin/block"
	"encoding/hex"
	"fmt"
	"github.com/fatih/color"
	"strconv"
	"time"
)

func (cli *CLI) printChain() {
	bci := cli.blockChain.Iterator()
	for {
		b := bci.Next()
		fmt.Printf("|=================================================================================|\n")
		fmt.Printf("|%-90s|\n", color.YellowString("[PreBlockHash]: %s", hex.EncodeToString(b.PrevBlockHash)))
		fmt.Printf("|%-90s|\n", color.RedString("[Data]: %x", b.Data))
		fmt.Printf("|%-90s|\n", color.BlueString("[Timestamp]: %d %v", b.Timestamp, time.Unix(b.Timestamp, 0)))
		fmt.Printf("|%-90s|\n", color.BlackString("[Hash]: %x", b.Hash))
		pow := block.NewProofOfWork(b)
		fmt.Printf("|%-90s|\n", color.GreenString("[Block Pow]: %s", strconv.FormatBool(pow.Validate())))
		fmt.Println("|=================================================================================|")

		fmt.Println()

		if len(b.PrevBlockHash) == 0 {
			break
		}
	}
}

func (cli *CLI) addBlock(args []string) {
	if len(args) != 3 || args[1] != "data" || args[2] == "" {
		fmt.Println("addblock command invalid")
		cli.printUsage()
	} else {
		cli.blockChain.AddBlock(args[2])
	}
}

func (cli *CLI) help() {
	fmt.Printf(color.RedString("%s", USAGE))
}
