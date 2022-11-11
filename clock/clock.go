package clock

import "fmt"

type Clock int

func New(h, m int) Clock {
	minutes := h*60 + m
	if minutes < 0 {
		return Clock(1440 + minutes%1440)
	}
	return Clock(minutes % 1440)
}

func (c Clock) Add(m int) Clock {
	return New(0, int(c)+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(0, int(c)-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}
