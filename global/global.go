package global

import (
	"sync"

	"gorm.io/gorm"
)

var Db *gorm.DB

var Cache map[string]map[int]interface{} = make(map[string]map[int]interface{}, 300)

var QueryStringQueue []string = make([]string, 0, 500)

var Lock sync.RWMutex
