package firecracker

import (
	"os"
	"path/filepath"
	"sync"

	firecrackersdk "github.com/firecracker-microvm/firecracker-go-sdk"

	"github.com/solo-io/unik/pkg/config"
	"github.com/solo-io/unik/pkg/state"
)

type FirecrackerProvider struct {
	config config.Firecracker
	state  state.State

	runningMachines map[string]*firecrackersdk.Machine
	mapLock         sync.RWMutex
}

func FirecrackerStateFile() string {
	return filepath.Join(config.Internal.UnikHome, "firecracker/state.json")

}
func firecrackerImagesDirectory() string {
	return filepath.Join(config.Internal.UnikHome, "firecracker/images/")
}

func firecrackerInstancesDirectory() string {
	return filepath.Join(config.Internal.UnikHome, "firecracker/instances/")
}

func firecrackerVolumesDirectory() string {
	return filepath.Join(config.Internal.UnikHome, "firecracker/volumes/")
}

func NewProvider(config config.Firecracker) (*FirecrackerProvider, error) {

	os.MkdirAll(firecrackerImagesDirectory(), 0777)
	os.MkdirAll(firecrackerInstancesDirectory(), 0777)
	os.MkdirAll(firecrackerVolumesDirectory(), 0777)

	p := &FirecrackerProvider{
		config:          config,
		state:           state.NewBasicState(FirecrackerStateFile()),
		runningMachines: map[string]*firecrackersdk.Machine{},
	}

	return p, nil
}

func (p *FirecrackerProvider) WithState(state state.State) *FirecrackerProvider {
	p.state = state
	return p
}

func getImagePath(imageName string) string {
	return filepath.Join(firecrackerImagesDirectory(), imageName, "boot.img")
}

func getVolumePath(volumeName string) string {
	return filepath.Join(firecrackerVolumesDirectory(), volumeName, "data.img")
}

func getInstanceDir(instanceName string) string {
	return filepath.Join(firecrackerInstancesDirectory(), instanceName)
}

func getImageDir(imageName string) string {
	return filepath.Join(firecrackerImagesDirectory(), imageName)
}
