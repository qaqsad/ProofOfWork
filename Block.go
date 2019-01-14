package BLC

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp int64
	PervBlockHash []byte
	Data []byte
	Hash []byte
	Nonce int
}

func (block *Block)setHash(){
//	将时间戳转化多维数组
timestamp := strconv.FormatInt(block.Timestamp,10)
//fmt.Println("timestamp:",timestamp)

timeBytes := []byte(timestamp)
//fmt.Println("timeBytes:",timeBytes)

//  拼接所有属性
blockBytes := bytes.Join([][]byte{timeBytes,block.PervBlockHash,block.Data},[]byte{})
//fmt.Println("blockBytes:",blockBytes)
// 	生成hash
hash := sha256.Sum256(blockBytes)
block.Hash = hash[:]

}
//	工厂方法创建区块
func NewBlock (data string,pervBlockHash []byte) *Block {
	//	nonce值
	block := &Block{time.Now().Unix(),pervBlockHash,[]byte(data),[]byte{},0}
	//	将block传进工作量证明方法中，创建pow对象，从而通过计算得出hash值
	pow := newProofOfWork(block)
	//	执行一次工作量证明

	nonce,hash := pow.Run()
	//
	//// 设置hash值
	block.Hash = hash[:]
	//
	////	设置nonce值
	block.Nonce = nonce
	return block
	}

//创建创世区块，返回创世区块
func NewGenesisBlock() *Block  {
return NewBlock("Genesis Block",[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
}