# DNS Exfiltration Script

🚨🚨🚨🚨  BEFORE USE CHECK THE TXT FILE'S PATH AND CHANGE ACCORDING WITH YOURS 🚨🚨🚨🚨

This script allows for the exfiltration of data using DNS queries. The script reads a secrets.txt file, divides the file into strings, and appends each string to a DNS query. The DNS server listens for the query and always responds with a valid IP address, and saves the part of the domain in the query that contains the string into a file named exfil.txt.

Requirements

Go 1.15 or later
The github.com/miekg/dns package, which can be installed using the following command:

> go get -u github.com/miekg/dns

Usage

Start the DNS server by running the script on one machine.
On another machine, run the client script that sends the DNS queries containing the secret data.
The server will listen for incoming queries, extract the data from the domain names, and write it to the exfil.txt file.
Note: Make sure to replace the IP address in the client script with the IP address of the machine running the DNS server.

Limitations

The exfiltrated data is limited by the length of the domain name in the DNS query, which is typically limited to 255 characters.
The script does not provide encryption or any other security measures, so the transmitted data can be easily intercepted if the network is not secure.
Conclusion

This script demonstrates a basic example of how data can be exfiltrated using DNS queries. However, it is not recommended to use this method in real-world scenarios as it lacks security measures and can easily be detected.
