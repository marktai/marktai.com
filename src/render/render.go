package render

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func PDF(url, filePath string) (string, string, error) {
	cmd := exec.Command("/usr/local/bin/wkhtmltopdf", url, filePath)
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

	line := url + "@pdf" + " took " + elapsed.String() + "\n"

	_, _ = f.WriteString(line)
	return stdOut.String(), stdErr.String(), err
}

func Image(url, filePath string, height, width int) (string, string, error) {
	heightString := strconv.Itoa(height)
	widthString := strconv.Itoa(width)
	cmd := exec.Command("/usr/local/bin/wkhtmltoimage", "--height", heightString, "--width", widthString, url, filePath)
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

	line := url + "@" + widthString + "x" + heightString + " took " + elapsed.String() + "\n"

	_, _ = f.WriteString(line)
	return stdOut.String(), stdErr.String(), err
}
