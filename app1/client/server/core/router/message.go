package router
import 	(
	// "gitlab.com/meta-node/client/models"
	"gitlab.com/meta-node/client/utils"
)

type Message struct {
	Type string      `json:"type"`
	Msg  interface{} `json:"message"`
}
// WSMessage struct
type WSMessage struct {
	Action string `json:"action"`
	Topic  string `json:"topic"`
	Data   string `json:"data"`
}
type Message1 struct {
	Platform string       `json:"platform"`
	Command string       `json:"command"`
	// Success  bool        `json:"success"`
	IsSocket  bool        `json:"isSocket"`
	Data    *utils.ResultTransformer      `json:"data"`
}
