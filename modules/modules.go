package modules

type BookSerializer struct {
	ID           int             `json:"id"`
	Type         string          `json:"type"`
	Slug         string          `json:"slug"`
	Name         string          `json:"name"`
	Namespace    string          `json:"namespace"`
	UserID       int             `json:"user_id"`
	User         *UserSerializer `json:"user"`
	Description  string          `json:"description"`
	CreatorID    int             `json:"creator_id"`
	Public       int             `json:"public"`
	LikesCount   int             `json:"likes_count"`
	WatchesCount int             `json:"watches_count"`
	CreatedAt    string          `json:"created_at"`
	UpdatedAt    string          `json:"updated_at"`
}

type BookDetailSerializer struct {
	BookSerializer

	// BookDetailSerializer 与 BookSerializer 不同的字段
	ItemsCount int    `json:"items_count"` // item_count
	TocYml     string `json:"toc_yml"`

	// 实际 API 返回的多出的字段
	ContentUpdatedAt string `json:"content_updated_at"`
	Serializer       string `json:"_serializer"`

	// 删除知识库返回的
	Toc        string `json:"toc"`
	PinnedAt   string `json:"pinned_at"`
	ArchivedAt string `json:"archived_at"`
}

type UserSerializer struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Login     string `json:"login"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`

	// 实际 API 返回的多出的字段
	Description      string `json:"description"`
	BooksCount       int    `json:"books_count"`
	PublicBooksCount int    `json:"public_books_count"`
	FollowersCount   int    `json:"followers_count"`
	FollowingCount   int    `json:"following_count"`
	Serializer       string `json:"_serializer"`
	SpaceID          int    `json:"space_id"`
	AccountID        int    `json:"account_id"`
	Public           int    `json:"public"`
}

type DocSerializer struct {
	ID               int             `json:"id"`
	Slug             string          `json:"slug"`
	Title            string          `json:"title"`
	UserID           int             `json:"user_id"`
	Format           string          `json:"format"`
	Public           int             `json:"public"`
	Status           int             `json:"status"`
	LikesCount       int             `json:"likes_count"`
	CommentsCount    int             `json:"comments_count"`
	ContentUpdatedAt string          `json:"content_updated_at"`
	Book             *BookSerializer `json:"book"`
	User             *UserSerializer `json:"user"`
	LastEditor       *UserSerializer `json:"last_editor"`
	CreatedAt        string          `json:"created_at"`
	UpdatedAt        string          `json:"updated_at"`
}

type DocDetailSerializer struct {
	DocSerializer

	// 相对 DocSerializer 多出的
	BookID    int    `json:"book_id"`
	Body      string `json:"body"`
	BodyDraft string `json:"body_draft"`
	BodyHTML  string `json:"body_html"`
	BodyLake  string `json:"body_lake"`
	CreatorID int    `json:"creator_id"`
	DeletedAt string `json:"deleted_at"`

	// 实际返回多出的
	Description       string `json:"description"`
	ViewStatus        int    `json:"view_status"`
	ReadStatus        int    `json:"read_status"`
	PublishedAt       string `json:"published_at"`
	FirstPublishedAt  string `json:"first_published_at"`
	DraftVersion      int    `json:"draft_version"`
	LastEditorID      int    `json:"last_editor_id"`
	WordCount         int    `json:"word_count"`
	Cover             string `json:"cover"`
	CustomDescription string `json:"custom_description"`
	Serializer        string `json:"_serializer"`
}
