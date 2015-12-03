package render

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func Image(url, filePath string, height, width int) (string, string, error) {
	cmd := exec.Command("/usr/local/bin/wkhtmltoimage", "--height", strconv.Itoa(height), "--width", strconv.Itoa(width), url, filePath)
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	start := time.Now()
	err := cmd.Run()

	elapsed := time.Now().Sub(start)

	f, err := os.OpenFile("/home/ubuntu/repos/marktai.com/upload/timings.txt", os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return stdOut.String(), stdErr.String(), err
	}

	defer f.Close()

	line := url + " took " + elapsed.String()

	_, _ = f.WriteString(line)
	return stdOut.String(), stdErr.String(), err
}
