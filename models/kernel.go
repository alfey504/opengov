package models

import "fmt"

type Kernel struct {
	kernel [][]float64
	x      int
	y      int
}

func CreateKernel(kernel [][]float64) (Kernel, error) {
	x := len(kernel)
	if x <= 0 {
		return Kernel{}, fmt.Errorf("kernel cannot be an empty slice in either dimension")
	}
	y := len(kernel[0])
	if y <= 0 {
		return Kernel{}, fmt.Errorf("kernel cannot be an empty slice in either dimension")
	}

	return Kernel{
		kernel: kernel,
		x:      x,
		y:      y,
	}, nil
}

func (kernel Kernel) GetDims() (int, int) {
	return kernel.x, kernel.y
}

func (kernel Kernel) GetKernel() [][]float64 {
	return kernel.kernel
}
