package odeskidb

import "os"

func isDatabaseInitalized() bool {

	if databasePath == "" {
		return false
	}

	return true
}

//WARNING: this function is extremely dangerous
func overwriteDatabase(newData []byte) error {
	if err := os.WriteFile(databasePath+"odeski.db", newData, 0666); err != nil {
		return err
	}

	return nil
}
