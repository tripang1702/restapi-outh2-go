package controller

// Controller example
type Controller struct {
	Conn Connection
}

type Connection struct {
	Server, Port, User, Password, Database string
}

// NewController example
func NewController() *Controller {
	return &Controller{}
}

func (co *Controller) SetParamEnv(server string, port string, user string, password string, database string) {
	var con Connection

	con.Server = server
	con.Port = port
	con.User = user
	con.Password = password
	con.Database = database

	co.Conn = con
}

func (co *Controller) NewConnection() *Connection {
	return &co.Conn
}
