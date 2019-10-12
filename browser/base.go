package browser

// base declare the methods for browser.core
type base interface {
	ID() string
	Name() string
	FullVersion() string

	match() bool
}

// core struct contains the raw User-Agent
type core struct {
	UA string
}

// FullVersion returns the full version info.
func (c core) FullVersion() string {
	return ""
}

// ID returns the identifier of the browser(in lowercase)
func (c core) ID() string {
	return ""
}

// Name returns the Name of the browser
func (c core) Name() string {
	return ""
}

// Match checks whether the User-Agent matches
func (c core) match() bool {
	return false
}
