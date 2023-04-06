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

		fmt.Printf("Prev hash: %x\n", b.PrevBlockHash)
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("Hash: %x\n", b.Hash)
		fmt.Printf("|=================================================================================|\n")
		fmt.Printf("|%-90s|\n", color.YellowString("[PreBlockHash]: %s", hex.EncodeToString(b.PrevBlockHash)))
		fmt.Printf("|%-90s|\n", color.RedString("[Data]: %x", b.Data))
		fmt.Printf("|%-90s|\n", color.BlueString("[Timestamp]: %d %v", b.Timestamp, time.Unix(b.Timestamp, 0)))
		fmt.Printf("|%-90s|\n", color.BlackString("[Hash]: %x", b.Hash))
		fmt.Println("|=================================================================================|")
		fmt.Printf("\n\n")
		pow := block.NewProofOfWork(b)
		fmt.Printf("|%-90s|\n", color.GreenString("[Block Pow]: %s", strconv.FormatBool(pow.Validate())))
		fmt.Println()

		if len(b.PrevBlockHash) == 0 {
			break
		}
	}
}
