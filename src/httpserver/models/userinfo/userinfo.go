package userinfo

import (
	"httpserver/utils"
)

type UserInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (self *UserInfo) SetValues() {
	self.Id = utils.RandInt()
	self.Name = "MimikFc7"
}
