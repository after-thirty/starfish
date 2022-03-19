package version

import (
	"bytes"
	"runtime"
	"strings"
	"text/template"
)

// Build information.
var (
	Version   string
	GoVersion = runtime.Version()
)

var versionInfoTmpl = `
{{.program}}, version {{.version}} 
  go version:       {{.goVersion}}
  platform:         {{.platform}}
`

func Print(program string) string {
	m := map[string]string{
		"program":   program,
		"version":   Version,
		"goVersion": GoVersion,
		"platform":  runtime.GOOS + "/" + runtime.GOARCH,
	}
	t := template.Must(template.New("version").Parse(versionInfoTmpl))

	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, "version", m); err != nil {
		panic(err)
	}
	return strings.TrimSpace(buf.String())
}
