package main

import (
	"context"
	"fmt"
	"gobang/property"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/olahol/melody"
	"github.com/spf13/cast"
	"google.golang.org/grpc"
)

func main() {
	r := gin.Default()
	m := melody.New()
	lobby := newLobby()
	go lobby.run()

	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := property.NewPropertyClient(conn)

	r.GET("/ws", func(c *gin.Context) {
		key, ok := c.GetQuery("loginKey")
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"code": 10001,
				"msg":  "key is null",
			})
			return
		}
		mp := make(jwt.MapClaims)
		token, err := jwt.ParseWithClaims(key, mp, func(t *jwt.Token) (interface{}, error) {
			return []byte(accessSecret), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusOK, gin.H{
				"code": 10001,
				"msg":  "key is invalid",
			})
			return
		}
		uid := cast.ToInt64(mp["uid"])
		name := fmt.Sprintf("user: %d", uid)
		player := newPlayer(uid, name, lobby)
		_, ok = players.Load(name)
		if ok {
			c.JSON(http.StatusOK, gin.H{
				"code": 10001,
				"msg":  "user already online",
			})
			return
		}
		players.Store(name, player)

		m.HandleRequestWithKeys(c.Writer, c.Request, map[string]interface{}{
			"key":    key,
			"player": player,
			"uid":    mp["uid"],
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

				res, err := client.UpdateScore(context.Background(), &property.UpdateScoreReq{
					Uid:   player.uid,
					Score: 5,
				})
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(res)
			}
		}
	})

	r.Run(":5001")
}
