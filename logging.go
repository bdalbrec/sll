package sll

//LogInfo is used for logging informational messages. It takes a message string and prints it to the log file.
// It returns an error.
func LogInfo(msg string) error {

	t := getTime()

	f, err := prepFile()
	if err != nil {
		return err
	}

	if _, err := f.Write([]byte(t + ";" + "INFO; " + msg + "\r\n")); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return err
}

//LogError is used for logging errors that occur. It takes a message string and an error and prints them to the log file.
// It returns an error.
func LogError(msg string, e error) error {

	t := getTime()

	f, err := prepFile()
	if err != nil {
		return err
	}

	if _, err := f.Write([]byte(t + ";" + "ERROR; " + msg + "; " + e.Error() + "\r\n")); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return err

}
