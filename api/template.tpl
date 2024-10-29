syntax = "v1"

info (
	title: // TODO: add title
	desc: // TODO: add description
	author: "{{.gitUser}}"
	email: "{{.gitEmail}}"
)

type request {
	// TODO: add members here and delete this comment
}

// Response This is a general response
type Response {
    Data interface{} `json:"data"`
    Flag bool        `json:"flag"`
    Msg  string      `json:"msg"`
    Code int         `json:"code"`
}

service {{.serviceName}} {
	@handler GetUser // TODO: set handler name and delete this comment
	get /users/id/:userId(request) returns(Response)

	@handler CreateUser // TODO: set handler name and delete this comment
	post /users/create(Response)
}