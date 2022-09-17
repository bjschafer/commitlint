package lint

type SeverityLevel int

const (
	Hint SeverityLevel = iota
	Warning
	Error
)

type Problem struct {
	Level   SeverityLevel `yaml:"level"`
	Name    string        `yaml:"name"`
	Message string        `yaml:"message"`
}
