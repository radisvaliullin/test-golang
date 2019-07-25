package main

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

// Need all possible ip in string: example 127001 -> [127.0.0.1, 12.70.0.1]
func main() {

	ipsIn := []string{
		"127001",
		"1270011",
		"1200011",
		"1250011",
	}

	for _, ip := range ipsIn {
		out, err := findIPs(ip)
		if err != nil {
			return
		}
		log.Printf("for ip %v possible ips - %v", ip, out)
	}
}

func findIPs(ip string) ([]string, error) {

	out := []string{}

	return findIPRec(ip, 0, []int{0, 0, 0, 0}, out)
}

func findIPRec(ip string, n int, ipOcts []int, out []string) ([]string, error) {

	if n == 3 {
		oct := ip
		if len(oct) > 3 {
			return out, nil
		}
		if len(oct) > 1 && strings.HasPrefix(oct, "0") {
			return out, errors.New("oct is not int, start with 0")
		}
		octInt, err := strconv.Atoi(oct)
		if err != nil {
			return out, errors.New("oct is not int")
		}
		if octInt > 255 {
			return out, errors.New("oct is more then 255")
		}
		ipOcts[3] = octInt

		ipOctsStr := []string{"", "", "", ""}
		for i, v := range ipOcts {
			ipOctsStr[i] = strconv.Itoa(v)
		}

		out = append(out, strings.Join(ipOctsStr, "."))
		return out, nil
	}

	for i := 3; i >= 1; i-- {
		if len(ip) < i {
			continue
		}
		oct := ip[:i]
		tail := ip[i:]
		if len(oct) > 1 && strings.HasPrefix(oct, "0") {
			continue
		}
		octInt, err := strconv.Atoi(oct)
		if err != nil {
			return out, errors.New("oct is not int")
		}
		if octInt > 255 {
			continue
		}

		ipOcts[n] = octInt
		out, err = findIPRec(tail, n+1, ipOcts, out)
		if err != nil {
			continue
		}
	}

	return out, nil
}
