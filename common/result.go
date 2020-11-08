package common

// Result 客户端与服务器数据交互
type Result struct {
	Status int         // Status: 200 success
	Data   interface{} // Return data
	Msg    string      // Return message
}
