package benchmark

/**
首先模拟线上的redis操作
kv数据库 主要包含以下三类操作
存数据
取数据
删数据
benchmark 主要分为以下几个部分
1. 单次读取 单线程反复 存读删 100次 判断会不会出现问题
2. 单线程 多次存读写 每次存10条数据 读10条数据 并删除其中的5条, 重复3次
3. 多线程 一个线程存数据, 一个线程读数据
4. 多线程 多个线程同时存数据,并同时读取数据,并修改对应的数据, 保证数据的一致性和可靠性
5. 另外还需要模拟热点读取, 当多个用户同时读写一个数据时, 保证数据不出错
*/

import (
	"bitcask-go-exp/db"
	"bitcask-go-exp/db/utils"
	"encoding/binary"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 单次读取 单线程反复 存读删 100次 判断会不会出现问题
func Test_benchmark_single_process_put_get_once(T *testing.T) {
	// 首先创建对应的DB对象
	cur_db := db.InitDB()
	key := []byte("1234")
	value := []byte("1234")

	cur_db.Put(key, value)

	// 测试 put 和 get方法
	b, _ := cur_db.Get(key)
	fmt.Println(b)
	cur_db.TrucDB()
}

// 单线程存数据100次
func Test_benchmark_single_process_put_get_100_times(T *testing.T) {
	cur_db := db.InitDB()
	res_map := make(map[string]string)

	for i := 0; i <= 1000; i++ {
		key := utils.GetRandomString()
		value := utils.GetRandomString()
		res_map[key] = value
		cur_db.Put([]byte(key), []byte(value))
	}

	i := 0

	// 打印当前的数据
	for key := range res_map {
		i++
		res, err := cur_db.Get([]byte(key))
		if err != nil {
			fmt.Println(err.Error())
		} else {
			assert := assert.New(T)
			assert.Equal(res_map[key], res, "test db [delete get] process is wrong count : "+strconv.Itoa(i))
		}

	}
}

func process(T *testing.T, db *db.DB, quant uint32) {
	for i := uint32(0); i < 1000; i++ {
		db.Put(utils.UInt32_2_Bytes(i), utils.UInt32_2_Bytes(uint32(i*quant)))
	}
	//
	res := uint32(0)
	assert := assert.New(T)
	for i := uint32(0); i < 1000; i++ {
		tmp := utils.UInt32_2_Bytes(i)
		byt, _ := db.Get(tmp)
		res = binary.BigEndian.Uint32(byt)
		// res = utils.Bytes_2_UInt32()
		assert.Equal(res, uint32(i*quant), "test db [delete get] process is wrong count : "+strconv.Itoa(int(i)))
	}
}

// 使用两个线程 同时读写
func Test_benchmark_double_process_put_get_100_times(T *testing.T) {
	cur_db := db.InitDB()
	// 开两个线程 每个读写100次
	go process(T, cur_db, uint32(100))

	go process(T, cur_db, uint32(1000))
}

// 模拟 一万个用户同时读和同时写100次
func Test_benchmark_10000_process_put_get_100_times(T *testing.T) {
	cur_db := db.InitDB()
	// 模拟一万个用户进行读取和写入
	for i := 0; i < 100000; i++ {
		go process(T, cur_db, uint32(i))
	}
}
