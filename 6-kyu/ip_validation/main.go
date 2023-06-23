package cw

import (
	"strconv"
	"strings"
	"unicode"
)

func Is_valid_ip(ip string) bool {
	octets := strings.Split(ip, ".")
	if len(octets) < 4 {
		return false
	}

	for _, octet := range octets {
		if octet[0] == '0' && len(octet) > 1 {
			return false
		}
		for _, c := range octet {
			if !unicode.IsDigit(c) {
				return false
			}
		}
		octetInt, _ := strconv.Atoi(octet)
		if octetInt > 255 {
			return false
		}
	}

	return true
}
