package producers

import (
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
)

// TextProducer implements the runtime.ProducerFunc interface. It creates a plain text producer that will write to HTTP
// responses.
func TextProducer() runtime.ProducerFunc {
	return func(writer io.Writer, i interface{}) error {

		// Type assert the io.ReadCloser.
		readCloser := i.(io.ReadCloser)

		// Read everything.
		data, err := ioutil.ReadAll(readCloser)
		if err != nil {
			return err
		}

		// Write the data to the writer.
		if _, err = writer.Write(data); err != nil {
			return err
		}

		// Close the io.ReadCloser.
		return readCloser.Close()
	}
}
