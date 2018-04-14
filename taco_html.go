package main

func metaRefreshHTML(content string) string {
	return `
  <!DOCTYPE html>
  <html>
    <head>
      <meta charset="utf-8">
      <title>taco.cc</title>
      <meta http-equiv="refresh" content="1">
    </head>
    <body>
      <pre>` + content + `</pre>
    </body>
  </html>
  `
}
