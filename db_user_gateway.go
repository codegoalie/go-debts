package main

import (
	"fmt"
	"go-debts/interfaces"
)

type dbUserGateway struct {
	handler interfaces.DbHandler
}

func (gateway dbUserGateway) fetchDebitorByUserId(userId int) debitor {
	row := gateway.handler.Query(fmt.Sprintf("SELECT debitor_id FROM users WHERE id = %d LIMIT 1", userId))
	var debitorId int
	row.Next()
	row.Scan(&debitorId)
	row = gateway.handler.Query(fmt.Sprintf("SELECT name FROM debitors WHERE id = %d LIMIT 1", debitorId))
	var name string
	row.Next()
	row.Scan(&name)
	return debitor{id: debitorId, name: name}
}
