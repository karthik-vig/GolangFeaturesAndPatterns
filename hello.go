package main

import (
	"bufio"
	"example/user/helloworld/daisychain"
	"example/user/helloworld/generators"
	"example/user/helloworld/generics"
	"example/user/helloworld/multiplexing"
	"example/user/helloworld/quitindicator"
	"example/user/helloworld/syncStepGoroutines"
	"example/user/helloworld/timeout"
	"fmt"
	"os"
)

type userDefinedFuncType func(x, y int) int

func someFuncOfUserDefinedFuncType(x, y int) int {
	return x + y
}

func higherOrderFunc(callback userDefinedFuncType) int {
	return callback(2, 3)
}

func exampleVariadicFunction(vals ...int) {
	for idx, val := range vals {
		fmt.Printf("The value at index: %d is %d\n", idx, val)
	}
}

func derterminedChannelDirection(getValueCh <-chan int, putValueCh chan<- int) {
	putValueCh <- 1 + <-getValueCh
}

func experimentWithReaderAndWriter() {
	//create a reader first with a location to read from
	//in this case it is the stadard console io from the os
	//there is also a variant of NewReader() with size
	//called NewReaderSize()
	reader := bufio.NewReader(os.Stdin)
	//now we can use reader with methods it implements
	//fmt.Println(reader.Buffered())
	var p []byte = make([]byte, 10)
	fmt.Println("Enter the content to be read by .Read()")
	numberOfChars, err := reader.Read(p)
	if err == nil {
		fmt.Println("From Read() method: ", p[:numberOfChars])
	}
	fmt.Println("Enter the content to be read by .ReadString()")
	someString, err2 := reader.ReadString('\n')
	if err2 == nil {
		fmt.Println("From ReadString() method: ", someString)
	}
	fmt.Println("Enter the content to be read by .ReadLine()")
	for {
		someStringByte, isPrefix, err3 := reader.ReadLine()
		if err3 != nil {
			break
		}
		if !isPrefix {
			fmt.Println("From ReadLine() method: (byte variant)", someStringByte)
			fmt.Println("From ReadLine() method: (string variant)", string(someStringByte))
			break
		}
		fmt.Println("From ReadLine() method: ", string(someStringByte))
	}
	// use a scanner to get input
	fmt.Println("Enter the content to be read by Scanner (type \"stop\" to quit scanner read)")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())  // string
		fmt.Println(scanner.Bytes()) // bytes
		if scanner.Text() == "stop" {
			break
		}
	}
	// custom writer
	// also has a variant called NewWriterSize()
	fmt.Println("Content Written to console by Writer")
	writer := bufio.NewWriter(os.Stdout)
	numberOfCharsWritten, charsWriteErr := writer.Write([]byte{97, 97})
	if charsWriteErr == nil {
		fmt.Println("Number of bytes to be written by .Write(): ", numberOfCharsWritten)
	}
	numberOfCharsWrittenFromString, stringWriteErr := writer.WriteString("\r\nhello world\r\n")
	if stringWriteErr == nil {
		fmt.Println("Number of bytes to be written by .WriteString()", numberOfCharsWrittenFromString)
	}
	// acutally writes to console the stuff in the
	// writer's buffer
	writer.Flush()

	// read and write content directly using os.Stdin and os.Stdout
	fmt.Println("Read Content using os.Stdin.Read(): ")
	var pp []byte = make([]byte, 10)
	os.Stdin.Read(pp)
	fmt.Println("Write content using os.Stdout.Write(): ")
	os.Stdout.Write(pp)
	os.Stdout.Write([]byte{97, 98})
}

func writeFileOperations() {
	// create a reader
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter data to be written to file")
	stringData, consoleReadStringErr := reader.ReadString('\n')
	if consoleReadStringErr != nil {
		fmt.Println("Error while reading data from console")
		return
	}
	// write byte data to file
	// convert string to a byte slice first
	fileWriteErr := os.WriteFile("D:\\Desktop\\test.txt", []byte(stringData), 0777)
	if fileWriteErr != nil {
		fmt.Println("Error occured while writing to file")
		return
	}
	fmt.Println("Writing to file was Successfull!!!")
}

