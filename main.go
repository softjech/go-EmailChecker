package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Domain -- hasMX -- hasSPF -- spfRecord -- hasDMARC -- dmarcRecord")
	for scanner.Scan() {
		domain := scanner.Text()
		if !checkDomain(domain) {
			fmt.Println("Invalid domain")
		}

		if err := scanner.Err(); err != nil {
			log.Fatal("\nError: could not read from the input: %v\n", err)
		}
	}
}

func checkDomain(domain string) bool {

	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: could not find MX records for %s: %v\n", domain, err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecord, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: could not find TXT records for %s: %v\n", domain, err)
	}
	for _, record := range txtRecord {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
		}
	}

	dRecord, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: could not find DMARC record for %s: %v\n", domain, err)
	}
	for _, record := range dRecord {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
		}
	}

	fmt.Printf("\n %s -- %t -- %t -- %s -- %t -- %s\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)

	return true
}
