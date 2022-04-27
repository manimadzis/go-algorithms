package sorts

// same interface as sort package has [sort.Interface]
type Sortable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
