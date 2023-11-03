package main

import (
	"flag"
	"os"
)

var (
	Name string
	Comment string
	Exec string
	Icon string
	Categories string
	Output string
)
func main() {
	
	flag.StringVar(&Name, "n", "", "Program name")
	flag.StringVar(&Comment, "d", "", "Program description")
	flag.StringVar(&Exec, "e", "", "Program executable")
	flag.StringVar(&Icon, "i", "", "Program icon")
	flag.StringVar(&Categories, "c", "", "Program category")
	flag.StringVar(&Output, "o", "", "Output")

	flag.Parse()

	if len(Name) == 0 && len(Comment) == 0 && len(Exec) == 0 && len(Icon) == 0 && len(Categories) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	} else {
		file, err := os.Create(Output)
		if err != nil {
			os.Exit(1)
		}
		defer file.Close()
		file.WriteString("[Desktop Entry]\n")
		file.WriteString("Name=" + Name + "\n")
		file.WriteString("Comment=" + Comment + "\n")
		file.WriteString("Exec=" + Exec + "\n")
		file.WriteString("Icon=" + Icon + "\n")
		file.WriteString("Categories=" + Categories + "\n")
	}
}
