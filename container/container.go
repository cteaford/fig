
package container

import (
    "fmt"
)

type getRequest struct {
	Name string
	ResponseChannel chan string
}

type addRequest struct {
	Name string
    Value string
}

type Container struct {
	nameToValue map[string]string
	requestChannel chan interface{}
}

func (c *Container) Init() {
	c.nameToValue = make(map[string]string)
	c.requestChannel = make(chan interface{})
	go func() {
        fmt.Println("supervisor routine started...")
        for { //TODO: implement a way to stop this :')
            request := <-c.requestChannel
            switch r := request.(type) {
                case getRequest:
                    r.ResponseChannel<-c.nameToValue[r.Name]
                case addRequest:
                    c.nameToValue[r.Name] = r.Value
                default:
                    fmt.Printf("Value: %v, Type: %T\n", r, r)

            }
		}
	}()
}

func (c *Container) Get(name string)string {
	rc := make(chan string)
	gr := getRequest{ Name: name, ResponseChannel: rc }
	c.requestChannel <- gr
    response := <-gr.ResponseChannel
    close(gr.ResponseChannel)
	return response
}

func (c *Container) Add(name string, value string) {
    ar := addRequest{ Name: name, Value: value }
    c.requestChannel <-ar
}


