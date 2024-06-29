package service

import (
	"context"
	"fmt"
	"os"
	"strings"

	pb "github.com/ladiesman2127/netconfig/api/gen/go"
	common "github.com/ladiesman2127/netconfig/internal/pkg"
)

type service struct {
	pb.UnimplementedNetConfigServer
}

func New() *service {
	return &service{}
}

func (s *service) UpdateHostname(ctx context.Context, in *pb.UpdateHostnameRequest) (*pb.Empty, error) {
	fmt.Println("UpdateHost")
	err := common.UpdateHostname(&in.NewHostname)
	if err != nil {
		return &pb.Empty{}, err
	}
	return &pb.Empty{}, nil
}

func (s *service) AddDns(ctx context.Context, in *pb.AddDNSRequest) (*pb.Empty, error) {
	fmt.Println("AddDns", in.NewDns)
	parsedAddress, err := common.ProcessAddress(&in.NewDns)
	if err == nil {
		return &pb.Empty{}, err
	}

	content, err := os.ReadFile("/etc/resolv.conf")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	lines = append(lines, parsedAddress+"\n")

	err = os.WriteFile("/etc/resolv.conf", []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (s *service) DeleteDns(ctx context.Context, in *pb.DeleteDNSRequst) (*pb.Empty, error) {

	parsedAddress, err := common.ProcessAddress(&in.DnsToRemove)
	if err == nil {
		return &pb.Empty{}, err
	}

	content, err := os.ReadFile("/etc/resolv.conf")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	var newLines []string

	for _, line := range lines {
		if !strings.Contains(line, parsedAddress) {
			newLines = append(newLines, line)
		}
	}

	err = os.WriteFile("/etc/resolv.conf", []byte(strings.Join(newLines, "\n")), 0644)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (s *service) GetAllDns(ctx context.Context, in *pb.Empty) (*pb.DnsListResponse, error) {
	addresses, err := getAllDns()
	if err != nil {
		return nil, nil
	}
	return &pb.DnsListResponse{Addresses: addresses}, nil
}

func getAllDns() ([]string, error) {
	content, err := os.ReadFile("/etc/resolv.conf")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	var addresses []string

	for _, line := range lines {
		if strings.HasPrefix(line, "nameserver") {
			parts := strings.Fields(line)
			if len(parts) > 1 {
				addresses = append(addresses, parts[1])
			}
		}
	}
	return addresses, nil
}
