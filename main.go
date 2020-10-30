package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"github.com/keegancsmith/shell"
)

func main() {
	regex, _ := regexp.Compile("^TF_VAR_[^=]+=.+$")
	envVars := os.Environ()
	for _, envVar := range envVars {
		if regex.MatchString(envVar) {
			pair := strings.SplitN(envVar, "=", 2)
			envVarWithoutTfVarLower := strings.ToLower(strings.TrimPrefix(pair[0], "TF_VAR_"))
			newEnvVarKey := "TF_VAR_"+envVarWithoutTfVarLower

			fmt.Println("export " + newEnvVarKey + "=" + shell.EscapeArg(pair[1]))
		}
	}
}
