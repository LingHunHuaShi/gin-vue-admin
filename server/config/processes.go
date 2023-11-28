package config

import (
	"errors"
	"log"
	"os/exec"
	"sync"
	"syscall"
)

const MaxSize = 4

type SysInferenceProcess struct {
	TaskID  uint
	Command *exec.Cmd
}

type SysProcessContainer struct {
	mutex     sync.Mutex
	processes []*SysInferenceProcess
}

func (container *SysProcessContainer) PushBack(process *SysInferenceProcess) error {
	container.mutex.Lock()
	defer container.mutex.Unlock()

	if len(container.processes) >= MaxSize {
		return errors.New("进程容器已满")
	}
	cmd := process.Command
	err := cmd.Start()
	if err != nil {
		log.Println("Failed to exec command.")
		return err
	}
	container.processes = append(container.processes, process)
	log.Println("Succeed to exec command.")
	return nil
}

func (container *SysProcessContainer) PauseProcess(TaskID uint) error {
	container.mutex.Lock()
	defer container.mutex.Unlock()

	for _, process := range container.processes {
		if process.TaskID == TaskID {
			if err := process.Command.Process.Signal(syscall.SIGSTOP); err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("无效的任务ID")
}

func (container *SysProcessContainer) KillProcess(TaskID uint) error {
	container.mutex.Lock()
	defer container.mutex.Unlock()

	for i, process := range container.processes {
		if process.TaskID == TaskID {
			_ = process.Command.Process.Kill()

			container.processes = append(container.processes[:i], container.processes[i+1:]...)
			return nil
		}
	}
	return errors.New("无效的任务ID")
}

func (container *SysProcessContainer) AwakeProcess(TaskID uint) error {
	container.mutex.Lock()
	defer container.mutex.Unlock()

	for _, process := range container.processes {
		if process.TaskID == TaskID {
			if err := process.Command.Process.Signal(syscall.SIGCONT); err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("无效的任务ID")
}
