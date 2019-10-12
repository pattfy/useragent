package platform

type base interface {
	ID() string
	Name() string
	Version() string

	match() bool
}

type core struct {
	UA string
}

func (c core) ID() string {
	return ""
}

func (c core) Name() string {
	return ""
}

func (c core) Version() string {
	return ""
}

func (c core) match() bool {
	return false
}
