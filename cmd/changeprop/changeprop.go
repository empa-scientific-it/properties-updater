package main

import (
	dao "empa/basi/properties-updater/pkg/dao"
	ioutils "empa/basi/properties-updater/pkg/io"
	"fmt"
	"github.com/magiconair/properties"
	"os"
)

func UpdateFile(propertyFile *os.File, kv dao.KeyValue) {
	p := properties.MustLoadFile(propertyFile.Name(), properties.UTF8)
	prev, ok, err := p.Set(kv.Key, kv.Value)
	if ok {
		fmt.Println("Updating property file: " + propertyFile.Name())
		fmt.Println("Setting property: " + kv.Key + " with value: " + prev + " to value: " + kv.Value)
		fmt.Println("The file is now:")
		fmt.Println(p)
		n, err := p.Write(propertyFile, properties.UTF8)
		if err == nil {
			fmt.Printf("Wrote %d bytes\n", n)
		} else {
			ioutils.HandleError(err)
		}
	} else {
		ioutils.HandleError(err)
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
	pf, err := os.OpenFile(propertyFile, os.O_WRONLY, 0600)
	if err == nil {
		UpdateFile(pf, dao.KeyValue{Key: propertyName, Value: propertyValue})
		pf.Sync()

	} else {
		ioutils.HandleError(err)

	}
	defer pf.Close()

}
