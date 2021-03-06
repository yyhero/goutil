package redisUtil

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var (
	redisPoolObj *RedisPool
)

func init() {
	redisPoolObj = NewRedisPool("testPool", "10.1.0.21:6379", "redis_pwd", 3, 500, 200, 10*time.Second, 5*time.Second)
}

func TestTestConnection(t *testing.T) {
	succeedCount := 0
	expectedCount := 5
	for i := 0; i < expectedCount; i++ {
		if err := redisPoolObj.TestConnection(); err != nil {
			fmt.Printf("%s:%s\n", time.Now(), err)
		} else {
			succeedCount += 1
			fmt.Printf("%s:%s\n", time.Now(), "ok")
		}
		time.Sleep(time.Second * 3)
	}

	if succeedCount != expectedCount {
		t.Errorf("ExecptedCount:%d, but got %d", expectedCount, succeedCount)
	}
}

func TestGetName(t *testing.T) {
	expectedName := "testPool"
	if actualName := redisPoolObj.GetName(); actualName != expectedName {
		t.Errorf("GetName should be %s, but got %s", expectedName, actualName)
	}
}

func TestGetAddress(t *testing.T) {
	exptectedAddress := "10.1.0.21:6379"
	if actualAddress := redisPoolObj.GetAddress(); actualAddress != exptectedAddress {
		t.Errorf("GetAddress should be %s, but got %s", exptectedAddress, actualAddress)
	}
}

func TestTest(t *testing.T) {
	/**
	if err := redisPoolObj.Test(); err != nil {
		t.Errorf("Test connection failed, err:%s", err)
		os.Exit(1)
	}
	**/
}

func TestExists(t *testing.T) {
	key := "testExists"
	if exists, err := redisPoolObj.Exists(key); err != nil {
		t.Errorf("Exists failed,err:%s", err)
	} else if exists {
		t.Errorf("it should be not exists, but now exists")
	}

	redisPoolObj.Set(key, "test")
	if exists, err := redisPoolObj.Exists(key); err != nil {
		t.Errorf("Exists failed,err:%s", err)
	} else if !exists {
		t.Errorf("it should be exists, but now not exists")
	}

	if count, err := redisPoolObj.Del(key); err != nil {
		t.Errorf("Del failed, err:%s", err)
	} else if count != 1 {
		t.Errorf("Del should return 1, but now return %d", count)
	}
}

func TestKeys(t *testing.T) {
	if keyList, err := redisPoolObj.Keys("*"); err != nil {
		t.Errorf("Keys failed,err:%s", err)
		return
	} else if len(keyList) != 0 {
		t.Errorf("there should be no keys, but now got %d.", len(keyList))
		return
	}

	redisPoolObj.Set("key1", "key1")
	if keyList, err := redisPoolObj.Keys("*"); err != nil {
		t.Errorf("Keys failed,err:%s", err)
		return
	} else if len(keyList) != 1 {
		t.Errorf("there should be 1 keys, but now got %d.", len(keyList))
		return
	}

	redisPoolObj.Set("key2", "key2")
	if keyList, err := redisPoolObj.Keys("*"); err != nil {
		t.Errorf("Keys failed,err:%s", err)
		return
	} else if len(keyList) != 2 {
		t.Errorf("there should be 2 keys, but now got %d.", len(keyList))
		return
	}

	if count, err := redisPoolObj.Del("key1", "key2"); err != nil {
		t.Errorf("Del failed, err:%s", err)
	} else if count != 2 {
		t.Errorf("Del should return 2, but now return %d", count)
	}
}

func TestExpire(t *testing.T) {
	key := "expire"
	if success, err := redisPoolObj.Expire(key, 2); err != nil {
		t.Errorf("Expire failed, err:%s", err)
	} else if success {
		t.Errorf("Expire expected fail, but now got success")
	}

	redisPoolObj.Set(key, "test")

	if success, err := redisPoolObj.Expire(key, 2); err != nil {
		t.Errorf("Expire failed, err:%s", err)
	} else if !success {
		t.Errorf("Expire expected success, but now got fail")
	}

	time.Sleep(2 * time.Second)
	if exists, err := redisPoolObj.Exists(key); err != nil {
		t.Errorf("Exists failed,err:%s", err)
	} else if exists {
		t.Errorf("it should be not exists, but now exists")
	}
}

