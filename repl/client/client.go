package client

type client struct {
}

func (c *client) ShowHelp(s string) string {

	return ""
}

func (c *client) Eval(s string) string {


}

func (c *client) EvalFile(f string) (string, error) {

	return "", nil
}
