package main

import "fmt"

type client struct {}

func (c *client) searchingInformationInNetwork(net network) {
	net.searchingInformationInGoogle()
}

type network interface {
	searchingInformationInGoogle()
}

type google struct {}

func (g *google) searchingInformationInGoogle()  {
	fmt.Println("The client searches for information on Google")
}

type opera struct {}

func (o *opera) searchingInformationInOpera() {
	fmt.Println("The client searches for information on Google")
}

type operaAdapter struct {
	operaNetwork *opera
}

func (oA *operaAdapter) searchingInformationInGoogle()  {
	fmt.Println("User opens Google in Opera")
	oA.operaNetwork.searchingInformationInOpera()
}

func asd(){
	client := &client{}
	google := &google{}
	client.searchingInformationInNetwork(google)

	opera := &opera{}
	//will give error
	//client.searchingInformationInNetwork(opera)

	operaAdapter := &operaAdapter{opera}
	client.searchingInformationInNetwork(operaAdapter)
}