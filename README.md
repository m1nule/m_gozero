# m_gozero

> goctl version 1.7.3 darwin/arm64
>
> rm -rf ~/.goctl/1.7.3
>
> mkdir -pv ~/.goctl/1.7.3/api
>
> vim ~/.goctl/1.7.3/api/handler.tpl

```go
package {{.PkgName}}

import (
"net/http"

"github.com/m1nule/m_gozero/response"
"github.com/zeromicro/go-zero/rest/httpx"
{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
return func(w http.ResponseWriter, r *http.Request) {
{{if .HasRequest}}var req types.{{.RequestType}}
if err := httpx.Parse(r, &req); err != nil {
response.Response(w, struct{}{}, err)
return
}

{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
if err != nil {
response.Response(w, resp, err)
} else {
{{if .HasResp}}httpx.OkJson(w, resp){{else}}httpx.Ok(w){{end}}
}
}
}
```
