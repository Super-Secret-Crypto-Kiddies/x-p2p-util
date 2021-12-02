package xutil

import (
	"encoding/json"
	"fmt"
)

// Peer Info Requests should just be made as a string message that says "peerinfo"

type PeerInfo struct {
	Type       PeerType
	Currencies []string
}

type PeerType int

const (
	Node = iota
	Server
)

func (p PeerInfo) Marshal() []byte {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Failed to marshal payment cancellation: %s", err)
	}
	return bytes
}

func UnmarshalPeerInfo(bytes []byte) (PeerInfo, error) {
	res := PeerInfo{}
	if err := json.Unmarshal(bytes, &res); err != nil {
		fmt.Printf("Failed to unmarshal response: %s", err)
		return res, err
	}
	return res, nil
}
