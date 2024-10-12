package container

import (
	containerd "github.com/containerd/containerd/v2/client"
)

type VesselContainer struct {
	ID        string
	Container containerd.Container
	Task      containerd.Task
}

func NewVesselContainer(id string, container containerd.Container, task containerd.Task) *VesselContainer {
	return &VesselContainer{
		ID:        id,
		Container: container,
		Task:      task,
	}
}
