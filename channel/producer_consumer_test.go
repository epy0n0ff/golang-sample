package main
import (
	"testing"
)

type Consumer struct {
	ch   chan *int
	done chan struct{}
}

type Producer struct {
	n int
	ch chan *int
}

func (p Producer) Start() {
	go func() {
		for i := 0; i < p.n; i++ {
			p.ch <- &i
		}
		defer close(p.ch)
	}()
}

func (c Consumer) Start() {
	go func() {
		for {
			i := <-c.ch
			if i == nil {
				break
			}
		}
		c.done <- struct{}{}
	}()
}

func BenchmarkP1C1Buffer1(b *testing.B) {
	ch := make(chan *int, 1)
	done := make(chan struct{})
	p := Producer{b.N,ch}
	c := Consumer{ch, done}

	b.ResetTimer()
	p.Start()
	c.Start()
	<-done
}


func BenchmarkP1C1Buffer2(b *testing.B) {
	ch := make(chan *int, 2)
	done := make(chan struct{})
	p := Producer{b.N,ch}
	c := Consumer{ch, done}

	b.ResetTimer()
	p.Start()
	c.Start()
	<-done
}

func BenchmarkP1C1Buffer10(b *testing.B) {
	ch := make(chan *int, 10)
	done := make(chan struct{})
	p := Producer{b.N,ch}
	c := Consumer{ch, done}

	b.ResetTimer()
	p.Start()
	c.Start()
	<-done
}

func BenchmarkP1C1Buffer100(b *testing.B) {
	ch := make(chan *int, 100)
	done := make(chan struct{})
	p := Producer{b.N,ch}
	c := Consumer{ch, done}

	b.ResetTimer()
	p.Start()
	c.Start()
	<-done
}

func BenchmarkP1C1Buffer1000(b *testing.B) {
	ch := make(chan *int, 1000)
	done := make(chan struct{})
	p := Producer{b.N,ch}
	c := Consumer{ch, done}

	b.ResetTimer()
	p.Start()
	c.Start()
	<-done
}

func BenchmarkP1C1Buffer10000(b *testing.B) {
	ch := make(chan *int, 10000)
	done := make(chan struct{})
	p := Producer{b.N,ch}
	c := Consumer{ch, done}

	b.ResetTimer()
	p.Start()
	c.Start()
	<-done
}
