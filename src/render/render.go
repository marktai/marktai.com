package render

import (
	"bytes"
	"log"
	"os/exec"
	"strconv"
)

func Image(url, filePath string, height, width int) (string, string, error) {
	log.Println(url, filePath)
	cmd := exec.Command("/usr/local/bin/wkhtmltoimage", "--height", strconv.Itoa(height), "--width", strconv.Itoa(width), url, filePath)
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	err := cmd.Run()

	return stdOut.String(), stdErr.String(), err
}
