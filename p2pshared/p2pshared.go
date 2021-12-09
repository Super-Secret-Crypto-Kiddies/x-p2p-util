package p2pshared

import (
	"x-bootstrap-node/xutil"

	"github.com/perlin-network/noise"
)

var BootstrapNodes = [...]string{
	"143.198.176.143", // DigitalOcean host
	// "10.142.89.192",   // rpi on LAN
	// "10.142.26.69", // thinkpad on LAN
}

func RegisterNodeMessages(n *noise.Node) error {

	n.RegisterMessage(xutil.PeerInfo{}, xutil.UnmarshalPeerInfo)

	n.RegisterMessage(xutil.PaymentIntent{}, xutil.UnmarshalPaymentIntent)
	n.RegisterMessage(xutil.PaymentResponse{}, xutil.UnmarshalPaymentResponse)

	n.RegisterMessage(xutil.PaymentCancellation{}, xutil.UnmarshalPaymentCancellation)
	n.RegisterMessage(xutil.PaymentCancellationResponse{}, xutil.UnmarshalPaymentCancellationResponse)

	return nil
}
