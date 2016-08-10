// Based on https://www.kluenter.de/garmin-ephemeris-files-and-linux/ and
// EPO_Downloader.rb in https://github.com/scrapper/postrunner (GPLv2)

// = EPO_Downloader.rb -- PostRunner - Manage the data from your Garmin sport devices.
//
// Copyright (c) 2015 by Chris Schlaeger <cs@taskjuggler.org>
//
// armstrong.go
//
// Copyright (c) 2016 by Steven Maude
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of version 2 of the GNU General Public License as
// published by the Free Software Foundation.

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// retrieveData makes a HTTP request to get Garmin EPO data and returns the body as []byte if successful.
func retrieveData() ([]byte, error) {
	url := "https://omt.garmin.com/Rce/ProtobufApi/EphemerisService/GetEphemerisData"
	// Data from https://www.kluenter.de/garmin-ephemeris-files-and-linux/
	data := []byte("\n-\n\aexpress\u0012\u0005de_DE\u001A\aWindows\"" +
		"\u0012601 Service Pack 1\u0012\n\b\x8C\xB4\x93\xB8" +
		"\u000E\u0012\u0000\u0018\u0000\u0018\u001C\"\u0000")

	resp, err := http.Post(url, "application/octet-stream", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// checkDataLength checks the EPO data length; if not as expected, returns an error.
func checkDataLength(data []byte) error {
	dataLength := len(data)
	// Each EPO data set is 2307 bytes long, with the first three bytes to be removed.
	if dataLength != 28*2307 {
		return fmt.Errorf("EPO data has unexpected length: %v", dataLength)
	}
	return nil
}

// cleanEPO removes the first three bytes from each block of 2307 bytes in EPO data,
// and returns a cleaned []byte.
func cleanEPO(rawEPOData []byte) []byte {
	var outData []byte
	for i := 0; i <= 27; i++ {
		offset := i * 2307
		outData = append(outData, rawEPOData[offset+3:offset+2307]...)
	}
	return outData
}

// main retrieves EPO data, checks it, cleans it and writes it to disk.
func main() {
	fmt.Println("Retrieving data from Garmin's servers...")
	rawEPOData, err := retrieveData()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Processing EPO.BIN...")
	err = checkDataLength(rawEPOData)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("raw", rawEPOData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	outData := cleanEPO(rawEPOData)

	err = ioutil.WriteFile("EPO.BIN", outData, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done! EPO.BIN saved.")
}