func TestGet(t *testing.T) {
	key := "get"

	if _, exists, err := redisPoolObj.Get(key); err != nil {
		t.Errorf("Get failed, err:%s", err)
	} else if exists {
		t.Errorf("Get should be not exists, but now exists.")
	}

	redisPoolObj.Set(key, "set")

	if value, exists, err := redisPoolObj.Get(key); err != nil {
		t.Errorf("Get failed, err:%s", err)
	} else if !exists {
		t.Errorf("Get should be  exists, but now not exists.")
	} else if value != "set" {
		t.Errorf("Get value should be %s, but now got %s", "set", value)
	}

	redisPoolObj.Del(key)
}

func TestHGet(t *testing.T) {
	key := "hget"
	field := "name"
	value := "jordan"

	if _, exists, err := redisPoolObj.HGet(key, field); err != nil {
		t.Errorf("HGET failed, err:%s", err)
	} else if exists {
		t.Errorf("HGET should be not exists, but now exists")
	}

	redisPoolObj.HSet(key, field, value)

	if actualValue, exists, err := redisPoolObj.HGet(key, field); err != nil {
		t.Errorf("HGET failed, err:%s", err)
	} else if !exists {
		t.Errorf("HGET should be exists, but now not exists")
	} else if value != actualValue {
		t.Errorf("HGET expected got %s, but now got %s", value, actualValue)
	}

	redisPoolObj.Del(key)
}

func TestSet2(t *testing.T) {
	key := "set2"
	val := "set2"
	expireVal := 10
	if err := redisPoolObj.Set2(key, val, Expire_Seond, expireVal); err != nil {
		t.Errorf("set2 error:%v", err)
	}
}

func TestSetDetail(t *testing.T) {
	key := "SetDetail"
	val := "SetDetail"
	expireVal := 20
	if err := redisPoolObj.SetDetail(key, val, Expire_Seond, expireVal, Set_NX); err != nil {
		t.Errorf("SetDetail error:%v", err)
	}

	val = "SetDetail2"
	if err := redisPoolObj.SetDetail(key, val, Expire_Seond, expireVal, Set_NX); err != nil {
		t.Errorf("SetDetail error:%v", err)
	}
}

func TestSetNX(t *testing.T) {
	key := "TestSetNx"
	value := "value"

	if success, err := redisPoolObj.SetNX(key, value); err != nil {
		t.Errorf("SETNX failed, err:%s", err)
	} else if !success {
		t.Errorf("SETNX should be succeed, but now failed")
	}

	if actualValue, exists, err := redisPoolObj.Get(key); err != nil {
		t.Errorf("SETNX.GET failed, err:%s", err)
	} else if !exists {
		t.Errorf("SETNX.GET should be exist, but now not exist")
	} else if actualValue != value {
		t.Errorf("SETNX.GET expected %s, but now get %s", value, actualValue)
	}

	value = "newvalue"
	if success, err := redisPoolObj.SetNX(key, value); err != nil {
		t.Errorf("SETNX failed, err:%s", err)
	} else if success {
		t.Errorf("SETNX should be failed, but now succeed")
	}
}

func TestHSet(t *testing.T) {
	key := "hget"
	field1 := "name"
	value1 := "jordan"
	field2 := "age"
	value2 := 32

	if err := redisPoolObj.HSet(key, field1, value1); err != nil {
		t.Errorf("HSET failed, err:%s", err)
	}

	if actualValue, exists, err := redisPoolObj.HGet(key, field1); err != nil {
		t.Errorf("HGET failed, err:%s", err)
	} else if !exists {
		t.Errorf("HGET should be exists, but now not exists")
	} else if actualValue != value1 {
		t.Errorf("HGET expected got %d, but now got %s", value1, actualValue)
	}

	if err := redisPoolObj.HSet(key, field2, value2); err != nil {
		t.Errorf("HSET failed, err:%s", err)
	}

	if actualValue, exists, err := redisPoolObj.HGet(key, field2); err != nil {
		t.Errorf("HGET failed, err:%s", err)
	} else if !exists {
		t.Errorf("HGET should be exists, but now not exists")
	} else if actualValue2, err := strconv.Atoi(actualValue); err != nil || actualValue2 != value2 {
		t.Errorf("HGET expected got %d, but now got %s", value2, actualValue)
	}

	redisPoolObj.Del(key)
}

