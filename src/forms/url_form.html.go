package forms

import (
	html "html/template"
)

var urlFormHtml = html.Must(html.New("urlFormHtml").Parse(`
{{- $url := .UrlField -}}
<div class="form-group {{if .HasUrlErr -}} has-error {{- end}}">
	<label class="control-label" for="urlFormHtml- {{- $url.Name -}} "> {{- $url.Label -}} :</label> Should include 'http://' or 'https://'
	<input type="url" class="form-control" id="urlFormHtml- {{- $url.Name -}} " name=" {{- $url.Name -}} "
		maxlength=" {{- $url.MaxChar -}} " value=" {{- .UrlNorm -}} ">
	{{ .UrlErrHtml }}
</div>
`))
