package main

import (
	"fmt"
	"time"
)

type PipeManager struct {
	pipes        map[int64]*Pipe
	spawnCounter float64
}

func NewPipeManager() PipeManager {
	return PipeManager{
		pipes:        make(map[int64]*Pipe),
		spawnCounter: 0,
	}
}

func (p *PipeManager) ManageLifecycle() {
	p.spawnCounter += dt

	if p.spawnCounter > 2 {
		newPipe := NewPipe()
		p.pipes[time.Now().UnixMicro()] = &newPipe
		p.spawnCounter = 0
		fmt.Println("new pipe!", p.pipes)
	}

	for key, pipe := range p.pipes {
		if pipe.x < -float64(pipe.width) {
			fmt.Println("deleting pipe!")
			delete(p.pipes, key)
		}
	}
}
