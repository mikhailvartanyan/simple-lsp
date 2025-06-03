package main

import (
	"bufio"
	"log"
	"os"
	"simple-lsp/rpc"
)

func main() {
	logger := getLogger("./log.txt")
	logger.Println("Logger initialised")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Splitter)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(logger, msg)
	}
}

func handleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	return log.New(logfile, "[simple-lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
