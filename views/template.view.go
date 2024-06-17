package templates

import (
	"fmt"
	"html/template"
	"io"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/outsourcing-api/configs"
)

var Components = []string{
	"navbar.html",
	"sidebar.html",
	"footer.html",
	"js.html",
	"css.html",
}

func init() {
	for index, component := range Components {
		Components[index] = fmt.Sprintf("./views/html/components/%s", component)
	}
}

type Template struct {
}

func (t *Template) Render(w io.Writer, path string, data any, c echo.Context) error {
	ctx := c.Request().Context()
	path = strings.ReplaceAll(path, ".", "/")
	path = fmt.Sprintf("./views/html/%s.html", path)
	components := Components
	temps := []string{
		path,
	}
	temps = append(temps, components...)
	tmpl, err := template.ParseFiles(temps...)
	fmt.Println(err)
	if err != nil {
		return err
	}

	mappedData, _ := data.(map[string]any)
	mappedData["app_base_url"] = configs.AppConfig.Url
	mappedData["api_base_url"] = fmt.Sprintf("%s/api/v%s", configs.AppConfig.Url, configs.AppConfig.ApiVersion)
	mappedData["public_path"] = fmt.Sprintf("%s/public", configs.AppConfig.Url)
	mappedData["app_name"] = configs.AppConfig.Name
	mappedData["context"] = ctx
	if err := tmpl.Execute(w, mappedData); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

var DefaultTemplate = Template{}
