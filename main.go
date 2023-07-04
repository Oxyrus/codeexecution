package main

import (
	"bytes"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/execute", executeHandler)
	http.ListenAndServe(":8000", nil)
}

func executeHandler(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	code := buf.String()

	tmpFile, err := createTempFile(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	output, err := executeGoCode(tmpFile.Name())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(output))
}

func createTempFile(code string) (*os.File, error) {
	tmpFile, err := os.CreateTemp("", "go_code_*.go")
	if err != nil {
		return nil, err
	}

	defer tmpFile.Close()

	_, err = tmpFile.WriteString(code)
	if err != nil {
		return nil, err
	}

	return tmpFile, nil
}

func executeGoCode(filename string) (string, error) {
	cmd := exec.Command("go", "run", filename)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
