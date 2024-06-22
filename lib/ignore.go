package lib

type Ignore struct {
	Paths []string
}

func (i *Ignore) IsIgnored(path string) bool {
	for _, p := range i.Paths {
		if p == path {
			return true
		}
	}
	return false
}

func (i *Ignore) Add(path string) {
	i.Paths = append(i.Paths, path)
}

func NewIgnore(paths []string) *Ignore {
	return &Ignore{
		Paths: paths,
	}
}
