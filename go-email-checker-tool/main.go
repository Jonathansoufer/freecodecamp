package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
	"log"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")

	for scanner.Scan(){
		domain := scanner.Text()
		checkDomain(domain)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input: %v\n",err)
	}
}

func checkDomain(domain string){
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string 

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	spfRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range spfRecords {
		if strings.Contains(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.Contains(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
		}
	}

	fmt.Printf("%s, %t, %t, %s, %t, %s\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
