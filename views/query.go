package views

import (
	"SZUCourse/global"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Query(c *gin.Context) {

	queryString := c.Query("q")

	if queryString == "" {
		c.JSON(200, map[int]string{0: "empty"})
		return
	}

	// read cache, readuce query database
	global.Lock.RLock()
	value, ok := global.Cache[queryString]
	global.Lock.RUnlock()
	if ok {
		c.JSON(200, value)
		fmt.Println("cache")
		return
	}

	courses := search(queryString)
	result := make(map[int]interface{})
	for i, course := range courses {
		result[i] = course
	}

	//high concurrency
	go cache(queryString, result)

	c.JSON(200, result)
}

func search(queryString string) []Course {
	courses1 := make([]Course, 0, 20)
	courses2 := make([]Course, 0, 20)
	r := global.Db.Limit(20).Where("teacher LIKE ?", "%"+queryString+"%").Find(&courses1)
	if r.RowsAffected < 20 {
		global.Db.Limit(20-int(r.RowsAffected)).Where("name LIKE ?", "%"+queryString+"%").Find(&courses2)
	}

	courses := append(courses1, courses2...)
	return courses
}

func cache(queryString string, result map[int]interface{}) {

	global.Lock.Lock()
	defer global.Lock.Unlock()

	global.QueryStringQueue = append(global.QueryStringQueue, queryString)
	if len(global.QueryStringQueue) > 200 {
		delete(global.Cache, global.QueryStringQueue[0])
		global.QueryStringQueue = global.QueryStringQueue[1:]
	}

	global.Cache[queryString] = result

}

type Course struct {
	Uid          int
	ID           string
	Name         string
	Dept         string
	Credit       string
	Duration     string
	Teacher      string
	Campus       string
	Type         string
	TimeAndPlace string
	Target       string
}