type Player struct {
	Name string
	Age  int
}

func TestHGetAll(t *testing.T) {
	key := "player"

	p := &Player{}
	if exists, err := redisPoolObj.HGetAll(key, p); err != nil {
		t.Errorf("HGETALL failed, err:%s", err)
	} else if exists {
		t.Errorf("HGETALL should be not exists, but now exists.")
	}

	p1 := &Player{
		Name: "jordan",
		Age:  32,
	}

	if err := redisPoolObj.HMSet(key, p1); err != nil {
		t.Errorf("HMSET failed, err:%s", err)
	}

	if exists, err := redisPoolObj.HGetAll(key, p); err != nil {
		t.Errorf("HGETALL failed, err:%s", err)
	} else if !exists {
		t.Errorf("HGETALL should be exists, but now not exists.")
	} else {
		fmt.Printf("player:%v\n", p)
	}

	redisPoolObj.Del(key)
}

func TestLRange(t *testing.T) {
	key := "list"
	start := 0
	stop := -1

	if list, err := redisPoolObj.LRange(key, start, stop); err != nil {
		t.Errorf("LRANGE failed, err:%s", err)
	} else if len(list) != 0 {
		t.Errorf("LRANGE expected 0 item, but now %d items", len(list))
	}

	if newCount, err := redisPoolObj.LPush(key, "1"); err != nil {
		t.Errorf("LPUSH failed, err:%s", err)
	} else if newCount != 1 {
		t.Errorf("LPUSH expected got 1, but now got %d", newCount)
	}

	if newCount, err := redisPoolObj.RPush(key, "2"); err != nil {
		t.Errorf("LPUSH failed, err:%s", err)
	} else if newCount != 2 {
		t.Errorf("LPUSH expected got 2, but now got %d", newCount)
	}

	if newCount, err := redisPoolObj.RPush(key, "3"); err != nil {
		t.Errorf("LPUSH failed, err:%s", err)
	} else if newCount != 3 {
		t.Errorf("LPUSH expected got 3, but now got %d", newCount)
	}

	if newCount, err := redisPoolObj.RPush(key, "3"); err != nil {
		t.Errorf("LPUSH failed, err:%s", err)
	} else if newCount != 4 {
		t.Errorf("LPUSH expected got 4, but now got %d", newCount)
	}

	if list, err := redisPoolObj.LRange(key, start, stop); err != nil {
		t.Errorf("LRANGE failed, err:%s", err)
	} else if len(list) != 4 {
		t.Errorf("LRANGE expected 4 item, but now %d items", len(list))
	} else {
		for _, item := range list {
			fmt.Println(item)
		}
	}

	if removeCount, err := redisPoolObj.LRem(key, 0, "3"); err != nil {
		t.Errorf("LRem failed, err:%s", err)
	} else if removeCount != 2 {
		t.Errorf("LREM expected got 2, but now got %d", removeCount)
	}

	if item, exists, err := redisPoolObj.LPop(key); err != nil {
		t.Errorf("LPOP failed, err:%s", err)
	} else if !exists {
		t.Errorf("LPOP should be exists, but now not exists")
	} else if item != "1" {
		t.Errorf("LPOP should got 1, but now got %s", item)
	}

	if item, exists, err := redisPoolObj.RPop(key); err != nil {
		t.Errorf("RPOP failed, err:%s", err)
	} else if !exists {
		t.Errorf("RPOP should be exists, but now not exists")
	} else if item != "2" {
		t.Errorf("RPOP should got 2, but now got %s", item)
	}

	if _, exists, err := redisPoolObj.RPop(key); err != nil {
		t.Errorf("RPOP failed, err:%s", err)
	} else if exists {
		t.Errorf("RPOP should be not exists, but now exists")
	}
}

