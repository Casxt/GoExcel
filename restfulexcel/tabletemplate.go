package restfulexcel

import "html/template"

type TemplateData struct {
	Sheets []Sheet
	Title  string
}

var totalTemplate = template.Must(template.New("total").Parse(`
{{template "header" .}}
{{template "body" .}}
{{template "footer" .}}
`))

var bodyTemplate = template.Must(totalTemplate.Parse(`
{{define "body"}}
<body class="container w-75 mt-5">
	<h1 class="text-center">{{.Title}}</h1>
	<form action="/test.xlsx" method="post">
		{{range .Sheets}}
			<table class="table table-bordered text-center">
				<!-- Table -->
				{{range .Table}}
					<tr>
						{{range .}}
							<th>{{.}}</th>
						{{end}}
					</tr>
				{{end}}
				<!-- Input -->
				
				<tr>
					<input type="hidden" name="sheet" value={{.Name}}>
					{{range $index, $value := .ColType}}
						<th class="p-0">
							{{if eq $value "select"}}
								<select name={{$index}} class="form-control border-0 rounded-0 text-center text-sm" style="width:5rem;" required>
									<option selected disabled hidden></option>
									<option value="true">Ture</option>
									<option value="false">False</option>
								</select>
							{{else}}
								<input type="{{$value}}" name={{$index}} class="form-control border-0 rounded-0 text-center" placeholder="input {{$value}}" required>
							{{end}}
						</th>
					{{end}}
				</tr>
			</table>
		{{end}}
		<button class="btn btn-primary rounded-0" type="submit" formmethod="post">Submit</button>
	</form>
</body>
{{end}}`))

var headerTemplate = template.Must(totalTemplate.Parse(`
{{define "header"}}
	<html>
		<head>
			<meta charset="utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
			<title>填写系统</title>
			<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">
		</head>
{{end}}`))

var footerTemplate = template.Must(totalTemplate.Parse(`
{{define "footer"}}
</html>
{{end}}`))
