package db

import (
	// "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

// loadIndexFromDataFile 前置方法,用于创造测试的方法
func Test_db_loadIndexFromDataFile_pre(t *testing.T) {
	// 单元测试
	assert := assert.New(t)
	cur_db := InitDB()
	// 首先创建对应的DB对象
	key := []byte("1234")
	value := []byte("1234")
	// 测试 put 和 get方法
	cur_db.Put(key, value)
	// cur_db.Put([]byte("1234"), []byte("1234"))
	res, _ := cur_db.Get(key)
	assert.Equal(res, value, "put get function have something wrong!")
}

// 新功能 添加新功能 从数据文件中重新构建索引
func Test_db_loadIndexFromDataFile(t *testing.T) {
	assert := assert.New(t)
	// Test_db_loadIndexFromDataFile_pre(t)
	cur_db := InitDB()
	cur_db.loadIndexFromDataFile()
	key := []byte("12345")
	value := []byte("12345")
	res, _ := cur_db.Get(key)
	assert.Equal(value, res, "put get function have something wrong!")
}

func db_syntax_test() {

}
