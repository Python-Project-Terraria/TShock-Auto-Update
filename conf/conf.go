package conf

import "os"

func Parce(file string) error {
	_, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	return nil
}
