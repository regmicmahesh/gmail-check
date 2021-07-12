package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)


var wg sync.WaitGroup
var client *http.Client

var inputFile = flag.String("input", "", "File to read email address from.")
var outputFile = flag.String("output", "", "File to write valid email address to.")

func init() {
	flag.Parse()
}

func check(w *bufio.Writer, mail string) error {
	defer wg.Done()
	defer writer.Flush()
	resp, err := client.Get("https://mail.google.com/mail/gxlu?email=" + mail)
	if err != nil {
		return err
	}

	if len(resp.Cookies()) > 0 {
		fmt.Fprintln(w, mail)
	}else{
	}
	return nil

}

var reader *bufio.Reader
var writer *bufio.Writer

func main() {

	t := &http.Transport{
		MaxIdleConns: 1000,
		MaxIdleConnsPerHost: 1000,
		MaxConnsPerHost: 1000,
	}

	client = &http.Client{
		Transport: t,
	}

	if len(*inputFile) != 0 {
		file, _ := os.Open(*inputFile)
		defer file.Close()
		reader = bufio.NewReader(file)
	} else {
		reader = bufio.NewReader(os.Stdin)
	}

	if len(*outputFile) != 0 {
		file, err := os.OpenFile(*outputFile, os.O_WRONLY|os.O_CREATE, 777)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		writer = bufio.NewWriter(file)
	} else {
		writer = bufio.NewWriter(os.Stdout)
	}

	for {

		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}
		wg.Add(1)
		go check(writer, strings.TrimSpace(line))

	}

	wg.Wait()

}
