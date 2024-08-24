package utils


func InitTmpOptionForTest() *Option{
	return &Option{
		DirPath:"/Users/bowen.yin/code/kv-projects/bitcask-go-exp/db/data/",
	}
}
// 保存一些对于db常用的常量
type Option struct{
	DirPath string // 数据路径
}

