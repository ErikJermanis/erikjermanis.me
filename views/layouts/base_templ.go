// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package layouts

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Base(title, description string, displayJsonLd bool) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><link rel=\"stylesheet\" type=\"text/css\" href=\"/public/styles.css\"><link rel=\"icon\" type=\"image/png\" href=\"/public/favicon.png\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if description != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<meta name=\"description\" content=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(description)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/layouts/base.templ`, Line: 12, Col: 50}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/layouts/base.templ`, Line: 14, Col: 17}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><script type=\"application/ld+json\">\n\t\t\t\t{\n\t\t\t\t\t\"@context\": \"https://schema.org\",\n\t\t\t\t\t\"@type\": \"Person\",\n\t\t\t\t\t\"name\": \"Erik Jermaniš\",\n\t\t\t\t\t\"url\": \"https://erikjermanis.me/\",\n\t\t\t\t\t\"jobTitle\": \"Software engineer\",\n\t\t\t\t\t\"sameAs\": [\n\t\t\t\t\t\t\"https://github.com/ErikJermanis\",\n\t\t\t\t\t\t\"https://x.com/ej_dev1\",\n\t\t\t\t\t\t\"https://www.linkedin.com/in/erik-jermanis/\"\n\t\t\t\t\t]\n\t\t\t\t}\n\t\t\t</script></head><body class=\"bg-slate-100 bg-coding-pattern dark:bg-neutral-900\"><button id=\"theme-switcher\" class=\"fixed top-2 right-2 text-white bg-neutral-900 dark:bg-slate-100 dark:text-black px-3 py-0.5 rounded-full text-sm uppercase\">light theme</button>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script type=\"text/javascript\">\n\t\t\t\tif (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {\n\t\t\t\t\tdocument.documentElement.classList.add('dark');\n\t\t\t\t\tdocument.getElementById('theme-switcher').textContent = 'light theme';\n\t\t\t\t} else {\n\t\t\t\t\tdocument.documentElement.classList.remove('dark');\n\t\t\t\t\tdocument.getElementById('theme-switcher').textContent = 'dark theme';\n\t\t\t\t}\n\t\t\t\tdocument.getElementById('theme-switcher').addEventListener('click', () => {\n\t\t\t\t\tif (document.documentElement.classList.contains('dark')) {\n\t\t\t\t\t\tdocument.documentElement.classList.remove('dark');\n\t\t\t\t\t\tlocalStorage.theme = 'light';\n\t\t\t\t\t\tdocument.getElementById('theme-switcher').textContent = 'dark theme';\n\t\t\t\t\t} else {\n\t\t\t\t\t\tdocument.documentElement.classList.add('dark');\n\t\t\t\t\t\tlocalStorage.theme = 'dark';\n\t\t\t\t\t\tdocument.getElementById('theme-switcher').textContent = 'light theme';\n\t\t\t\t\t}\n\t\t\t\t});\n\t\t\t</script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
