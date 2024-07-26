package day20

type Pulse string

const HIGH Pulse = "high"
const LOW Pulse = "low"

// -- Flip Flop --

type FlipFlop struct {
	ident       string
	on          bool
	outputs     []string
	shouldPulse bool
}

func (ff *FlipFlop) receive(p Pulse, from string) {
	// fmt.Printf("%v -%v-> %v\n", from, p, ff.ident)
	if p == HIGH {
		ff.shouldPulse = false
		return
	}

	ff.shouldPulse = true
	ff.on = !ff.on
}

func (ff *FlipFlop) pulse() Pulse {
	if !ff.shouldPulse {
		return ""
	}
	if ff.on {
		return HIGH
	}
	return LOW
}

func (ff *FlipFlop) receivers() []string {
	return ff.outputs
}

func (ff *FlipFlop) addInput(ident string) {}
func (ff *FlipFlop) inputCount() int       { return 0 }

// -- Conjunction --

type Conjunction struct {
	ident      string
	inputs     []string
	outputs    []string
	lastPulses map[string]Pulse
}

func (c *Conjunction) receive(p Pulse, from string) {
	// fmt.Printf("%v -%v-> %v\n", from, p, c.ident)
	if c.lastPulses == nil {
		c.lastPulses = make(map[string]Pulse)
	}
	c.lastPulses[from] = p
}

func (c *Conjunction) pulse() Pulse {
	for _, ident := range c.inputs {
		if val, ok := c.lastPulses[ident]; !ok || val == LOW {
			return HIGH
		}
	}
	return LOW
}

func (c *Conjunction) receivers() []string {
	return c.outputs
}

func (c *Conjunction) addInput(ident string) {
	c.inputs = append(c.inputs, ident)
}

func (c *Conjunction) inputCount() int {
	return len(c.inputs)
}

// -- Module

type Module interface {
	receive(Pulse, string)
	pulse() Pulse
	receivers() []string
	addInput(string)
	inputCount() int
}
