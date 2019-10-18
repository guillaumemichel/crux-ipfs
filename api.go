package template

/*
The api.go defines the methods that can be called from the outside. Most
of the methods will take a roster so that the service knows which nodes
it should work with.

This part of the service runs on the client or the app.
*/

import (
	"go.dedis.ch/onet"

	"go.dedis.ch/cothority/v3"
	"go.dedis.ch/onet/v3/log"
	"go.dedis.ch/onet/v3/network"
)

// ServiceName is used for registration on the onet.
const ServiceName = "MyTestSimulation"

// Client is a structure to communicate with the template
// service
type Client struct {
	*onet.Client
}

// NewClient instantiates a new template.Client
func NewClient() *Client {
	return &Client{Client: onet.NewClient(cothority.Suite, ServiceName)}
}

// Clock chooses one server from the Roster at random. It
// sends a Clock to it, which is then processed on the server side
// via the code in the service package.
//
// Clock will return the time in seconds it took to run the protocol.
func (c *Client) Clock(r *onet.Roster) (*ClockReply, error) {
	dst := r.RandomServerIdentity()
	log.Lvl4("Sending message to", dst)
	reply := &ClockReply{}
	err := c.SendProtobuf(dst, &Clock{r}, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// Count will return the number of times `Clock` has been called on this
// service-node.
func (c *Client) Count(si *network.ServerIdentity) (int, error) {
	reply := &CountReply{}

	err := c.SendProtobuf(si, &Count{}, reply)
	if err != nil {
		return -1, err
	}
	return reply.Count, nil
}

// GenSecret asks a random node to generate a secret key
func (c *Client) GenSecret(r *onet.Roster) (*GenSecretReply, error) {
	dst := r.RandomServerIdentity()
	log.Lvl4("Sending message to", dst)
	reply := &GenSecretReply{}
	err := c.SendProtobuf(dst, &GenSecret{}, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil

}