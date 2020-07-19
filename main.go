package main

import (
	"fmt"

	// "github.com/spf13/cobra/cobra/cmd"
)

// func main() {
// 	fmt.Print("this is where the code will be executed")
// 	cmd.Execute
// }

func main() {
	host := "0.0.0.0:8888"
	http.HandleFunc("/", getCandidateInfo)
	logger.Log.Info("Server started: http://" + host)

	err := http.ListenAndServe(host, nil)
	if err != nil {
		logger.Log.Error(err)
		return
	}
}