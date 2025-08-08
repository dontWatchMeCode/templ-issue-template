//go:build mage

package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

var Default = Find

func shellRun(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	fmt.Printf("#> Running: %s %v\n", name, args)
	return cmd.Run()
}

func Find() error {
	_, err := exec.Command("sh", "-c", "command -v fzf").Output()
	if err != nil {
		List()
		return nil
	}

	out, err := exec.Command("sh", "-c", "mage -l | grep -v 'Targets:' | grep -v '* default target' | grep -v 'find*' | awk '{print $1}' | sed '/^$/d' | fzf").Output()
	if err != nil {
		return err
	}
	target := string(out[:len(out)-1]) //? remove newline
	if target == "" {
		return nil
	}
	return shellRun("mage", target)
}

func RemoveTemplFiles() error {
	return shellRun("find", ".", "-type", "f", "-name", "*_templ.go", "-delete")
}

func Build() error {
	if err := RemoveTemplFiles(); err != nil {
		return err
	}

	if err := shellRun("templ", "generate", "-v"); err != nil {
		return err
	}

	if err := shellRun("go", "build", "-o", ".bin/main", "main.go"); err != nil {
		return err
	}

	return nil
}

func TidyAndUpdate() error {
	if err := shellRun("go", "install", "github.com/a-h/templ/cmd/templ@latest"); err != nil {
		return err
	}

	if err := shellRun("go", "get", "-u"); err != nil {
		return err
	}

	if err := shellRun("go", "mod", "tidy"); err != nil {
		return err
	}

	if err := shellRun("gofmt", "-w", "."); err != nil {
		return err
	}

	if err := shellRun("templ", "fmt", "."); err != nil {
		return err
	}

	if err := Build(); err != nil {
		return err
	}

	return nil
}

func List() error {
	return shellRun("mage", "-l")
}

func Dev() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	templCmd := exec.CommandContext(ctx, "templ", "generate", "--watch", "--open-browser=false")
	templCmd.Stdout = os.Stdout
	templCmd.Stderr = os.Stderr
	templCmd.Start()

	airCmd := exec.CommandContext(ctx, "air", "-c", "./air.toml")
	airCmd.Stdout = os.Stdout
	airCmd.Stderr = os.Stderr
	airCmd.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	cancel()
	templCmd.Wait()
	airCmd.Wait()
	return nil
}
