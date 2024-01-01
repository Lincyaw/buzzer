package stackcorruption

import (
	"buzzer/pkg/ebpf/ebpf"
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	prog, err := ebpf.New(1000, 1, 8)
	if err != nil {
		panic(err)
	}
	g := &Generator{}
	prog.Instructions = g.Generate(prog)
	for _, v := range prog.Instructions {
		fmt.Println(v.GeneratePoc())
		fmt.Println(v.GenerateBytecode())
	}
}