func readFileOperations() {
	fmt.Println("Attempting to read from file")
	fileData, fileReadErr := os.ReadFile("D:\\Desktop\\test.txt")
	if fileReadErr != nil {
		fmt.Printf("File read operation faced error:\n%s.\nExiting...\n", fileReadErr.Error())
		return
	}
	fmt.Println("The data in the file is: ", string(fileData))
}

func appendWriteToFile() {
	file, fileErr := os.OpenFile("D:\\Desktop\\test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if fileErr != nil {
		fmt.Println("Error occurred while attempting to open the file")
		return
	}
	defer file.Close()
	someMsg := "So, this part will be appended to the file..."
	//var fileData []byte
	//fileData = append(fileData, []byte(someMsg)...)
	numberOfBytesWritten, fileWriteErr := file.Write([]byte(someMsg))
	if fileWriteErr != nil {
		fmt.Println("Error occurred while appending data to the file: \n", fileWriteErr.Error())
		return
	}
	fmt.Println("The append-write was successful. Number of bytes written: ", numberOfBytesWritten)
}

func main() {
	fmt.Println("Main function started!!!")
	fmt.Println("Select a Number for a Demo:")
	fmt.Println("1) Concurrency Pattern - Generator")
	fmt.Println("2) Concurrency Pattern - Multiplexing/fanin")
	fmt.Println("3)  Concurrency Pattern - Sync Goroutines")
	fmt.Println("4) Concurrency Pattern - Timeout (channel has to reply within given time)")
	fmt.Println("5) Concurrency Pattern - Timeout (channel finishes all comuunication within given time)")
	fmt.Println("6) Concurrency Pattern - Two way quit communication with Goroutine")
	fmt.Println("7) Concurrency Pattern - Daisy Chain")
	fmt.Println("8) Higher order functions")
	fmt.Println("9) Variadic functions")
	fmt.Println("10) Determine the channel type as taker or giver of value")
	fmt.Println("11) I/O Operatoins")
	fmt.Println("12) Write File Operations")
	fmt.Println("13) Read File Operations")
	fmt.Println("14) Append write to File Operations")
	fmt.Println("15) Test Generics")
	fmt.Println("Enter any other inputs to quit")
	reader := bufio.NewReader(os.Stdin)
	input, inputErr := reader.ReadString('\n')

	if inputErr != nil {
		fmt.Println("Error occurred tyring to read from console. Exiting...")
		return
	}

	// remove \r and \n from the end of the string
	input = input[:len(input)-2]

	switch input {
	case "1":
		//generators example
		receiveOnlyChannel := generators.GeneratorEx(0)
		for {
			val, ok := <-receiveOnlyChannel
			fmt.Println(val, ok)
			if !ok {
				break
			}
		}

	case "2":
		// multiplexing or fanin
		multCh1 := generators.GeneratorEx(100)
		multCh2 := generators.GeneratorEx(200)
		multCommonCh := multiplexing.Multiplex(multCh1, multCh2)
		for {
			val, ok := <-multCommonCh
			fmt.Println(val, ok)
			if !ok {
				break
			}
		}

	case "3":
		// sync step locked goroutines
		syncStepGoroutines.Synched()

	case "4":
		//timeouts
		timeout.TimeoutType1()

	case "5":
		//timeouts
		timeout.TimeoutType2()

	case "6":
		// example of using channel to inidcate end of goroutine
		// and perform a indication that it can quit safely, finally
		quitindicator.QuitIndicatorExample()

	case "7":
		// example of daisy-chaining channels between goroutines
		daisychain.DaisyChainExample()

	case "8":
		fmt.Println(higherOrderFunc(someFuncOfUserDefinedFuncType))

	case "9":
		//variadic function
		exampleVariadicFunction(1, 2, 3, 4)

	case "10":
		//make determistic channel
		getValueCh := make(chan int)
		putValueCh := make(chan int)
		go derterminedChannelDirection(getValueCh, putValueCh)
		getValueCh <- 0
		fmt.Println(<-putValueCh)
	case "11":
		// io operations
		fmt.Println("I/O Operations selected")
		experimentWithReaderAndWriter()
	case "12":
		// write file operations
		writeFileOperations()
	case "13":
		// read file operations
		readFileOperations()
	case "14":
		// append write data to file
		appendWriteToFile()
	case "15":
		// test generics
		generics.TestGenerics()
	}
}
