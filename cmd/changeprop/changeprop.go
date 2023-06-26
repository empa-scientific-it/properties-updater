package main

import (
	dao "empa/basi/properties-updater/pkg/dao"
	ioutils "empa/basi/properties-updater/pkg/io"
	"errors"
	"fmt"
	"github.com/magiconair/properties"
	"os"
)

type Mode struct {
	slug string
}

var (
	Replace   = Mode{"replace"}
	Create    = Mode{"create"}
	Uncomment = Mode{"uncomment"}
	Comment   = Mode{"comment"}
	Unknown   = Mode{"unknown"}
)

func FromString(s string) (Mode, error) {
	switch s {
	case Replace.slug:
		return Replace, nil
	case Create.slug:
		return Create, nil
	}
	return Unknown, errors.New("unknown mode: " + s)
}

func UpdateFile(propertyFile *os.File, kv dao.KeyValue, mode Mode) {
	p := properties.MustLoadFile(propertyFile.Name(), properties.UTF8)
	prev, ok, err := p.Set(kv.Key, kv.Value)

	fmt.Println(ok)
	if (ok && mode == Replace) || (!ok && mode == Create) {
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
	if len(argsWithoutProg) != 4 {
		fmt.Println("Usage: properties-updater <property-file> <mode> <property-name> <property-value>")
		return
	}
	fmt.Println(argsWithoutProg)

	propertyFile := argsWithoutProg[0]
	mode, err := FromString(argsWithoutProg[1])
	ioutils.HandleError(err)
	propertyName := argsWithoutProg[2]
	propertyValue := argsWithoutProg[3]
	pf, err := os.OpenFile(propertyFile, os.O_WRONLY, 0600)
	if err == nil {
		UpdateFile(pf, dao.KeyValue{Key: propertyName, Value: propertyValue}, mode)
		pf.Sync()

	} else {
		ioutils.HandleError(err)

	}
	defer pf.Close()

}
