package api

import (
	"fmt"
	"os/exec"
)

func Algorithm(algorithm string) {
	fmt.Println(algorithm)
	pyName := fmt.Sprintf("./lib/%s.py", algorithm)
	fmt.Println(pyName)

	// cmd := exec.Command("python", pyName, "arg1", "arg2")
	cmd := exec.Command("python", pyName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
