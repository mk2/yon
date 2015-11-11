package kit

type Repl interface {
	GetClient() ReplClient
	GetPrimaryServer() ReplServer
	GetServers() []ReplServer
	Eval(string) string
	EvalFile(string) (string, error)
}

type ReplClient interface {
	ShowHelp(string) string
}

type ReplServer interface {
	Start() error
	Send(string) error
	Receive(int) (string, error)
}
