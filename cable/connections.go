package cable

import (
	"sync"

	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
)

const wsConnMap = "WS_CONN_MAP"

var globalConnMap = WsConnectionMap{
	conns: make(map[string]*websocket.Conn),
	mut:   sync.Mutex{},
}

type WsConnectionMap struct {
	conns map[string]*websocket.Conn
	mut   sync.Mutex
}

func (wsConnMap *WsConnectionMap) Add(id string, ws *websocket.Conn) {
	wsConnMap.mut.Lock()
	wsConnMap.conns[id] = ws
	wsConnMap.mut.Unlock()
}

func (wsConnMap *WsConnectionMap) Delete(id string) {
	wsConnMap.mut.Lock()
	delete(wsConnMap.conns, id)
	wsConnMap.mut.Unlock()
}

func FetchWsConnectionMap(c *gin.Context) *WsConnectionMap {
	result, _ := c.Get(wsConnMap)
	return result.(*WsConnectionMap)
}

func AddWsConnectionMap(c *gin.Context) {
	c.Set(
		wsConnMap,
		&globalConnMap,
	)
}
