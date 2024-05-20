package entities

type RPCdata struct {
	Name string
	To   string
	Args []interface{}
	Err  string
}
