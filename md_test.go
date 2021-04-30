package md

import "testing"

func TestMBlock_Write(t *testing.T) {
	b := NewBlock()
	b.Add(BlockValue("标题1", H1))
	b.Add(BlockValue("zhengwen", Default))

	b.Add(BlockValue("标题2", H2), Value("标题2", H3))
	b.Write("test.md")
}
