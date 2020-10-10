package client

import (
	"encoding/json"
	"github.com/jdxj/yuque/models"
	"strconv"
)

type Response struct {
	Data json.RawMessage `json:"data"`
}

type Doc struct {
	Create bool `json:"create"`
}

type Abilities struct {
	Update  bool `json:"update"`
	Destroy bool `json:"destroy"`
	Doc     Doc  `json:"doc"`
}

type BookSerializer struct {
	ID               int                   `json:"id"`
	Type             string                `json:"type"`
	Slug             string                `json:"slug"`
	Name             string                `json:"name"`
	Namespace        string                `json:"namespace"`
	UserID           int                   `json:"user_id"`
	User             models.UserSerializer `json:"user"`
	Description      string                `json:"description"`
	CreatorID        int                   `json:"creator_id"`
	Public           int                   `json:"public"`
	ItemsCount       int                   `json:"items_count"`
	LikesCount       int                   `json:"likes_count"`
	WatchesCount     int                   `json:"watches_count"`
	ContentUpdatedAt string                `json:"content_updated_at"`
	CreatedAt        string                `json:"created_at"`
	UpdatedAt        string                `json:"updated_at"`
	Serializer       string                `json:"_serializer"`
}

type BookDetailSerializer struct {
	BookSerializer
	Toc        string `json:"toc"`
	TocYml     string `json:"toc_yml"`
	PinnedAt   string `json:"pinned_at"`
	ArchivedAt string `json:"archived_at"`

	Abilities Abilities `json:"abilities"`
}

type DocSerializer struct {
	ID                int            `json:"id"`
	Slug              string         `json:"slug"`
	Title             string         `json:"title"`
	Description       string         `json:"description"`
	UserID            int                   `json:"user_id"`
	BookID            int                   `json:"book_id"`
	Format            string                `json:"format"`
	Public            int                   `json:"public"`
	Status            int                   `json:"status"`
	ViewStatus        int                   `json:"view_status"`
	ReadStatus        int                   `json:"read_status"`
	LikesCount        int                   `json:"likes_count"`
	CommentsCount     int                   `json:"comments_count"`
	ContentUpdatedAt  string                `json:"content_updated_at"`
	Book              BookSerializer        `json:"book"`
	LastEditor        models.UserSerializer `json:"last_editor"`
	CreatedAt         string                `json:"created_at"`
	UpdatedAt         string                `json:"updated_at"`
	PublishedAt       string                `json:"published_at"`
	FirstPublishedAt  string                `json:"first_published_at"`
	DraftVersion      int                   `json:"draft_version"`
	LastEditorID      int                   `json:"last_editor_id"`
	WordCount         int                   `json:"word_count"`
	Cover             string                `json:"cover"`
	CustomDescription string                `json:"custom_description"`
	Serializer        string                `json:"_serializer"`
}

type DocDetailSerializer struct {
	DocSerializer

	Creator       models.UserSerializer `json:"creator"`
	Body          string                `json:"body"`
	BodyDraft     string                `json:"body_draft"`
	BodyHTML      string                `json:"body_html"`
	BodyLake      string                `json:"body_lake"`
	BodyDraftLake string                `json:"body_draft_lake"`
	CreatorID     int                   `json:"creator_id"`
	DeletedAt     string                `json:"deleted_at"`
	Hits          int                   `json:"hits"`
}

type Toc struct {
	Type        string `json:"type"`
	Title       string `json:"title"`
	Uuid        string `json:"uuid"`
	Url         string `json:"url"`
	PrevUuid    string `json:"prev_uuid"`
	SiblingUuid string `json:"sibling_uuid"`
	ChildUuid   string `json:"child_uuid"`
	ParentUuid  string `json:"parent_uuid"`
	Level       int    `json:"level"`
	OpenWindow  int    `json:"open_window"`
	Visible     int    `json:"visible"`
	Depth       int    `json:"depth"`
	Slug        string `json:"slug"`

	ID    json.RawMessage `json:"id"`
	DocID json.RawMessage `json:"doc_id"`
}

func (toc *Toc) UnmarshalID() string {
	var res string
	if toc.Type == "DOC" {
		tmp := new(int)
		_ = json.Unmarshal(toc.ID, tmp)
		res = strconv.Itoa(*tmp)
	} else if toc.Type == "LINK" {
		_ = json.Unmarshal(toc.ID, &res)
	}
	return res
}

func (toc *Toc) UnmarshalDocID() string {
	var res string
	if toc.Type == "DOC" {
		tmp := new(int)
		_ = json.Unmarshal(toc.DocID, tmp)
		res = strconv.Itoa(*tmp)
	} else if toc.Type == "LINK" {
		_ = json.Unmarshal(toc.DocID, &res)
	}
	return res
}
