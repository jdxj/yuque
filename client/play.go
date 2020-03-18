package client

import (
	"log"
	"sync"

	"github.com/jdxj/yuque/utils"

	"github.com/jdxj/yuque/modules"
)

// CreateDocAmount 创建 amount 个文档, 该文档在 AutoCreate 知识库下
func (c *Client) CreateDocAmount(amount int) {
	repos, err := c.ListOwnUserRepositories()
	if err != nil {
		log.Fatalln(err)
	}

	var autoCreated *modules.BookSerializer
	for _, repo := range repos {
		if repo.Name == RepositoryNamePrefix {
			autoCreated = repo
			break
		}
	}

	if autoCreated == nil {
		log.Fatalf("not found auto create reposity")
	}

	log.Println(autoCreated.Namespace)

	for i := 0; i < amount; i++ {
		log.Println("创建请求")
		title := utils.GenRandString(6)
		docReq := NewCreateDocRequestSlug(title, title, Intranet, Markdown)

		log.Println("发送数据")
		if _, err := c.CreateDoc(autoCreated.Namespace, docReq); err != nil {
			log.Println(err)
		}
	}
}

// DeleteAutoCreate 删除 auto create 知识库
func (c *Client) DeleteAutoCreate() {
	repos, err := c.ListOwnUserRepositories()
	if err != nil {
		log.Fatalln(err)
	}

	for _, repo := range repos {
		if repo.Name == RepositoryNamePrefix {
			if _, err := c.DeleteRepository(repo.Namespace); err != nil {
				log.Println(err)
			} else {
				log.Printf("delete: %s\n", repo.Namespace)
			}
		}
	}
}

// 创建 amount 个知识库, 每个知识库有中有100个文档
func (c *Client) CreateRepoDoc(amount int) {
	go c.continueCreateRepo(amount)
	c.continueCreate100Doc()
}

// continueCreateRepo 连续创建 auto create 知识库
func (c *Client) continueCreateRepo(amount int) {
	for i := 0; i < amount; i++ {
		repoReq := NewCreateRepositoryRequestSlug(RepositoryNamePrefix, "", Book, Open)
		book, err := c.CreateUserRepository(repoReq)
		if err != nil {
			log.Println(err)
			break
		} else {
			log.Printf("AutoCreate finish: %s, id: %d\n", book.Slug, i)
		}

		// c.user 必不为空
		user := c.user.Login
		slug := book.Slug
		namespace := user + "/" + slug
		c.namespaceTask <- namespace
	}

	close(c.namespaceTask)
}

func (c *Client) continueCreate100Doc() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			c.create100Doc()
			wg.Done()
		}()
	}

	wg.Wait()
	log.Println("wg done")
}

func (c *Client) create100Doc() {
	for namespace := range c.namespaceTask {
		for j := 0; j < 100; j++ {
			title := utils.GenRandString(6)
			docReq := NewCreateDocRequestSlug(title, title, Intranet, Markdown)
			if doc, err := c.CreateDoc(namespace, docReq); err != nil {
				log.Println(err)
				continue
			} else {
				log.Printf("create doc finish: %s, id: %d\n", doc.Slug, j)
			}
		}
	}
}
