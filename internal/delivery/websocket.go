package delivery

// import (
// 	"context"
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"sync"
// 	"time"

// 	"github.com/gorilla/websocket"
// )

// type WsIncoming struct {
// 	Action string          `json:"action"`
// 	Method string          `json:"method"`
// 	Tag    string          `json:"tag"`
// 	Data   json.RawMessage `json:"data"`
// }

// type WsOutgoing struct {
// 	Action string `json:"action"`
// 	Tag    string `json:"tag"`
// 	Data   any    `json:"data,omitempty"`
// 	Error  string `json:"error,omitempty"`
// }

// // WsRouter is the global registry for WebSocket handlers, allowing other files
// // like http.go to register routes automatically.
// var WsRouter = make(map[string]map[string]func(context.Context, json.RawMessage) (any, error))

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

// const (
// 	pongWait   = 60 * time.Second
// 	pingPeriod = (pongWait * 9) / 10
// )

// // HandleWebSocket manages upgrading the HTTP connection to a WebSocket.
// func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println("WebSocket upgrade failed:", err)
// 		return
// 	}
// 	defer conn.Close()

// 	log.Println("🔌 WebSocket connection established")

// 	var writeMutex sync.Mutex

// 	// Helper to safely write messages concurrently
// 	writeMsg := func(msgType int, data []byte) error {
// 		writeMutex.Lock()
// 		defer writeMutex.Unlock()
// 		return conn.WriteMessage(msgType, data)
// 	}

// 	// Helper to safely write JSON concurrently
// 	writeJSON := func(v any) error {
// 		writeMutex.Lock()
// 		defer writeMutex.Unlock()
// 		return conn.WriteJSON(v)
// 	}

// 	conn.SetReadDeadline(time.Now().Add(pongWait))
// 	conn.SetPongHandler(func(string) error {
// 		conn.SetReadDeadline(time.Now().Add(pongWait))
// 		return nil
// 	})

// 	// Start a background goroutine to send periodic automatic pings
// 	go func() {
// 		ticker := time.NewTicker(pingPeriod)
// 		defer ticker.Stop()
// 		for range ticker.C {
// 			if err := writeMsg(websocket.PingMessage, nil); err != nil {
// 				return // Connection closed or failed, stop pinging
// 			}
// 		}
// 	}()

// 	for {
// 		msgType, p, err := conn.ReadMessage()
// 		if err != nil {
// 			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
// 				log.Printf("Error reading message: %v", err)
// 			} else {
// 				log.Printf("WebSocket connection closed: %v", err)
// 			}
// 			break
// 		}

// 		// Ignore non-text messages (e.g., binary frames)
// 		if msgType != websocket.TextMessage {
// 			continue
// 		}

// 		// Handle manual text-based "ping" from terminal clients like wscat
// 		if string(p) == "ping" {
// 			_ = writeMsg(websocket.TextMessage, []byte("pong"))
// 			continue
// 		}

// 		go func(msg []byte) {
// 			var incoming WsIncoming
// 			if err := json.Unmarshal(msg, &incoming); err != nil {
// 				log.Printf("Error unmarshaling ws message: %v", err)
// 				_ = writeJSON(WsOutgoing{Error: "Invalid message format"})
// 				return
// 			}

// 			handlerFunc := WsRouter[incoming.Method][incoming.Action]
// 			if handlerFunc == nil {
// 				log.Printf("No websocket handler found for method %s, action %s", incoming.Method, incoming.Action)
// 				_ = writeJSON(WsOutgoing{Action: incoming.Action, Tag: incoming.Tag, Error: "Action not found"})
// 				return
// 			}

// 			data, handlerErr := handlerFunc(r.Context(), incoming.Data)

// 			outgoing := WsOutgoing{
// 				Action: incoming.Action,
// 				Tag:    incoming.Tag,
// 				Data:   data,
// 			}
// 			if handlerErr != nil {
// 				outgoing.Error = handlerErr.Error()
// 			}

// 			if writeErr := writeJSON(outgoing); writeErr != nil {
// 				log.Println("WebSocket write error:", writeErr)
// 			}
// 		}(p)
// 	}
// }
