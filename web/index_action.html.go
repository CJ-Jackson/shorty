package web

import (
	html "html/template"
)

const indexActionHeadHtml = `
<style type="text/css">
#welcomebox {
	margin-top: 2em;
	background: url("/static/img/landscape.jpg") center no-repeat;
}

#welcomebox .introbox {
	color: white;
	background-color: rgba(0, 0, 0, 0.5);
	padding: 1% 0 1% 4%;
	max-width: 35%;
	border-radius: 1em;
}
</style>
`

var indexActionBodyHtml = html.Must(html.New("indexActionBodyHtml").Parse(`
<div class="jumbotron" id="welcomebox">
	<div class="introbox">
		<h1>Welcome</h1>
		<p>to URL Shorter Demo</p>
	</div>
</div>
{{- .Flash -}}
<form method="post" novalidate>
	{{- .Csrf.Field -}}
	{{ .UrlForm.GetHtml }}
	<input class="form-control" type="submit" value="Submit">
</form>
<br>
<p><strong>Project URL:</strong> <a href="https://github.com/CJ-Jackson/shorty">https://github.com/CJ-Jackson/shorty</a>
<strong>Rules:</strong> Please don't do anything daft. :D
<strong>Note:</strong> URLs will expire in 24 hours</p>

<p>Powered by Google Go and MongoDB</p>
`))
