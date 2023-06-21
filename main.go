package main


import(
	"github.com/magiconair/properties"
	"fmt"
	"os"
)

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
	p.Set(propertyName, propertyValue)
	fmt.Println("Updating property file: " + propertyFile)
	fmt.Println("Setting property: " + propertyName + " to value: " + propertyValue)
	fmt.Println(p)
	p.Write(propertyFile, properties.UTF8)
}