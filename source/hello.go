// hello how are you doing
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>htmx Go Example</title>
    <script src="/static/htmx.min.js" defer></script>
</head>
<body>
	<div hx-post="/test" hx-trigger="every 2s"> Hello this is it</div>

</body>
</html>
`))

type PageVariables struct {
	Items []string
}

var items []string

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/test", test_press)
	http.HandleFunc("/add", AddItem)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func test_press(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HELLO")
	http.Error(w, "NOOOO", 286)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	pageVariables := PageVariables{
		Items: items,
	}

	err := templates.ExecuteTemplate(w, "", pageVariables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AddItem(w http.ResponseWriter, r *http.Request) {
	newItem := r.FormValue("item")
	items = append(items, newItem)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
