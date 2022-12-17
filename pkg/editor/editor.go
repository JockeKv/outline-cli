package editor

import (
	"fmt"
	"os"
	"os/exec"
	"outline/pkg/api"
)

func EditDoc(doc *api.Document) (*api.Document, error) {
	if doc == nil {
		return nil, fmt.Errorf("no document")
	}
	data := []byte(*doc.Text)
	file := *doc.Id
	filename, err := writeFile(file, data)
	if err != nil {
		return nil, err
	}

	edited, err := editFile(filename)
	if err != nil {
		return nil, err
	}

	if edited {
		data, err = os.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		*doc.Text = string(data)
		err = removeFile(file)
		if err != nil {
			return nil, err
		}
		return doc, nil
	}
	err = removeFile(file)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func editFile(file string) (bool, error) {
	initialStat, err := os.Stat(file)
	if err != nil {
		return false, err
	}

	editor := os.Getenv("EDITOR")
	cmd := exec.Command(editor, file)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return false, err
	}

	stat, err := os.Stat(file)
	if err != nil {
		return false, err
	}

	if stat.Size() != initialStat.Size() || stat.ModTime() != initialStat.ModTime() {
		fmt.Println("File edited")
		return true, nil
	}
	fmt.Println("File not edited")
	return false, nil
}

func writeFile(file string, data []byte) (string, error) {

	filename := fmt.Sprintf("/tmp/%s.md", file)

	err := os.WriteFile(filename, data, 0733)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func removeFile(file string) error {

	filename := fmt.Sprintf("/tmp/%s.md", file)

	err := os.Remove(filename)
	if err != nil {
		return err
	}
	return nil
}
