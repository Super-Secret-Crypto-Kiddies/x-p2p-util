package xutil

import "github.com/perlin-network/noise"

func RegisterNodeMessages(n *noise.Node) error {

	n.RegisterMessage(PeerInfo{}, UnmarshalPeerInfo)

	n.RegisterMessage(PaymentIntent{}, UnmarshalPaymentIntent)
	n.RegisterMessage(PaymentResponse{}, UnmarshalPaymentResponse)

	n.RegisterMessage(PaymentCancellation{}, UnmarshalPaymentCancellation)
	n.RegisterMessage(PaymentCancellationResponse{}, UnmarshalPaymentCancellationResponse)

	return nil
}
