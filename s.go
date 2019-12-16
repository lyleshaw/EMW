package main

import "github.com/gin-gonic/gin"

func (s *Service) AddNew(c *gin.Context) (int, interface{})  {
	Thing := net(things)
	err := c.ShouldBindJSON(Thing)
	if err != nil {
		return s.makeErrJSON()
	}
}
