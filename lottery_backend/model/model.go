package model

import (
	"lottery/database"
)

func init() {
	conn := database.RedisPool.Get()
	defer conn.Close()

	// 获取key为10的值
	val, err := conn.Do("GET", "10")
	if err != nil {
		return
	}

	// 如果值为0才执行
	if val == nil {
		var inventorys []Inventory
		database.DB.Find(&inventorys)

		// 将库存信息存入redis
		for _, inventory := range inventorys {
			conn.Do("APPEND", inventory.Id, inventory.Count)
		}
	}
}
