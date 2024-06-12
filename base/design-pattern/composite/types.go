package composite

// Composite 组织
type Composite interface {
	AddComposite(composite Composite)
	Features()
}
