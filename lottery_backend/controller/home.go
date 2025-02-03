package controller

import (
	"crypto/rand"
	"lottery/database"
	"lottery/model"
	"lottery/utils"
	"math/big"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func GetHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func GetAllPrize(c *gin.Context) {
	var inventorys []model.Inventory
	database.DB.Find(&inventorys)

	for i := range inventorys {
		inventorys[i].Count = 0
	}
	c.JSON(200, gin.H{
		"inventorys": inventorys,
	})
}

func Lottery(c *gin.Context) { //解决高并发下redis库存保持原子性问题
	for i := 0; i < 10; i++ {
		dataCount := make(map[int]int, 20) //记录每一件奖品的数量
		dataInfo := make(map[int]string, 20)
		//设立抽奖算法
		conn := database.RedisPool.Get()
		defer conn.Close()

		// 获取所有奖品的数量
		var total = 0
		inventoryKeys, err := redis.Strings(conn.Do("KEYS", "*"))
		if err != nil {
			c.JSON(500, gin.H{"error": "获取奖品列表失败"})
			return
		}

		utils.Blob(inventoryKeys)
		// 遍历所有奖品ID,累加数量
		var i = 1
		for _, key := range inventoryKeys {
			count, err := redis.Int(conn.Do("GET", key))
			if count == 0 || count < 0 {
				continue
			}
			if err != nil {
				continue
			}
			total += count
			dataInfo[i] = key
			dataCount[i] = total
			i++
		}

		if len(dataCount) == 0 {
			c.JSON(200, gin.H{
				"message": "奖品已抽完",
				"id":      "0",
			})
			return
		}

		//设置二分查找区间算法
		//map[1:1000 2:2000 3:2200 4:2500 5:3500 6:4500 7:4600 8:5100 9:5500]
		//map[1:10 2:11 3:12 4:13 5:15 6:16 7:17 8:18 9:19]
		n, err := rand.Int(rand.Reader, big.NewInt(int64(total)))
		if err != nil {
			panic(err)
		}
		target := int(n.Int64()) + 1 //[1,tatal]
		var id = utils.Binary(target, dataCount, dataInfo)
		_, err2 := conn.Do("DECR", id) //redis本身是单线程的，所以这里不用担心并发问题
		count, _ := redis.Int(conn.Do("GET", id))
		if err2 != nil {
			c.JSON(500, gin.H{"error": "奖品数量减少失败"})
			return
		} else if count < 0 { //但并发情况可能导致负数
			continue //进行后端第二次抽奖
		}

		c.JSON(200, gin.H{
			"message": "抽奖成功",
			"id":      id,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "抽奖失败",
		"id":      "0",
	})
}
