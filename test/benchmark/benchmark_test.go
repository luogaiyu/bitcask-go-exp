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
	"time"

	// "encoding/binary"
	"fmt"
	"errors"
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

	// 测试 put 和 get方法
	cur_db.Put(key, value)
	b, _ := cur_db.Get(key)
	fmt.Println(string(b))
	cur_db.TrucDB()
}

// 单线程存数据1000次
// 20240824 发现 如果多次进行数据存放操作, 会导致数据被重复写入同一个 offset 这个是因为修改了 offset的逻辑导致的, 看起来offset 是有必要的
// 20240825: 序列化 反序列化的逻辑有问题 20240825: 成功修复
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
			assert.Equal(res_map[key], string(res), "test db [delete get] process is wrong count : "+strconv.Itoa(i))
		}
	}
	cur_db.TrucDB()
}

// 测试db 删除功能
func Test_benchmark_single_process_delete_1_times(T *testing.T) {

}
// goroutin 是
// 多线程 如果数据不及时消费 会阻塞 子线程或者主线程
func Test_process(T *testing.T){
	cur_db := db.InitDB()
	ch := make(chan error, 10)
	// ch := make(chan int , 4)
	go process(T,cur_db,1,1,1,ch)
	
	// 手动读取并检查通道状态
	for {
		value, ok := <-ch
		if !ok {
			// 通道已关闭，退出循环
			fmt.Println("Channel closed and all data received")
			break
		}
		fmt.Println(value.Error()) // 打印从通道中接收到的值
	}
}

func process(T *testing.T, db *db.DB, key uint32, value uint32, count uint32, ch chan error) {
	assert := assert.New(T)
	ch <- errors.New("start")
	for i := uint32(0); i < 100; i++ {
		db.Put(utils.UInt32_2_Bytes(key), utils.UInt32_2_Bytes(uint32(value)))
		time.Sleep(30*time.Millisecond)
		res,_ := db.Get(utils.UInt32_2_Bytes(key))

		err_bool := assert.Equal(res,utils.UInt32_2_Bytes(value) , "wrong!!! ")
		
		if(err_bool){
			ch <- errors.New("wrong!!!" + strconv.Itoa(int(count)) + "")
		}
	}
	close(ch)
}

// 使用两个线程 同时读写
// 可以实现 多线程的测试
func Test_benchmark_double_process_put_get_100_times(T *testing.T) {
	cur_db := db.InitDB()
	ch := make(chan error, 10)
	// 开两个线程 每个读写100次
	go process(T, cur_db, 1,1,1,ch)
	go process(T, cur_db, 1,2,2,ch)
	go process(T, cur_db, 1,3,3,ch)
	go process(T, cur_db, 1,4,4,ch)
	
	// 手动读取并检查通道状态
	for {
		value, ok := <-ch
		if !ok {
			// 通道已关闭，退出循环
			fmt.Println("Channel closed and all data received")
			break
		}
		fmt.Println(value.Error()) // 打印从通道中接收到的值
	}
}

// 模拟 一万个用户同时读和同时写100次
func Test_benchmark_10000_process_put_get_100_times(T *testing.T) {
	// cur_db := db.InitDB()
	// // 模拟一万个用户进行读取和写入
	// for i := 0; i < 100000; i++ {
	// 	// go process(T, cur_db, uint32(i))
	// }
}
