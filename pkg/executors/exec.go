package executors

import (
	"bytes"
	"os/exec"
)

// CreateDirReq :
type CreateDirReq struct {
	Name  string
	Force bool
}

// CreateDirs : create directories
func CreateDirs(dirs ...CreateDirReq) error {
	var cmd *exec.Cmd
	for _, dir := range dirs {
		if dir.Force {
			cmd = exec.Command("mkdir", "-p", dir.Name)
		} else {
			cmd = exec.Command("mkdir", dir.Name)
		}

		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func GoInit(projectDir string) (string, error) {
	cmd := exec.Command("go", "mod", "init")
	cmd.Dir = projectDir
	var errb bytes.Buffer
	cmd.Stderr = &errb

	output, err := cmd.Output()
	if err != nil {
		return errb.String(), err
	}

	return string(output), nil
}

func GoImport(projectDir string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", `goimports -w -l $(find . -type f -name '*.go' -not -path "./vendor/*")`)
	cmd.Dir = projectDir
	var errb bytes.Buffer
	cmd.Stderr = &errb

	output, err := cmd.Output()
	if err != nil {
		return errb.String(), err
	}

	return string(output), nil
}
