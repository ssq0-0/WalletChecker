package utils

import (
	"fmt"
	"regexp"
)

func ParseProxy(proxy string) (string, error) {
	userHost := regexp.MustCompile(`^([a-zA-Z0-9]+:[a-zA-Z0-9]+)@([0-9]+\.[0-9]+\.[0-9]+\.[0-9]+:[0-9]+)$`)
	hostUser := regexp.MustCompile(`^([0-9]+\.[0-9]+\.[0-9]+\.[0-9]+:[0-9]+)@([a-zA-Z0-9]+:[a-zA-Z0-9]+)$`)

	if matches := userHost.FindStringSubmatch(proxy); matches != nil {
		return fmt.Sprintf("http://%s", proxy), nil
	}

	if matches := hostUser.FindStringSubmatch(proxy); matches != nil {
		ipPort := matches[1]
		userPass := matches[2]

		return fmt.Sprintf("http://%s@%s", ipPort, userPass), nil
	}

	return "", fmt.Errorf("неправильный формат прокси: %s", proxy)
}
