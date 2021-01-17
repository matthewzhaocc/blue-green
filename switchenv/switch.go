package main

import (
	"os"
	"os/exec"

	"github.com/matthewzhaocc/common"
	log "github.com/sirupsen/logrus"
)

const (
	// GREEN the RED of the netflix
	GREEN = "green"
	// BLUE the BLACK of the netflix
	BLUE = "blue"
)
// ApplyManifest applies the manifest of the environment
func ApplyManifest(color string) error {
	command := "apply -f " + color + "-env.yml"
	cmd := exec.Command("kubectl", command)
	return cmd.Run()
}
// FindEnvColor finds the color of the environment
func FindEnvColor() string {
	// thisis is a generic implementation based on environment variables, feel free to eidt it if you want to use a different method
	return os.Getenv("CURRENT_ENV_COLOR")
}
func init() {
	log.SetOutput(os.Stdout)
}
func main() {
	if FindEnvColor() == GREEN {
		// applying blue environment manifest
		log.Println("applying *blue* environment")
		err := ApplyManifest(BLUE)
		if common.CheckError(err) {
			log.Println(err.Error())
		}
	} else if FindEnvColor() == BLUE {
		// applying the green environment manifest
		log.Println("applying *green* environment")
		err := ApplyManifest(GREEN)
		if common.CheckError(err) {
			log.Println(err.Error())
		}
	} else {
		log.Println("Can't find the current environment colorz")
	}
	
}