func TestIncr(t *testing.T) {
	key := "OrderSeed"
	var count int64
	var expected int64
	var err error
	var effected int

	if count, err = redisPoolObj.Incr(key); err != nil {
		t.Errorf("Incr falied, err:%s", err)
	}

	expected = 1
	if count != expected {
		t.Errorf("Expected %d, but now got %d.", expected, count)
	}

	if count, err = redisPoolObj.Incr(key); err != nil {
		t.Errorf("Incr falied, err:%s", err)
	}
	expected = 2
	if count != expected {
		t.Errorf("Expected %d, but now got %d.", expected, count)
	}

	if effected, err = redisPoolObj.Del(key); err != nil {
		t.Errorf("Del failed, err:%s", err)
	} else if effected != 1 {
		t.Errorf("Del should return 1, but now return %d", count)
	}
}

func TestIncrBy(t *testing.T) {
	key := "OrderSeed"
	var count int64
	var expected int64
	var err error
	var effected int

	if count, err = redisPoolObj.IncrBy(key, 2); err != nil {
		t.Errorf("Incr falied, err:%s", err)
	}

	expected = 2
	if count != expected {
		t.Errorf("Expected %d, but now got %d.", expected, count)
	}

	if count, err = redisPoolObj.IncrBy(key, 3); err != nil {
		t.Errorf("Incr falied, err:%s", err)
	}
	expected = 5
	if count != expected {
		t.Errorf("Expected %d, but now got %d.", expected, count)
	}

	if effected, err = redisPoolObj.Del(key); err != nil {
		t.Errorf("Del failed, err:%s", err)
	} else if effected != 1 {
		t.Errorf("Del should return 1, but now return %d", count)
	}
}

func TestDecr(t *testing.T) {
	key := "OrderSeed"
	var count int64
	var expected int64
	var err error
	var effected int

	if count, err = redisPoolObj.Decr(key); err != nil {
		t.Errorf("Incr falied, err:%s", err)
	}

	expected = -1
	if count != expected {
		t.Errorf("Expected %d, but now got %d.", expected, count)
	}

	if count, err = redisPoolObj.Decr(key); err != nil {
		t.Errorf("Incr falied, err:%s", err)
	}
	expected = -2
	if count != expected {
		t.Errorf("Expected %d, but now got %d.", expected, count)
	}

	if effected, err = redisPoolObj.Del(key); err != nil {
		t.Errorf("Del failed, err:%s", err)
	} else if effected != 1 {
		t.Errorf("Del should return 1, but now return %d", count)
	}
}

func TestDecrBy(t *testing.T) {
	key := "OrderSeed"
	var count int64
	var expected int64
	var err error
	var effected int

	if count, err = redisPoolObj.DecrBy(key, 2); err != nil {
		t.Errorf("Incr falied, err:%s", err)
	}

	expected = -2
	if count != expected {
		t.Errorf("Expected %d, but now got %d.", expected, count)
	}

	if count, err = redisPoolObj.DecrBy(key, 3); err != nil {
		t.Errorf("Incr falied, err:%s", err)
	}
	expected = -5
	if count != expected {
		t.Errorf("Expected %d, but now got %d.", expected, count)
	}

	if effected, err = redisPoolObj.Del(key); err != nil {
		t.Errorf("Del failed, err:%s", err)
	} else if effected != 1 {
		t.Errorf("Del should return 1, but now return %d", count)
	}
}

func TestSAdd(t *testing.T) {
	key := "TestSAdd"

	redisPoolObj.Del(key)
	newCount, err := redisPoolObj.SAdd(key, "1")
	if err != nil {
		t.Error(err)
		return
	}
	if newCount <= 0 {
		t.Errorf("影响记录数和写入记录数不匹配1：%v", newCount)
	}

	newCount, err = redisPoolObj.SAdd(key, "2", "3")
	if err != nil {
		t.Error(err)
		return
	}
	if newCount != 2 {
		t.Errorf("影响记录数和写入记录数不匹配2：%v", newCount)
	}
}

