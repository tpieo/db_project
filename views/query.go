package views

import (
	"SZUCourse/global"

	"github.com/gin-gonic/gin"
)

func Query(c *gin.Context) {
	q := c.Query("q")
	courses1 := make([]Course, 0, 20)
	courses2 := make([]Course, 0, 20)
	r := global.Db.Limit(20).Where("teacher LIKE ?", "%"+q+"%").Find(&courses1)
	if r.RowsAffected < 20 {
		global.Db.Limit(20-int(r.RowsAffected)).Where("name LIKE ?", "%"+q+"%").Find(&courses2)
	}

	courses := append(courses1, courses2...)

	h := make(map[int]interface{})
	for i, course := range courses {
		h[i] = course
	}
	c.JSON(200, h)
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
