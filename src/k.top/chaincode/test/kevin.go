package main

type Kevin struct {
	name string
}

type router struct {
}

func NewRouter(name string) *router {
	return nil
}

type ChaincodeFunc func(int) string

type ChaincodeHandlerFunc func(args []string, responseWriter *peerResponseWriter)

type peerResponseWriter interface {
}

func createAc(args []string, responseWriter *peerResponseWriter) {
	//return string(args[0])
}

func (r *router) handle(path string, fn ChaincodeHandlerFunc, middleware ...interface{}) *router {
	return nil
}

func (r *router) method(m string) *router {
	return nil
}

const (
	invokeMethod = `INVOKE`
	queryMethod  = `QUERY`
)

func main() {
	NewRouter(`chain-code-router`).handle(`createAc`, createAc).method(invokeMethod)
}
