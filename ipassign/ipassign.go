package ipassign

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/huin/goupnp/dcps/internetgateway2"
	nat "github.com/libp2p/go-libp2p-nat"
	"golang.org/x/sync/errgroup"
)

// NAT Traversal from libp2p

func GetNatMapping() string {
	nat, err := nat.DiscoverNAT(context.Background())
	if err != nil {
		fmt.Println("Failed to discover NAT.")
	}
	// defer nat.Close()
	mapping, err := nat.NewMapping("tcp", 9871)
	if err != nil {
		fmt.Println("Failed to initialize NAT mapping")
	}
	defer mapping.Close()
	addr, _ := mapping.ExternalAddr()
	fmt.Printf("Mapping %d to %s\n", mapping.InternalPort(), addr)
	return addr.String()
}

// Half-baked UPnP support - currently unused
// *Shamelessly ripped from https://github.com/huin/goupnp/blob/master/GUIDE.md

type RouterClient interface {
	AddPortMapping(
		NewRemoteHost string,
		NewExternalPort uint16,
		NewProtocol string,
		NewInternalPort uint16,
		NewInternalClient string,
		NewEnabled bool,
		NewPortMappingDescription string,
		NewLeaseDuration uint32,
	) (err error)

	GetExternalIPAddress() (
		NewExternalIPAddress string,
		err error,
	)

	GetGenericPortMappingEntry(NewPortMappingIndex uint16) (
		NewRemoteHost string,
		NewExternalPort uint16,
		NewProtocol string,
		NewInternalPort uint16,
		NewInternalClient string,
		NewEnabled bool,
		NewPortMappingDescription string,
		NewLeaseDuration uint32,
		err error,
	)

	// GetListOfPortMappings(
	// 	NewStartPort uint16,
	// 	NewEndPort uint16,
	// 	NewProtocol string,
	// 	NewManage bool,
	// 	NewNumberOfPorts uint16,
	// ) (
	// 	NewPortListing string,
	// 	err error,
	// )
}

func PickRouterClient(ctx context.Context) (RouterClient, error) {
	tasks, _ := errgroup.WithContext(ctx)
	// Request each type of client in parallel, and return what is found.
	var ip1Clients []*internetgateway2.WANIPConnection1
	tasks.Go(func() error {
		var err error
		ip1Clients, _, err = internetgateway2.NewWANIPConnection1Clients()
		return err
	})
	var ip2Clients []*internetgateway2.WANIPConnection2
	tasks.Go(func() error {
		var err error
		ip2Clients, _, err = internetgateway2.NewWANIPConnection2Clients()
		return err
	})
	var ppp1Clients []*internetgateway2.WANPPPConnection1
	tasks.Go(func() error {
		var err error
		ppp1Clients, _, err = internetgateway2.NewWANPPPConnection1Clients()
		return err
	})

	if err := tasks.Wait(); err != nil {
		return nil, err
	}

	// Trivial handling for where we find exactly one device to talk to, you
	// might want to provide more flexible handling than this if multiple
	// devices are found.
	switch {
	case len(ip2Clients) == 1:
		return ip2Clients[0], nil
	case len(ip1Clients) == 1:
		return ip1Clients[0], nil
	case len(ppp1Clients) == 1:
		return ppp1Clients[0], nil
	default:
		return nil, errors.New("multiple or no services found")
	}
}

func UpnpTest() {

	// clients, _, _ := internetgateway2.NewWANIPConnection2Clients()
	// client := clients[0]

	client, err := PickRouterClient(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	externalIP, err := client.GetExternalIPAddress()
	if err != nil {
		panic(err)
	}
	fmt.Println("Our external IP address is: ", externalIP)

	host, port, protocol, _, _, _, _, _, err := client.GetGenericPortMappingEntry(1900)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("host: %s\nport: %d\nprotocol: %s\n", host, port, protocol)

	// listing, err := client.GetListOfPortMappings(0, 8000, "TCP", false, 20)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Listing: %s\n", listing)

}

// Conventional REST APIs for getting IPv4 & IPv6 addresses

type APIResponse struct {
	Query string `json:"query"`
}

func GetIPv4Address() string {
	resp, err := http.Get("http://ip-api.com/json/")

	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var apiResp APIResponse
	json.Unmarshal(body, &apiResp)

	return apiResp.Query
}

func GetIPv6Address() string {
	resp, err := http.Get("http://api6.ipify.org/")

	if err != nil {
		resp, err = http.Get("http://v6.ident.me/") // fallback
		if err != nil {
			panic(err)
		}
	}

	body, _ := ioutil.ReadAll(resp.Body)

	return (string)(body)

}
