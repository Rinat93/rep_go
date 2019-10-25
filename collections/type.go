package collections

type Collections interface {
	New()
	Add(name string, dataArr []string)
	AddAsync(name string, dataArr []string, res chan<- *Collection)
	Remove()
}
