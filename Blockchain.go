package BLC
//新建区块链结构体
type Blockchain struct {
	//创建有序区块
	Blocks []*Block
}
//新增区块
func (blockchain *Blockchain) AddBlock (data string)  {
//	创建新的block
prevBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
newBlock := NewBlock(data,prevBlock.Hash)


//	将区块链添加到Blocks
blockchain.Blocks = append(blockchain.Blocks,newBlock)
}
//创建带有创世区块节点的区块链
func NewBolckchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}