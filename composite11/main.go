package composite11

type IOrganization interface {
	Count() int
}

type Employee struct {
	Name string
}

func (e Employee) Count() int {
	return 1
}

type Department struct {
	Name             string
	SubOrganizations []IOrganization
}

func (d *Department) Count() int {
	c := 0
	for _, org := range d.SubOrganizations {
		c += org.Count()
	}
	return c
}

func (d *Department) AddSub(org IOrganization) {
	d.SubOrganizations = append(d.SubOrganizations, org)
}

func NewOrganization() IOrganization {
	root := &Department{Name: "name"}
	for i := 0; i < 10; i++ {
		root.AddSub(&Employee{})
		root.AddSub(&Employee{Name: "sub"})
	}
	return root
}
