package command

import (
	"bufio"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"sync"

	"github.com/utkarshkrsingh/ripple/internal/config"
	"github.com/utkarshkrsingh/ripple/internal/log"
)

var wg sync.WaitGroup
var dependsOn = []string{}
var visited = make(map[string]bool)
var completed = make(map[string]bool)
var added = make(map[string]bool)

func ShowAllTasks() {
	log.Logger.Info("All the tasks are: ")
	for k := range config.Sections {
		if k == "variables" {
			continue
		}
		log.Logger.Println("-", k)
	}
}

func RunByName(taskName string) error {
	exists, ok := config.Sections[taskName]
	if !exists || !ok {
		return fmt.Errorf("%v does not exists", taskName)
	}

	err := fetchDependencies(taskName)
	if err != nil {
		return err
	}

	if len(dependsOn) > 0 {
	    log.Logger.Info("Running dependency...")
		for i := len(dependsOn) - 1; i >= 0; i-- {
			if !completed[dependsOn[i]] {
				completed[dependsOn[i]] = true
				command := config.Config.Get(fmt.Sprintf("%v.cmd", dependsOn[i])).(string)
				desc := config.Config.Get(fmt.Sprintf("%v.desc", dependsOn[i]))
				command = replacePlaceholders(command)
				log.Logger.Info(desc)
				wg.Add(1)
				runCommand(command, &wg)
			} else {
				log.Logger.Infof("Already executed %v...", dependsOn[i])
			}
		}
		wg.Wait()
	    log.Logger.Info("Dependency completed...")

	}

    mainCommand := config.Config.Get(fmt.Sprintf("%v.cmd", taskName)).(string)
    mainCommand = replacePlaceholders(mainCommand)
    wg.Add(1)
    go runCommand(mainCommand, &wg)
    wg.Wait()
	log.Logger.Info("Execution completed...")

	return nil
}

func fetchDependencies(taskName string) error {
	exists, ok := config.Sections[taskName]
	if !exists || !ok {
		log.Logger.Errorf("%v does not exists", taskName)
		return fmt.Errorf("%v does not exists", taskName)
	}

	if visited[taskName] {
		return nil
	}
	visited[taskName] = true

	if config.Config.IsSet(fmt.Sprintf("%v.depends_on", taskName)) {
		if deps, ok := config.Config.Get(fmt.Sprintf("%v.depends_on", taskName)).([]interface{}); ok {
			for _, dep := range deps {
				if !added[dep.(string)] {
					dependsOn = append(dependsOn, dep.(string))
					added[dep.(string)] = true

				}
			}
		} else if dep, ok := config.Config.Get(fmt.Sprintf("%v.depends_on", taskName)).(string); ok {
			if !added[dep] {
				dependsOn = append(dependsOn, dep)
				added[dep] = true

			}
		}
	}

	for _, dep := range dependsOn {
		err := fetchDependencies(dep)
		if err != nil {
			return err
		}
	}

	return nil
}

func replacePlaceholders(cmd string) string {
	variables := config.Config.GetStringMap("variables")

	for key, rawValue := range variables {

		value, exists := rawValue.(string)
		if !exists {
			continue
		}

		cmd = strings.ReplaceAll(cmd, fmt.Sprintf("${%s}", key), value)
		cmd = strings.ReplaceAll(cmd, fmt.Sprintf("${%s}", strings.ToUpper(key)), value)
	}

	return cmd
}

func runCommand(command string, wg *sync.WaitGroup) error {
	defer wg.Done()

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Logger.Errorf("Failed to create stdout pipe: %v", err)
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Logger.Errorf("Failed to create stderr pipe: %v", err)
		return err
	}

	if err := cmd.Start(); err != nil {
		log.Logger.Errorf("Failed to start command: %v", err)
		return err
	}

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			log.Logger.Info(scanner.Text())
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			log.Logger.Error(scanner.Text())
		}
	}()

	if err := cmd.Wait(); err != nil {
		log.Logger.Errorf("Command execution failed: %v", err)
		return err
	}

	return nil
}
