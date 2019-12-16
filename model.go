package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type student struct {
	gorm.Model `json:"-"`
	Content string `json:"stu_id"`
	DDL string `json:"phone"`
}

type messageForm struct {
	TId string `json:"stu_id"`
	Content string `json:"stu_id"`
	DDL string `json:"phone"`
}


func (s *Service) makeErrJSON(httpStatusCode int, errCode int, msg interface{}) (int, interface{}) {
	return httpStatusCode, gin.H{"error": errCode, "msg": fmt.Sprint(msg)}
}

func (s *Service) makeSuccessJSON(data interface{}) (int, interface{}) {
	return 200, gin.H{"error": 0, "msg": "success", "data": data}
}
