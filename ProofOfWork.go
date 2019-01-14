package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
	"math"
)
var (
	maxNonce =	math.MaxInt64
	)

const targetBites  = 20
//	创教pow结构体
type ProofOfWork struct {
//	当前需要验证的区块
	block *Block
//	大数存储
	target *big.Int
}
// 拼接数据 返回字节组
func (pow *ProofOfWork)prepareData(nonce int) []byte {
	data :=	bytes.Join([][]byte{
		pow.block.PervBlockHash,
		pow.block.Data,
		IntToHex(pow.block.Timestamp),
		IntToHex(int64(targetBites)),
		IntToHex(int64(nonce)),
	},[]byte{})
	return data
}
func (pow *ProofOfWork)	Run() (int ,[]byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	//	打印数据信息
	fmt.Printf("Mining the block containing\"%s\"\n",pow.block.Data)
	for nonce < maxNonce {
	data := pow.prepareData(nonce)
	hash = sha256.Sum256(data)
	fmt.Printf("\r%x",hash)
	hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target)== -1 {
			break
		}else{
			nonce++
		}
	}
	return nonce,hash[:]
}

// 创建newProofOfWork的方法
func newProofOfWork(block *Block) *ProofOfWork {
//	输出target
	target := big.NewInt(1)
	target.Lsh(target,uint(256-targetBites))
	fmt.Println(target)
	fmt.Println("-----------")
	pow := &ProofOfWork{block,target}

	return pow
}

