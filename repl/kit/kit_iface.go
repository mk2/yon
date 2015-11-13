package kit

type ReplClient interface {
	ShowHelp(string) string
	Eval(string) string
	EvalFile(string) (string, error)
}

type Repl interface {
	GetClient() ReplClient
	GetPrimaryServer() ReplServer
	GetServers() []ReplServer
}


type ReplServer interface {
	Send(string) error
	Receive(int) string
}
