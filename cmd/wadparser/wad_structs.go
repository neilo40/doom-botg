package main

type Wad struct {
	WadType string
	Levels  []Level
}

type Level struct {
	Name         string
	Linedefs     []Linedef
	Start        Vertex
	Exit         Vertex
	DoorSwitches []DoorSwitch
}

type Linedef struct {
	A              Vertex
	B              Vertex
	NonTraversable bool
	SectorTag      *int     // may be nil
	LineType       *int     // may be nil
	RightSidedef   *Sidedef // may be nil
	LeftSidedef    *Sidedef // may be nil
}

type Vertex struct {
	X float64
	Y float64
}

type DoorSwitch struct {
	A        Vertex
	B        Vertex
	Switched bool
}

type Sidedef struct {
	SectorTag int
	Sector    *Sector // may be nil
}

type Sector struct {
	SectorType    int
	Tag           int
	FloorHeight   int
	CeilingHeight int
}

type Lump struct {
	Name string
	Data []byte
}

type Thing struct {
	Position Vertex
	Facing   int
	DoomId   int
}
