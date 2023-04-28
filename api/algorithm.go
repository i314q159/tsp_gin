package api

import (
	"fmt"
	"os/exec"
	"strings"
)

func Algorithm(algorithm string, cp []string) {
	fmt.Println(algorithm)
	pyName := fmt.Sprintf("./lib/%s.py", algorithm)

	args := strings.Replace(strings.Trim(fmt.Sprint(cp), "[]"), " ", " ", -1)
	// ["1,2" "3,4"] => "1,2 3,4"

	cmd := exec.Command("python", pyName, args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(string(out))
}
