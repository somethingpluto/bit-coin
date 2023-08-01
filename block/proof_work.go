package block

import (
	"bitcoin/constant"
	"bitcoin/utils"
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/fatih/color"
	"math/big"
	"time"
)

const TIME_FORMAT = "2006-01-02 15:04:05"

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork
// @Description: 新建工作证明
// @param b
// @return *ProofOfWork
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-constant.TARGET_BITES))
	pow := &ProofOfWork{b, target}
	return pow
}

// prepareData
// @Description: 结合nonce生成区块数据
// @receiver pow
// @param nonce
// @return []byte
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{pow.block.PrevBlockHash, pow.block.Data, utils.IntToHex(pow.block.Timestamp), utils.IntToHex(int64(constant.TARGET_BITES)), utils.IntToHex(int64(nonce))}, []byte{})
	return data
}

// Run
// @Description: 工作计算hash
// @receiver pow
// @return int
// @return []byte
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	start := time.Now()
	startFormat := start.Format(TIME_FORMAT)
	fmt.Printf(color.YellowString("Mining the block %s\n", pow.block.Data))
	fmt.Printf("start time: %v\n", startFormat)
	for nonce < constant.MAX_NONCE {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
		fmt.Printf(color.GreenString("nonce: %d  hash: %x\r", nonce, hash))
		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf(color.RedString("%x\n", hash))
			duration := time.Since(start)
			fmt.Printf(color.RedString("spend time: %v\n", duration))
			break
		} else {
			nonce++
		}
	}
	fmt.Printf("\n\n")
	return nonce, hash[:]
}

// Validate
// @Description: 检验区块hash
// @receiver pow
// @return bool
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}
