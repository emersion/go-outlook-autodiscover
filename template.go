package autodiscover

var mailTemplate = `<?xml version="1.0" encoding="utf-8" ?>
<Autodiscover xmlns="http://schemas.microsoft.com/exchange/autodiscover/responseschema/2006">
	<Response xmlns="http://schemas.microsoft.com/exchange/autodiscover/outlook/responseschema/2006a">
		{{if .DisplayName}}<User><DisplayName>{{.DisplayName}}</DisplayName></User>{{end}}

		<Account>
			<AccountType>email</AccountType>
			<Action>settings</Action>
			{{if .Image}}<Image>{{.Image}}</Image>{{end}}
			{{if .ServiceHome}}<ServiceHome>{{.ServiceHome}}</ServiceHome>{{end}}

			{{if .Imap}}
			<Protocol>
				<Type>IMAP</Type>
				<Server>{{.Imap.Hostname}}</Server>
				{{if .Imap.Port}}<Port>{{.Imap.Port}}</Port>{{end}}
				{{if .Imap.Username}}<LoginName>{{.Imap.Username}}</LoginName>{{end}}
				<SSL>{{if .Imap.Tls}}on{{else}}off{{end}}</SSL>
				<SPA>off</SPA>
			</Protocol>
			{{end}}

			{{if .Smtp}}
			<Protocol>
				<Type>SMTP</Type>
				<Server>{{.Smtp.Hostname}}</Server>
				{{if .Smtp.Port}}<Port>{{.Smtp.Port}}</Port>{{end}}
				{{if .Smtp.Username}}<LoginName>{{.Smtp.Username}}</LoginName>{{end}}
				<SSL>{{if .Smtp.Tls}}on{{else}}off{{end}}</SSL>
				<SPA>off</SPA>
			</Protocol>
			{{end}}
		</Account>
	</Response>
</Autodiscover>
`
