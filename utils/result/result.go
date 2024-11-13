package result

const (
	OK             = 200
	NoContent      = 204
	PartialContent = 206
)

var resultMsg = map[int]string{
	OK:             "成功",
	NoContent:      "成功",
	PartialContent: "成功",
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func ResultMsg(code int) string {
	return resultMsg[code]
}

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 2XX 成功
// 200 OK ：表示从客户端发来的请求在服务器端被正常处理了
// 204 No Content：表示服务器接收的请求已经成功处理，但是返回的响应报文不包含实体的主体部分。
// 206 Partial Content ：表示客户端进行了范围请求，并且服务器成功执行了这部分Get请求。

// 3XX 重定向
// 301 Moved Permanently ：永久重定向，表示请求的资源已经被分配了新的URI，以后使用资源现在所指的URI
// 302 Found ：临时重定向。该状态码表示请求的资源被分配了新的URI 希望用户本次能使用新的URI进行访问
// 303 See Other：请求对应的资源存在着另一个URI，应使用GET方法重新定向获取请求的资源
// 303和302状态相似，303明确表示应当采用GET方法获取
// 304 Not Modified ： 表示客户端发送附带条件的请求时，服务端允许请求访问资源，但未满足条件

// 4XX 客户端错误
// 400 Bad Request : 请求报文存在语法错误
// 401 Unauthorized ：发送的请求需要通过HTTP认证，若之前已经进行过1次请求，表示用户认证失败
// 403 Forbidden ： 该状态码表明对请求资源的访问被服务器拒绝（可能是权限问题）
// 404 Not Found ：服务器上无法找到请求的资源，或者服务端拒绝请求且不想说明理由

// 5XX 服务端错误
// 500 Internal Server Error : 服务端在执行请求时发生错误
// 503 Service Unavailable :表示服务器暂时处于超负载或者在停机维护，现在无法处理请求。
