package model

type Crunchies []Crunchy

func GetCrunchies() Crunchies {
	return Crunchies{
		{Name: "Bob"}, {Name: "Tim"}, {Name: "Sarah"},
	}
}
