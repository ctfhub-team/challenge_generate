package cmdutil

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func exec_cmd(command string) {
	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.Env = append(cmd.Env, "COMPOSE_DOCKER_CLI_BUILD=1", "DOCKER_BUILDKIT=1")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalln("stderr pipe ", err)
	}
	defer stderr.Close()
	if err := cmd.Start(); err != nil {
		log.Fatalln("start ", err)
	}
	go func() {
		serr := bufio.NewReader(stderr)
		for {
			line, _, err2 := serr.ReadLine()
			if err2 != nil || io.EOF == err2 {
				break
			}
			fmt.Println(string(line))
		}
	}()
	if err := cmd.Wait(); err != nil {
		log.Fatalln("wait ", err)
	}
}

// Deprecated 有 Bug 不建议使用
func exec_cmd_attach(command string) {
	cmd := exec.Command("/bin/bash", "-c", command)
	log.Println(cmd.Args)
	cmd.Env = append(cmd.Env, "COMPOSE_DOCKER_CLI_BUILD=1", "DOCKER_BUILDKIT=1", "COMPOSE_INTERACTIVE_NO_CLI=1")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalln("stdin pipe ", err)
	}
	defer stdin.Close()
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalln("stdout pipe ", err)
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		log.Fatalln("start ", err)
	}
	go func() {
		out := bufio.NewReader(stdout)
		for {
			line, _, err2 := out.ReadLine()
			if err2 != nil || io.EOF == err2 {
				break
			}
			fmt.Println(string(line))
		}
	}()

	io.Copy(stdin, bufio.NewReader(os.Stdin))

	if err := cmd.Wait(); err != nil {
		log.Fatalln("wait ", err)
	}
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
	exec_cmd_attach("docker-compose exec challenge sh")
}

func Log() {
	exec_cmd("docker-compose logs")
}

func Save() {
	fmt.Printf("TODO")
	// exec_cmd("docker-compose save")
}
