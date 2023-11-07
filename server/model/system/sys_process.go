package system

import (
	"errors"
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

func (container *SysProcessContainer) Push_back(process *SysInferenceProcess) error {
	container.mutex.Lock()
	defer container.mutex.Unlock()

	if len(container.processes) >= MaxSize {
		return errors.New("进程容器已满")
	}
	container.processes = append(container.processes, process)
	return nil
}

func (container *SysProcessContainer) pause_process(TaskID uint) error {
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

func (container *SysProcessContainer) kill_process(TaskID uint) error {
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

func (container *SysProcessContainer) awake_process(TaskID uint) error {
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
