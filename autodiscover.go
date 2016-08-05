package autodiscover

import (
	"io"
	"text/template"
)

var t = template.Must(template.New("autodiscover").Parse(mailTemplate))

// A Microsoft Outlook autodiscover file.
type Config struct {
	DisplayName string
	Image string
	ServiceHome string

	Imap *Imap
	Smtp *Smtp
}

type Imap struct {
	Hostname string
	Port int
	Tls bool

	Username string
}

type Smtp struct {
	Hostname string
	Port int
	Tls bool

	Username string
}

func (c *Config) WriteTo(w io.Writer) error {
	return t.Execute(w, c)
}
