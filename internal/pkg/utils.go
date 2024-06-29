package pkg

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
)

func ProcessAddress(address *string) (string, error) {
	processedAddress := ""
	nets := strings.Split(*address, ".")
	if len(nets) != 4 {
		return "", errors.New("invalid address")
	}
	for idx, net := range nets {
		i, err := strconv.Atoi(net)
		if err != nil || i < 0 || i > 255 {
			return "", errors.New("invalid address")
		}
		processedAddress += strings.TrimSpace(strconv.Itoa(i))
		if idx < len(nets)-1 {
			processedAddress += "."
		}
	}
	return processedAddress, nil
}

func UpdateHostname(newHostName *string) error {

	cmd := exec.Command("hostnamectl", "set-hostname", *newHostName)
	if err := cmd.Run(); err != nil {
		if ok := err.(*exec.ExitError); ok != nil {
			return err
		}
	}
	return nil
}
