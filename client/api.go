package client

import "errors"

const (
	APIDomain = "https://www.yuque.com/api/v2"

	// user
	APIUsers = APIDomain + "/users/%s"
	APIUser  = APIDomain + "/user"

	// repo
	APIUsersRepos  = APIDomain + "/users/%s/repos"
	APIGroupsRepos = APIDomain + "/groups/%s/repos"
	APIRepos       = APIDomain + "/repos/%s"
	APIReposToc    = APIDomain + "/repos/%s/toc"

	// doc
	APIReposDocs       = APIDomain + "/repos/%s/docs"
	APIReposDocsDetail = APIDomain + "/repos/%s/docs/%s"

	APIUsersGroups  = APIDomain + "/users/%s/groups"
	APIPublicGroups = APIDomain + "/groups"

	DefaultUserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36"

	SlugLength = 6

	RepositoryNamePrefix = "AutoCreate"
)

var (
	ErrCodeNotDefine = errors.New("未定义的错误码")

	ErrMsg = map[int]error{
		200: errors.New("成功"),
		400: errors.New("请求的参数不正确，或缺少必要信息，请对比文档"),
		401: errors.New("需要用户认证的接口用户信息不正确"),
		403: errors.New("缺少对应功能的权限"),
		404: errors.New("数据不存在，或未开放"),
		500: errors.New("服务器异常"),
	}
)

const (
	Book   = "Book"
	Design = "Design"
	All    = "All"
)

const (
	Private          = iota // 私密
	Open                    // 所有人
	SpaceMember             // 空间成员
	SpaceOpen               // 空间所有人
	RepositoryMember        // 知识库成员
)

const (
	Markdown = "markdown"
	Lake     = "lake"
)
