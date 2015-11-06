package kit

type Repl interface {
	GetClient() ReplClient
	GetPrimaryServer() ReplServer
	GetServers() []ReplServer
}

type ReplClient interface {
	Eval(string)
	EvalFile(string)
}

type ReplServer interface {
}
