package main

import (
	"fmt"
	"github.com/magiconair/properties"
	"os"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 3 {
		fmt.Println("Usage: properties-updater <property-file> <property-name> <property-value>")
		return
	}
	propertyFile := argsWithoutProg[0]
	propertyName := argsWithoutProg[1]
	propertyValue := argsWithoutProg[2]
	p := properties.MustLoadFile(propertyFile, properties.UTF8)
	prev, ok, err := p.Set(propertyName, propertyValue)
	if ok {
		fmt.Println("Updating property file: " + propertyFile)
		fmt.Println("Setting property: " + propertyName  + " with value: " + prev + " to value: " + propertyValue)
		fmt.Println("The file is now:")
		fmt.Println(p)
		outFile, err := os.Create(propertyFile)
		handleError(err)
		defer outFile.Close()
		p.Write(outFile, properties.UTF8)

	}else{
		handleError(err)
	}

}
