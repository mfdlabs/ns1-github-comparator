package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	ns1 "gopkg.in/ns1/ns1-go.v2/rest"
)

var (
	ns1url       = flag.String("ns1-url", "", "The URL for the NS1 API, warning: It has to be prepended with a slash or ns1 client will just strip out the last part of the URL")
	fetchTimeout = flag.Duration("fetch-timeout", 30*time.Second, "The timeout for the fetch operation in the background, this cannot be less than 30 seconds")
	githubAuthor = flag.String("github-author", "", "The author name for the commit message")
)

func main() {
	flag.Usage = func() {
		os.Stderr.WriteString("Usage of ns1-cli:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *ns1url == "" {
		flag.Usage()
		os.Exit(1)
	}

	// check if the fetch timeout is less than 30 seconds
	if *fetchTimeout < 30*time.Second {
		panic("fetch-timeout cannot be less than 30 seconds")
	}

	// verify if the ns1-to-zone-file binary is available at $PWD/lib/ns1-to-zone-file
	_, err := os.Stat("lib/ns1-to-zone-file")

	if err != nil {
		panic("ns1-to-zone-file binary not found, please download it at `https://github.com/mfdlabs-grid-development/ns1-to-zone-file/releases/download/v1.0.0/ns1-to-zone-file` and extract it to $PWD/lib/ns1-to-zone-file")
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client := ns1.NewClient(nil, ns1.SetEndpoint(*ns1url))

	// listen for a keyboard interrupt of the key 'F'
	// this will purge the cache and exit the program
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// The next step is to run the workflow loop in the background, making sure execution doesn't stop until SIGINT or SIGTERM is received.
	// but not blocking the main thread.
	go func() {
		for {
			// Time how long it takes to do the work
			start := time.Now()
			doWork(client)
			log.Printf("Work took %s\n", time.Since(start))
			log.Printf("Sleeping for %s\n", *fetchTimeout)
			time.Sleep(*fetchTimeout)
		}
	}()

	// We have a special case here, for SIGUSR1 or SIGUSR2, we want to do the same thing as SIGINT or SIGTERM, but we want to purge the cached dns files (not repo, the json files that hold state)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGUSR2)
	defer func() {
		s := <-sig

		switch s {
		case syscall.SIGUSR1:
			// Purge the cache in here
			log.Printf("Received %s, purging cache...\n", s)
			purgeCache()
			os.Exit(0)
		case syscall.SIGUSR2:
			// Purge the cache in here
			log.Printf("Received %s, purging cache...\n", s)
			purgeCache()
			os.Exit(0)
		default:
			log.Printf("Received signal: %s. Exiting...\n", s)
			os.Exit(0)
		}

	}()
}

func doWork(client *ns1.Client) {
	// this function will do the following:
	// 1. Fetch the list of zones from ns1
	// 2. Iterate over the zones and serialize them to json
	// 3. Write the json to the cache directory
	// 4. Run the ns1-to-zone-file binary to generate the dns files.

	// 1. Fetch the list of zones from ns1
	zones, _, err := client.Zones.List()
	if err != nil {
		log.Fatal(err)
	}

	var commitMessage string = "Update to NS1 zones."

	// 2. Iterate over the zones and serialize them to json
	for _, zone := range zones {
		// serialize the zone to json using encoding/json
		json, err := json.Marshal(zone)

		if err != nil {
			log.Fatal(err)
		}

		// write the json to the cache directory
		err = writeToCache(json, zone.Zone, &commitMessage)
		if err != nil {
			log.Fatal(err)
		}
	}

	// 3. Run the ns1-to-zone-file binary to generate the dns files.
	err = runNs1ToZoneFile()

	if err != nil {
		log.Fatal(err)
	}

	// 4. Push the dns files to the repo
	err = pushToRepo(commitMessage)

	if err != nil {
		log.Fatal(err)
	}
}

func pushToRepo(commitMessage string) error {
	// Push the dns files to the repo
	cmd := exec.Command("git", "add", "ns1")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		return err
	}

	author := *githubAuthor

	if author == "" {
		author = "NS1 Changes Detector <ns1-changes@ops.vmminfra.local>"
	}

	cmd = exec.Command("git", "commit", "--author", author, "-m", "\""+commitMessage+"\"")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()

	if err != nil {
		return err
	}

	cmd = exec.Command("git", "push")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()

	if err != nil {
		return err
	}

	return nil
}

func runNs1ToZoneFile() error {
	// Cache directory is $PWD/ns1_dns_json_cache
	// Run the ns1-to-zone-file binary to generate the dns files.
	cmd := exec.Command("lib/ns1-to-zone-file", "-zone-files-path", "ns1_dns_json_cache", "-dns-zone-files-output-path", "ns1")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func writeToCache(json []byte, zone string, commitMessage *string) error {
	// this function will write the json to the cache directory
	// the cache directory is $PWD/ns1_dns_json_cache

	// If the zone already exists in the output directory, we need to find differences between the two files and write them to the output directory.
	if _, err := os.Stat("ns1_dns_json_cache/" + zone + ".json"); err == nil {
		// We need to write this one to temp files while we compare the two files.
		// then delete it from temp files when finished.

		// 1. Write this json to a temp file
		tmpFileName := "tmp_" + zone + ".json"
		tmpFilePath := "ns1_dns_json_cache/" + tmpFileName
		err := ioutil.WriteFile(tmpFilePath, json, 0644)

		if err != nil {
			return err
		}

		// 2. Check old file against new file
		isDifferent := deepCompare(tmpFilePath, "ns1_dns_json_cache/"+zone+".json")

		// 3. If the files are different, append the commit message.
		if isDifferent {
			*commitMessage += fmt.Sprintf("\n~The Zone (%s) was modified since the last check.\n", zone)
		}

	}

	// create the cache directory if it doesn't exist
	err := os.MkdirAll("ns1_dns_json_cache", 0755)
	if err != nil {
		return err
	}

	// write the json to the cache directory
	err = ioutil.WriteFile(fmt.Sprintf("ns1_dns_json_cache/%s.json", zone), json, 0644)
	if err != nil {
		return err
	}

	return nil
}

func purgeCache() {
	// Cache directory is $PWD/ns1_dns_json_cache
	// Purge the cache in here
	err := os.RemoveAll("ns1_dns_json_cache")

	if err != nil {
		log.Printf("Error purging cache: %s\n", err)
	}
}

const chunkSize = 64000

func deepCompare(file1, file2 string) bool {
	// Check file size ...

	f1, err := os.Open(file1)
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	f2, err := os.Open(file2)
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	for {
		b1 := make([]byte, chunkSize)
		_, err1 := f1.Read(b1)

		b2 := make([]byte, chunkSize)
		_, err2 := f2.Read(b2)

		if err1 != nil || err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return true
			} else if err1 == io.EOF || err2 == io.EOF {
				return false
			} else {
				log.Fatal(err1, err2)
			}
		}

		if !bytes.Equal(b1, b2) {
			return false
		}
	}
}
