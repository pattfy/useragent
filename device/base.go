package device

type base interface {
	ID() string
	Name() string
	UA() string

	match() bool
}

type core struct {
	ua string
}

func (c core) ID() string {
	return ""
}

func (c core) Name() string {
	return ""
}

func (c core) match() bool {
	return false
}

func (c core) UA() string {
	return c.ua
}
