package cable

import (
	"log"
	"net/http"
	"time"

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

	defer wsConnMap.Delete(userId)
	defer ws.Close(websocket.StatusNormalClosure, "Finished work")
	wsConnMap.Add(userId, ws)

	for i := 0; i < 100; i++ {
		ws.Write(c, websocket.MessageText, []byte(userId))
		time.Sleep(1 * time.Second)
	}
}
