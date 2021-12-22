package main

import "testing"

func BenchmarkDirectCall(b *testing.B) {
	c := &Cat{Name: "draven"}
	for n := 0; n < b.N; n++ {
		// MOVQ	AX, "".c+24(SP)
		// MOVQ	AX, (SP)
		// CALL	"".(*Cat).Quack(SB)
		c.Quack()
	}
}
func BenchmarkDynamicDispatch(b *testing.B) {
	c := Duck(&Cat{Name: "draven"})
	for n := 0; n < b.N; n++ {
		// MOVQ	"".d+56(SP), AX
		// MOVQ	24(AX), AX
		// MOVQ	"".d+64(SP), CX
		// MOVQ	CX, (SP)
		// CALL	AX
		c.Quack()
	}
}
