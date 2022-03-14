package main

import (
	"fmt"
	"time"
)

const minTopY = 40
const maxTopY = 150
const topYVariance = 30

type PipeManager struct {
	pipePairs    map[int64]*PipePair
	spawnCounter float64
	lastTopY     int
}

func NewPipeManager() PipeManager {
	return PipeManager{
		pipePairs:    make(map[int64]*PipePair),
		spawnCounter: 0,
		lastTopY:     0,
	}
}

func (p *PipeManager) ManageLifecycle() {
	p.spawnCounter += dt

	if p.spawnCounter > 2 {
		variance := randIntBetween(-topYVariance, topYVariance)
		topY := int(clamp(float64(p.lastTopY+variance), minTopY, maxTopY))
		p.lastTopY = topY

		newPipePair := NewPipePair(float64(topY))
		p.pipePairs[time.Now().UnixMicro()] = &newPipePair
		p.spawnCounter = 0
		fmt.Println("new pipe pair!", p.pipePairs)
	}

	for key, pipePair := range p.pipePairs {
		if pipePair.top.x < -float64(pipePair.top.width) {
			fmt.Println("deleting pipe!")
			delete(p.pipePairs, key)
		}
	}
}