func TestSCard(t *testing.T) {
	key := "TestSCard"
	values := []string{"1", "2", "3"}
	redisPoolObj.Del(key)
	if _, err := redisPoolObj.SAdd(key, values...); err != nil {
		t.Error(err)
		return
	}

	count, err := redisPoolObj.SCard(key)
	if err != nil {
		t.Error(err)
		return
	}
	if count != len(values) {
		t.Errorf("TestSCard获取到的记录数不正确：%v", count)
	}
}

func TestSIsMember(t *testing.T) {
	key := "TestSIsMember"
	values := []string{"1"}
	redisPoolObj.Del(key)
	if _, err := redisPoolObj.SAdd(key, values...); err != nil {
		t.Error(err)
		return
	}

	isMember, err := redisPoolObj.SIsMember(key, "1")
	if err != nil {
		t.Error(err)
		return
	}
	if isMember == false {
		t.Errorf("TestSIsMember本该是成员，但返回不是成员")
		return
	}

	isMember, err = redisPoolObj.SIsMember(key, "2")
	if err != nil {
		t.Error(err)
		return
	}
	if isMember {
		t.Errorf("TestSIsMember本不是成员，但返回是成员")
		return
	}
}

func TestSMembers(t *testing.T) {
	key := "TestSMembers"
	values := []string{"1"}
	redisPoolObj.Del(key)
	if _, err := redisPoolObj.SAdd(key, values...); err != nil {
		t.Error(err)
		return
	}

	result, err := redisPoolObj.SMembers(key)
	if err != nil {
		t.Error(err)
		return
	}
	if len(result) != len(values) {
		t.Errorf("TestSMembers数量不正确")
		return
	}

	redisPoolObj.Del(key)
	result, err = redisPoolObj.SMembers(key)
	if err != nil {
		t.Error(err)
		return
	}
	if len(result) != 0 {
		t.Errorf("TestSMembers数量不正确2")
		return
	}
}

func TestSRandMember(t *testing.T) {
	key := "TestSRandMember"
	values := []string{"1"}
	redisPoolObj.Del(key)
	if _, err := redisPoolObj.SAdd(key, values...); err != nil {
		t.Error(err)
		return
	}

	result, err := redisPoolObj.SRandMember(key, 1)
	if err != nil {
		t.Error(err)
		return
	}
	if len(result) != len(values) {
		t.Errorf("TestSRandMember数量不正确")
		return
	}

	redisPoolObj.Del(key)
	result, err = redisPoolObj.SMembers(key)
	if err != nil {
		t.Error(err)
		return
	}
	if len(result) != 0 {
		t.Errorf("TestSRandMember数量不正确2")
		return
	}
}

func TestSPop(t *testing.T) {
	key := "TestSPop"
	values := []string{"1"}
	redisPoolObj.Del(key)
	if _, err := redisPoolObj.SAdd(key, values...); err != nil {
		t.Error(err)
		return
	}

	result, err := redisPoolObj.SPop(key)
	if err != nil {
		t.Error(err)
		return
	}
	if result != values[0] {
		t.Errorf("TestSPop获取结果不正确")
		return
	}

	var count int
	count, err = redisPoolObj.SCard(key)
	if err != nil {
		t.Error(err)
		return
	}
	if count != 0 {
		t.Errorf("TestSPop未正常移除")
		return
	}
}

func TestSRem(t *testing.T) {
	key := "TestSRem"
	values := []string{"1"}
	redisPoolObj.Del(key)
	if _, err := redisPoolObj.SAdd(key, values...); err != nil {
		t.Error(err)
		return
	}

	result, err := redisPoolObj.SRem(key, "1")
	if err != nil {
		t.Error(err)
		return
	}
	if result != 1 {
		t.Errorf("TestSRem数量不正确")
		return
	}
}
