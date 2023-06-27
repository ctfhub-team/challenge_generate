package cmdutil

import (
	"bytes"
	"os/exec"
)

func exec_cmd(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.Env = append(cmd.Env, "COMPOSE_DOCKER_CLI_BUILD=1", "DOCKER_BUILDKIT=1")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func Auto() {
	Stop()
	Build()
	Run()
	Bash()
}

func Build() {
	exec_cmd("docker-compose build")
}

func Run() {
	exec_cmd("docker-compose up -d")
}

func Stop() {
	exec_cmd("docker-compose down")
}

func Bash() {
	exec_cmd("docker-compose exec challenge bash")
}
