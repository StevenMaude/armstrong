// Based on https://www.kluenter.de/garmin-ephemeris-files-and-linux/ and
// EPO_Downloader.rb in https://github.com/scrapper/postrunner (GPLv2)
// and information in this Garmin forum post:
// https://forums.garmin.com/outdoor-recreation/outdoor-recreation-archive/f/fenix-5-series/207977/epo-expired/1166435

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
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"flag"
)

// Garmin forum post:
// "…each EPO SET is 2304 bytes"
const epoLength = 2304

// retrieveDataEPO makes a HTTP request to get data and returns the body as []byte if successful.
func retrieveDataEPO() ([]byte, error) {
	url := "https://epodownload.mediatek.com/EPO.DAT"

	c := &http.Client{
		Timeout: 20 * time.Second,
	}
	resp, err := c.Get(url)
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

// checkDataLengthEPO checks the EPO data length; if not as expected, returns an error.
func checkDataLengthEPO(data []byte) error {
	dataLength := len(data)
	if dataLength != 120*epoLength {
		return fmt.Errorf("EPO data has unexpected length: %v", dataLength)
	}
	return nil
}

// trimEPOData trims the MediaTek data sufficiently for a Garmin watch.
func trimEPOData(data []byte) []byte {
	// Garmin forum post:
	// "…even with such a clean file, the Garmin watches use only a max of one-digit
	// days, ie 9 days of data…"
	const days = 9
	// Garmin forum post:
	// "…each EPO SET … has 6 hours of satellite locations…"
	const epoSetsCount = 4 * days
	return data[:(epoSetsCount * epoLength)]
}

// downloadFileEPO retrieves EPO data, checks it, cleans it and writes it to disk.
func downloadFileEPO() {
	fmt.Println("Retrieving EPO data from Mediatek's servers...")
	rawEPOData, err := retrieveDataEPO()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Processing EPO.BIN...")
	err = checkDataLengthEPO(rawEPOData)
	if err != nil {
		log.Fatal(err)
	}

	outData := trimEPOData(rawEPOData)

	err = ioutil.WriteFile("EPO.BIN", outData, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done! EPO.BIN saved.")
}


// retrieveData makes a HTTP request to get Garmin EPO data and returns the body as []byte if successful.
func retrieveDataCPE() ([]byte, error) {
	url := "https://api.gcs.garmin.com/ephemeris/cpe/sony?coverage=WEEKS_1"

	c := &http.Client{
		Timeout: 20 * time.Second,
	}
	resp, err := c.Get(url)
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

// checkDataLengthCPE errors out of CPE data is empty
// - CPE internal data format is not yet known
func checkDataLengthCPE(data []byte) error {
	dataLength := len(data)
	if dataLength == 0 {
		return fmt.Errorf("CPE data has unexpected length of zero")
	}
	return nil
}

// Retrieves CPE data, checks it and writes it to disk.
func downloadFileCPE() {
	fmt.Println("Retrieving CPE data from Garmin's servers...\n")
	rawCPEData, err := retrieveDataCPE()
	if err != nil {
		log.Fatal(err)
	}

	err = checkDataLengthCPE(rawCPEData)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("CPE.BIN", rawCPEData, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done! CPE.BIN saved.\n")
}


func main() {
	var modeCPE bool
	flag.BoolVar(&modeCPE, "cpe", false, "Download CPE format ephemeris data (instead of EPO)")
	flag.Parse()

	if modeCPE == true {
		downloadFileCPE()
	} else {
		downloadFileEPO()
	}
}
