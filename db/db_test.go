package db

import (
	"fmt"
	"testing"
	// "github.com/google/btree"
)

// 用来测试 DB的常规操作

// 测试KV数据库的 新增获取删除 api, 但是现在有个问题, 怎么保存索引, 也就是中间状态?
// 目前的实现方法 是通过磁盘中文件重新对 数据创建中间状态
func db_put_test() {
	
}

func db_get_test() {

}

func db_delete_test() {

}

func Test_db_loadIndexFromDataFile(t *testing.T) {
	cur_db := InitDB()
	
	res,err := cur_db.Get([]byte("12345"))


	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println(string(res))
	}

	// item:= db.index.Get([]byte("1234"))
	// if item != nil {
	// 	item = item.(*Item)
	// }else{
	// 	log.Println("index 返回为空")
	// }
	

	// lgRecord := db.fio.Read(item.lgPos)
	// fmt.Println("||||")
	
	
	// fmt.Println(lgRecord.RType)
	// fmt.Println(lgRecord.KeySize)
	// fmt.Println(lgRecord.ValueSize)
	// fmt.Println(string(lgRecord.Key))
	// fmt.Println("---------")
	// fmt.Println(string(lgRecord.Value))

	// fmt.Println()
	// fmt.Println(db.index.Get([]byte("4321")))
}

func db_syntax_test(){

}
