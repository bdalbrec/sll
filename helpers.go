package sll

import (
	"fmt"
	"os"
	"time"
)

// prepFile is used to set up the file for writing by the LogX functions. It returns a pointer to a file and an error
func prepFile() (*os.File, error) {
	fileName := "logs/dumbLog.log"

	err := manageFile(fileName)
	if err != nil {
		err = fmt.Errorf("Error in prepFile. Could not manage file - %v\n", err)
		return nil, err
	}

	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Errorf("Error in prepFile: Could not open or create log file. - %v\n", err)
		return nil, err
	}

	return f, nil

}

// getTime returns a formatted datetime string used in the log files
func getTime() string {
	return time.Now().Format("2006/01/02 15:04:05.999")
}

// getDate returns an formatted datetime string used for file renames
func getDate() string {
	return time.Now().Format("20060102150405") // using this as a timestamp for renaming files, so slashes donn't work
}

// manageFile determines when the log file has grown over a set size limit and renames it with the current
// datetime+filename.log allowing a new logfile to be created in its place while preserving the log history
func manageFile(fn string) error {
	fi, err := os.Stat(fn)
	if err != nil {
		err = fmt.Errorf("Error in manageFile: Could not retrieve file stat. - %v\n", err)
		return err
	}

	if fi.Size() > 15000 {
		err := os.Rename(fn, getDate()+fn)
		if err != nil {
			err = fmt.Errorf("Error in manageFile: Could not rename file. - %v\n", err)
			return err
		}
	}

	return err

}
