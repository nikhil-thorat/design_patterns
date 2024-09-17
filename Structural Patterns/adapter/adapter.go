package main

import "fmt"

type Client struct{}

type Computer interface {
	InsertIntoLightningPort()
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("INSERTING LIGHTNING CONNECTOR INTO COMPUTER")
	com.InsertIntoLightningPort()
}

type Mac struct{}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("LIGHTNING CONNECTOR PLUGGED INTO MAC")
}

type Windows struct{}

func (w *Windows) InsertIntoLightningPort() {
	fmt.Println("LIGHTNING CONNECTOR PLUGGED INTO WINDOWS")
}

type WindowsAdapter struct {
	windowsMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("CONVERTING LIGHTNING SIGNAL TO USB")
	w.windowsMachine.InsertIntoLightningPort()
}

func main() {

	client := &Client{}

	mac := &Mac{}
	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{windowsMachine: windowsMachine}
	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
