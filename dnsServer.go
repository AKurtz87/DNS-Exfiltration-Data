package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/miekg/dns"
)

func removeChars(s string) string {
	// Remove first 4 and last 4 elements from the byte array
	res := s[ len(s)-11 : len(s)-1  ]

	return string(res)
}

func writeToFile(str string) error {
	f, err := os.OpenFile("exfil.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Failed to open file: %v", err)
	}
	defer f.Close()

	// WRITE ALL STRINGS ON THE SAME LINE
	//if _, err := f.WriteString(str); err != nil {
	// WRITE ALL STRINGS ON A NEW LINE
	if _, err := f.WriteString(str + "\n"); err != nil {
		return fmt.Errorf("Failed to write to file: %v", err)
	}

	return nil
}

func handleRequest(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)

	file, err := os.Open("/Users/ak/Desktop/CHAT_GPT/DNS_DATA_SHARE/server/host.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) != 2 {
			fmt.Println("Error: invalid input format")
			continue
		}

		domain := fields[0]
		ip := fields[1]
		ipAddr := net.ParseIP(ip)
		if ipAddr == nil {
			fmt.Printf("Error: invalid IP address for domain %s\n", domain)
			continue
		}

		fmt.Println(r.Question[0].Name)

		//if r.Question[0].Name == domain+"." {
			writeToFile(removeChars(r.Question[0].Name))
			msg.Answer = append(msg.Answer, &dns.A{
				Hdr: dns.RR_Header{
					Name:   r.Question[0].Name,
					Rrtype: dns.TypeA,
					Class:  dns.ClassINET,
					Ttl:    60,
				},
				A: ipAddr,
			})
		//}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	w.WriteMsg(&msg)
}

func main() {
	dns.HandleFunc(".", handleRequest)
	

	err := dns.ListenAndServe("127.0.0.1:53", "udp", nil)
	if err != nil {
		fmt.Println(err)
	}

	
}

