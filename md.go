package md

import (
	"fmt"
	"io/ioutil"
)

const (
	Default = iota
	H1
	H2
	H3
	H4
)

type Option func(*mBlock)

type Block interface {
	Add(...Option)
	Write(path string) error
	NewChild() Block
	Parent() Block
	children() []Block
}

func NewBlock() Block {
	return &mBlock{}
}

type mBlock struct {
	title int
	buf   []byte
	b     []*mBlock
}

func (m *mBlock) Add(opt ...Option) {
	if m.b == nil {
		m.b = newB()
	}
	for _, o := range opt {
		o(m)
	}
}

func (m *mBlock) Write(path string) error {
	return ioutil.WriteFile(path, m.parse(), 0700)
}

func (m *mBlock) NewChild() Block {
	m1 := &mBlock{}
	m.b = append(m.b, m1)
	return m1
}

func (m *mBlock) Parent() Block {
	if len(m.b) < 1 {
		return nil
	}
	return m.b[0]
}

func (m *mBlock) children() []Block {
	var b []Block
	for _, v := range m.b {
		b = append(b, v)
	}
	return b
}


func Value(content string, title int) Option {
	return func(m *mBlock) {
		var prefix string
		for i := 0; i < title; i++ {
			prefix += "#"
		}
		if prefix != "" {
			prefix += " "
		}
		content = fmt.Sprintf("%s%s\n", prefix, content)
		m.b = append(m.b, &mBlock{
			title: title,
			buf:   []byte(content),
		})
	}
}

func BlockValue(content string, title int) Option {
	return func(m *mBlock) {
		var prefix string
		for i := 0; i < title; i++ {
			prefix += "#"
		}
		if prefix != "" {
			prefix += " "
		}
		content = fmt.Sprintf("%s%s\n", prefix, content)
		m.b = append(m.b, &mBlock{
			title: title,
			buf:   []byte(content),
		})
	}
}

func (m *mBlock) parse() []byte {
	var temp []byte
	if len(m.b) <= 1 {
		return nil
	}
	stack := []*mBlock{m}
	for len(stack) > 0 {
		size := len(stack)
		var result [][]byte
		for i := 0; i < size; i++ {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if node == nil {
				continue
			}
			result = append(result, node.buf)
			for i := 1; i < len(node.b); i++ {
				stack = append(stack, node.b[i])
			}
		}
		for i := len(result) - 1; i >= 0; i-- {
			temp = append(temp, result[i]...)
		}
	}
	return temp
}

func newB() []*mBlock {
	return make([]*mBlock, 1)
}
