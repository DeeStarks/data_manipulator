package main

import (
	"fmt"
	"os"
	"flag"

	"github.com/go-gota/gota/dataframe"
)

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func main() {
	var binderPath string
	var nSplits int
	const (
		defaultPath = ""
		pathUsage = "path to file, prefixed with \"@\""
		defaultNSplits = 100
		nSplitsUsage = "number of rows to split file by"
	)
	
	flag.StringVar(&binderPath, "filepath", defaultPath, pathUsage)
	flag.StringVar(&binderPath, "f", defaultPath, pathUsage+" (shorthand)")
	flag.IntVar(&nSplits ,"splits", defaultNSplits, nSplitsUsage)
	flag.IntVar(&nSplits ,"s", defaultNSplits, nSplitsUsage+" (shorthand)")
	flag.Parse()

	if binderPath == "" {
		fmt.Println("Please provide a filepath")
		return
	}

	binderPath = binderPath[1:]
	fmt.Println("==========================================================")
	fmt.Println("Reading data from", binderPath)
	fmt.Println("Number of splits:", nSplits)
	fmt.Println("==========================================================")

	df, err := os.Open(binderPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer df.Close()

	binderDF := dataframe.ReadCSV(df)

	var totalPages int = binderDF.Nrow() / nSplits
	var currentPage int
	for i := 0; i < binderDF.Nrow(); i += nSplits {
		right := i + nSplits - 1
		if currentPage == totalPages {
			right = binderDF.Nrow() - 1
		}
		newDF := binderDF.Subset(makeRange(i, right))
		// Check if directory "chunks" exists and create it if not
		if _, err := os.Stat("chunks"); os.IsNotExist(err) {
			os.Mkdir("chunks", 0755)
		}
		f, err := os.Create(fmt.Sprintf("chunks/file%d.csv", currentPage+1))
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		newDF.WriteCSV(f)
		currentPage++
	}
	fmt.Printf("Total number of files: %d\n", currentPage)
	fmt.Println("==========================================================")
}
