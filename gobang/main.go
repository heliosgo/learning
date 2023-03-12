package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

func main() {
	r := gin.Default()
	m := melody.New()
	lobby := newLobby()
	go lobby.run()

	r.GET("/ws", func(c *gin.Context) {
		name, ok := c.GetQuery("name")
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"code": 10001,
				"msg":  "name is null",
			})
		}
		m.HandleRequestWithKeys(c.Writer, c.Request, map[string]interface{}{
			"name":   name,
			"player": newPlayer(name, lobby),
		})
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		sli := strings.Split(string(msg), ":")
		if len(sli) < 2 {
			return
		}
		op, val := sli[0], sli[1]
		player := s.Keys["player"].(*Player)
		switch op {
		case "table":
			player.sitDown(val)
			s.Keys["table"] = val
			m.BroadcastFilter(
				[]byte(fmt.Sprintf("欢迎 %s 进入房间", player.name)),
				func(p *melody.Session) bool {
					return p.Keys["table"] == val
				},
			)

		case "drop":
			xy := strings.Split(val, ",")
			if len(xy) < 2 {
				return
			}
			x, _ := strconv.Atoi(xy[0])
			y, _ := strconv.Atoi(xy[1])
			win, ok := player.drop(x, y)
			if !ok {
				return
			}
			m.BroadcastFilter(
				[]byte(player.table.getStringState()),
				func(p *melody.Session) bool {
					return p.Keys["table"] == p.Keys["table"]
				},
			)
			if win {
				m.BroadcastFilter(
					[]byte(fmt.Sprintf("winner: %s !", player.name)),
					func(p *melody.Session) bool {
						return p.Keys["table"] == p.Keys["table"]
					},
				)
			}
		}
	})

	r.Run(":5001")
}
