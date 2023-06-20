// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: crawler.proto

package crawlerproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Crawler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceLink         string `protobuf:"bytes,1,opt,name=source_link,json=sourceLink,proto3" json:"source_link,omitempty"`
	FeedLink           string `protobuf:"bytes,2,opt,name=feed_link,json=feedLink,proto3" json:"feed_link,omitempty"`
	CrawlType          string `protobuf:"bytes,3,opt,name=crawl_type,json=crawlType,proto3" json:"crawl_type,omitempty"`
	ArticleDiv         string `protobuf:"bytes,4,opt,name=article_div,json=articleDiv,proto3" json:"article_div,omitempty"`
	ArticleTitle       string `protobuf:"bytes,5,opt,name=article_title,json=articleTitle,proto3" json:"article_title,omitempty"`
	ArticleDescription string `protobuf:"bytes,6,opt,name=article_description,json=articleDescription,proto3" json:"article_description,omitempty"`
	ArticleLink        string `protobuf:"bytes,7,opt,name=article_link,json=articleLink,proto3" json:"article_link,omitempty"`
	ArticleAuthors     string `protobuf:"bytes,8,opt,name=article_authors,json=articleAuthors,proto3" json:"article_authors,omitempty"`
	Schedule           string `protobuf:"bytes,9,opt,name=schedule,proto3" json:"schedule,omitempty"`
	ArticlesSourceId   int32  `protobuf:"varint,10,opt,name=articles_source_id,json=articlesSourceId,proto3" json:"articles_source_id,omitempty"`
}

func (x *Crawler) Reset() {
	*x = Crawler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Crawler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Crawler) ProtoMessage() {}

func (x *Crawler) ProtoReflect() protoreflect.Message {
	mi := &file_crawler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Crawler.ProtoReflect.Descriptor instead.
func (*Crawler) Descriptor() ([]byte, []int) {
	return file_crawler_proto_rawDescGZIP(), []int{0}
}

func (x *Crawler) GetSourceLink() string {
	if x != nil {
		return x.SourceLink
	}
	return ""
}

func (x *Crawler) GetFeedLink() string {
	if x != nil {
		return x.FeedLink
	}
	return ""
}

func (x *Crawler) GetCrawlType() string {
	if x != nil {
		return x.CrawlType
	}
	return ""
}

func (x *Crawler) GetArticleDiv() string {
	if x != nil {
		return x.ArticleDiv
	}
	return ""
}

func (x *Crawler) GetArticleTitle() string {
	if x != nil {
		return x.ArticleTitle
	}
	return ""
}

func (x *Crawler) GetArticleDescription() string {
	if x != nil {
		return x.ArticleDescription
	}
	return ""
}

func (x *Crawler) GetArticleLink() string {
	if x != nil {
		return x.ArticleLink
	}
	return ""
}

func (x *Crawler) GetArticleAuthors() string {
	if x != nil {
		return x.ArticleAuthors
	}
	return ""
}

func (x *Crawler) GetSchedule() string {
	if x != nil {
		return x.Schedule
	}
	return ""
}

func (x *Crawler) GetArticlesSourceId() int32 {
	if x != nil {
		return x.ArticlesSourceId
	}
	return 0
}

type TestResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Articles       []*Article      `protobuf:"bytes,1,rep,name=articles,proto3" json:"articles,omitempty"`
	ArticlesSource *ArticlesSource `protobuf:"bytes,2,opt,name=articles_source,json=articlesSource,proto3" json:"articles_source,omitempty"`
}

func (x *TestResult) Reset() {
	*x = TestResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResult) ProtoMessage() {}

