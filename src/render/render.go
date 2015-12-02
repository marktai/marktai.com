package render

import (
	"bytes"
	"log"
	"os/exec"
)

func Image(url, filePath string) (string, string, error) {
	log.Println(url, filePath)
	cmd := exec.Command("/usr/local/bin/wkhtmltoimage", url, filePath)
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	err := cmd.Run()

	return stdOut.String(), stdErr.String(), err
}
