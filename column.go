package dbf

// Column represents a table column
type Column struct {
	Name     string
	Type     byte
	Address  uint32
	Length   uint8
	Position int
}

func (c *Column) String() string {
	return c.Name
}
