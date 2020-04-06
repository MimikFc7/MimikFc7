package weather

import (
	"httpserver/utils"
)

type Weather struct {
	Id      int    `json:"id"`
	Feeling string `json:"feeling"`
}

func (self *Weather) SetValues() {
	self.Id = utils.RandInt()
	self.Feeling = utils.RandString()
}
