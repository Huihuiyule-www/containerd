package containerd

// runtime handles containers, containers handle their own actions.
type Runtime interface {
	Create(id, bundlePath string) (Container, error)
}
