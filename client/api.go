package client

import "errors"

// API 路径
const (
	APIDomain = "https://www.yuque.com/api/v2"
	APIHello  = APIDomain + "/hello"

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

	// group
	APIUsersGroups       = APIDomain + "/users/%s/groups"
	APIGroups            = APIDomain + "/groups"
	APIGroupsDetail      = APIDomain + "/groups/%s"
	APIGroupsUsers       = APIDomain + "/groups/%s/users"
	APIGroupsUsersUpdate = APIDomain + "/groups/%s/users/%s"

	SlugLength = 6

	RepositoryNamePrefix = "AutoCreate"
)

// DefaultUserAgent 默认的 user agent
const DefaultUserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36"

// 知识库类型
const (
	Book   = "Book"
	Design = "Design"
	All    = "All"
)

// 知识库访问权限
const (
	Private          = iota // 私密
	Open                    // 所有人
	SpaceMember             // 空间成员
	SpaceOpen               // 空间所有人
	RepositoryMember        // 知识库成员
)

// 文档格式
const (
	Markdown = "markdown"
	Lake     = "lake"
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