func (x *TestResult) ProtoReflect() protoreflect.Message {
	mi := &file_crawler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResult.ProtoReflect.Descriptor instead.
func (*TestResult) Descriptor() ([]byte, []int) {
	return file_crawler_proto_rawDescGZIP(), []int{1}
}

func (x *TestResult) GetArticles() []*Article {
	if x != nil {
		return x.Articles
	}
	return nil
}

func (x *TestResult) GetArticlesSource() *ArticlesSource {
	if x != nil {
		return x.ArticlesSource
	}
	return nil
}

type Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Link        string `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
	Published   string `protobuf:"bytes,4,opt,name=published,proto3" json:"published,omitempty"`
	Authors     string `protobuf:"bytes,5,opt,name=authors,proto3" json:"authors,omitempty"`
}

func (x *Article) Reset() {
	*x = Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_crawler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_crawler_proto_rawDescGZIP(), []int{2}
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Article) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *Article) GetPublished() string {
	if x != nil {
		return x.Published
	}
	return ""
}

func (x *Article) GetAuthors() string {
	if x != nil {
		return x.Authors
	}
	return ""
}

type ArticlesSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Link        string `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
	FeedLink    string `protobuf:"bytes,4,opt,name=feed_link,json=feedLink,proto3" json:"feed_link,omitempty"`
	Image       string `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *ArticlesSource) Reset() {
	*x = ArticlesSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticlesSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticlesSource) ProtoMessage() {}

func (x *ArticlesSource) ProtoReflect() protoreflect.Message {
	mi := &file_crawler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticlesSource.ProtoReflect.Descriptor instead.
func (*ArticlesSource) Descriptor() ([]byte, []int) {
	return file_crawler_proto_rawDescGZIP(), []int{3}
}

func (x *ArticlesSource) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ArticlesSource) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ArticlesSource) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *ArticlesSource) GetFeedLink() string {
	if x != nil {
		return x.FeedLink
	}
	return ""
}

func (x *ArticlesSource) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type NewArticlesCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *NewArticlesCount) Reset() {
	*x = NewArticlesCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawler_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewArticlesCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewArticlesCount) ProtoMessage() {}

func (x *NewArticlesCount) ProtoReflect() protoreflect.Message {
	mi := &file_crawler_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewArticlesCount.ProtoReflect.Descriptor instead.
func (*NewArticlesCount) Descriptor() ([]byte, []int) {
	return file_crawler_proto_rawDescGZIP(), []int{4}
}

func (x *NewArticlesCount) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_crawler_proto protoreflect.FileDescriptor

var file_crawler_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf3, 0x02, 0x0a, 0x07, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x65, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x65, 0x65, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x61,
	0x77, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x61, 0x77, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x5f, 0x64, 0x69, 0x76, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x69, 0x76, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2f,
	0x0a, 0x13, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69,
	0x6e, 0x6b, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x10, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x6c, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x24, 0x0a, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52,
	0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x0f, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x73, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x0e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x73, 0x22, 0x8f, 0x01, 0x0a, 0x0e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69,
	0x6e, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x65, 0x65, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x10, 0x4e, 0x65, 0x77, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32,
	0x91, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x29, 0x0a, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x53, 0x53, 0x43, 0x72, 0x61,
	0x77, 0x6c, 0x65, 0x72, 0x12, 0x08, 0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x1a, 0x0b,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x2c, 0x0a,
	0x11, 0x54, 0x65, 0x73, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x72, 0x61, 0x77, 0x6c,
	0x65, 0x72, 0x12, 0x08, 0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x1a, 0x0b, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x05, 0x43,
	0x72, 0x61, 0x77, 0x6c, 0x12, 0x08, 0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x1a, 0x11,
	0x2e, 0x4e, 0x65, 0x77, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6e, 0x65, 0x69, 0x6c, 0x2d, 0x67, 0x6f, 0x2d, 0x70, 0x68, 0x61, 0x6e, 0x2f, 0x6e,
	0x65, 0x77, 0x73, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crawler_proto_rawDescOnce sync.Once
	file_crawler_proto_rawDescData = file_crawler_proto_rawDesc
)

func file_crawler_proto_rawDescGZIP() []byte {
	file_crawler_proto_rawDescOnce.Do(func() {
		file_crawler_proto_rawDescData = protoimpl.X.CompressGZIP(file_crawler_proto_rawDescData)
	})
	return file_crawler_proto_rawDescData
}

var file_crawler_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_crawler_proto_goTypes = []interface{}{
	(*Crawler)(nil),          // 0: Crawler
	(*TestResult)(nil),       // 1: TestResult
	(*Article)(nil),          // 2: Article
	(*ArticlesSource)(nil),   // 3: ArticlesSource
	(*NewArticlesCount)(nil), // 4: NewArticlesCount
}
var file_crawler_proto_depIdxs = []int32{
	2, // 0: TestResult.articles:type_name -> Article
	3, // 1: TestResult.articles_source:type_name -> ArticlesSource
	0, // 2: CrawlerService.TestRSSCrawler:input_type -> Crawler
	0, // 3: CrawlerService.TestCustomCrawler:input_type -> Crawler
	0, // 4: CrawlerService.Crawl:input_type -> Crawler
	1, // 5: CrawlerService.TestRSSCrawler:output_type -> TestResult
	1, // 6: CrawlerService.TestCustomCrawler:output_type -> TestResult
	4, // 7: CrawlerService.Crawl:output_type -> NewArticlesCount
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_crawler_proto_init() }
func file_crawler_proto_init() {
	if File_crawler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crawler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Crawler); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_crawler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_crawler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Article); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_crawler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticlesSource); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_crawler_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewArticlesCount); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_crawler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crawler_proto_goTypes,
		DependencyIndexes: file_crawler_proto_depIdxs,
		MessageInfos:      file_crawler_proto_msgTypes,
	}.Build()
	File_crawler_proto = out.File
	file_crawler_proto_rawDesc = nil
	file_crawler_proto_goTypes = nil
	file_crawler_proto_depIdxs = nil
}
