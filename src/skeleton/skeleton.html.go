package skeleton

import (
	html "html/template"
)

var skeletonHtml = html.Must(html.New("skeletonHtml").Parse(`<!DOCTYPE html>
<html lang="en-GB">
	<head>
		<title>{{.Title}} | Shorty</title>
		<link rel="stylesheet" href="/static/css/bootstrap.css">
		<link rel="stylesheet" href="/static/css/bootstrap-theme.min.css">
		<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
		{{.Head}}
	</head>
	<body>
		<div class="container">
			{{.Body}}
			{{- .Footer -}}
			<script src="/static/js/jquery-2.2.3.min.js"></script>
			<script src="/static/js/bootstrap.min.js"></script>
			{{.Javascript}}
		</div>
	</body>
</html>`))
