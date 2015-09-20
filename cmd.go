package shot

type Cmd struct {
	Url  string
	Done chan struct{}
}
