package configmodify

import (
	"fmt"
	"os"
)

// IsFileExist : file exists or not
func IsFileExist(filepath string) error {
	st, err := os.Stat(filepath)
	if err != nil {
		return fmt.Errorf("%s doesn't exist", filepath)
	}
	if st.IsDir() {
		return fmt.Errorf("%s is not a file, but a directory", filepath)
	}
	return nil
}
