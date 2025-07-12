package record

type Role struct {
	Index   int     `yaml:"index"`
	Script  string  `yaml:"script"`
	Records Records `yaml:"records,omitempty"`
}
