# go-outlook-autodiscover

[![GoDoc](https://godoc.org/github.com/ProtonMail/go-outlook-autodiscover?status.svg)](https://godoc.org/github.com/ProtonMail/go-outlook-autodiscover)

Generate [Microsoft Outlook autodiscover files](https://technet.microsoft.com/en-us/library/cc511507(v=office.14).aspx).

## Usage

```go
package main

import (
	"os"
	"github.com/ProtonMail/go-outlook-autodiscover"
)

func main() {
	c := &autodiscover.Config{
		DisplayName: "Mail Account",
		Imap: &mobileconfig.Imap{
			Hostname: "mail.nsa.gov",
			Port: 993,
			Tls: true,
			Username: "root",
		},
		Smtp: &mobileconfig.Smtp{
			Hostname: "mail.nsa.gov",
			Port: 25,
			Tls: false,
			Username: "root",
		},
	}

	if err := c.WriteTo(os.Stdout); err != nil {
		panic(err)
	}
}
```

## License

MIT
