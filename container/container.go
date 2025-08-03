
package container

type getRequest struct {
	Name string
	ResponseChannel chan string
}

type Container struct {
	nameToValue map[string]string
	getChannel  chan *getRequest
}

func (c *Container) Init() {
	c.nameToValue = make(map[string]string)
	c.getChannel = make(chan *getRequest)
	go func() {
        for { // implement a way to stop this :')
			gr := <-c.getChannel
			gr.ResponseChannel<-c.nameToValue[gr.Name]
		}
	}()
}

func (c *Container) Get(name string)string {
	rc := make(chan string)
	gr := &getRequest{ Name: name, ResponseChannel: rc }
	c.getChannel<-gr
	return <-gr.ResponseChannel
}

func (c *Container) Add(name string, value string) {
    c.nameToValue[name] = value
}


