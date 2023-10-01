package cable

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
)

func HandleWebsocket(c *gin.Context) {
	userId := c.GetString("UserId")
	wsConnMap := FetchWsConnectionMap(c)
	ws, err := websocket.Accept(c.Writer, c.Request, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	wsConnMap.Add(userId, ws)

	go func() {
		defer wsConnMap.Delete(userId)
		defer ws.Close(websocket.StatusNormalClosure, "Finished work")

		for {
			_, msg, err := ws.Read(c)
			if err != nil {
				log.Printf("[WS=%s] Client disconnected: %s", userId, err)
				break
			}
			log.Printf("[WS=%s] Received message: %s", userId, msg)
		}
	}()
}
