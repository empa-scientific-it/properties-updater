package main

import (
	"testing"
	ioutils "empa/basi/properties-updater/pkg/io"
	dao "empa/basi/properties-updater/pkg/dao"
	"os"
	"github.com/magiconair/properties"

)


func writeKeyValue(file *os.File, kv dao.KeyValue) {
	p := properties.MustLoadFile(file.Name(), properties.UTF8)
	p.Set(kv.Key, kv.Value)
	p.Write(file, properties.UTF8)
}



func TestSetProp(t *testing.T) {
	cases := []struct {
		orig dao.KeyValue
		new  dao.KeyValue
	}{
		{dao.KeyValue{Key: "a", Value: "1"}, dao.KeyValue{Key: "a", Value: "2"}},
	}
	for _, c := range cases {
		// Prepare the file with original value
		tf, err := ioutils.TempFile()
		ioutils.HandleError(err)
		writeKeyValue(tf, c.orig)
		tf.Close()
		// Update the file with new value
		tfNew, err  := os.OpenFile(tf.Name(), os.O_WRONLY, 0600)
		ioutils.HandleError(err)
		UpdateFile(tfNew, c.new, Replace)
		tfNew.Close()
		p := properties.MustLoadFile(tfNew.Name(), properties.UTF8)
		newVal, ok := p.Get(c.orig.Key)
		defer os.Remove(tf.Name())
		if ok{
			if newVal != c.new.Value {
				t.Errorf("Expected %s, got %s", c.new.Value, newVal)
			}
		}
	}
}
