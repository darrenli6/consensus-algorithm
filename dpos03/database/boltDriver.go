package database


// 将交易信息 候选人记录 区块 存储到数据库中

const dbFile = "dpos_blockchain_%s.db"

//区块的bucket
const BlockBucket= "dpos_blocks"

//受托人的bucket
const DelegateBucket =" dpos_delegate"

//记录交易记录的bucket

const TransfersBucket="dpos_transfer"

const LastHash ="lastHash"

//初始化本地数据库
