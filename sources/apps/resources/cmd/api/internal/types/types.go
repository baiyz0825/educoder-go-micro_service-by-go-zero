// Code generated by goctl. DO NOT EDIT.
package types

type File struct {
	ID            int64  `json:"id"`
	UUID          int64  `json:"uuid"`
	Name          string `json:"name"`
	ObfuscateName string `json:"obfuscateName"`
	Size          int64  `json:"size"`
	Owner         int64  `json:"owner"`
	Status        int64  `json:"status"`
	FileType      int64  `json:"fileType"`
	Class         int64  `json:"class"`
	Suffix        string `json:"suffix"`
	DownloadAllow int64  `json:"downloadAllow"`
	Link          string `json:"link"`
	FilePoster    string `json:"filePoster"`
	CreateTime    int64  `json:"createTime"`
	UpdateTime    int64  `json:"updateTime"`
}

type OnlineText struct {
	ID         int64  `json:"id"`
	UUID       int64  `json:"uuid"`
	TypeSuffix int64  `json:"typeSuffix"`
	Owner      int64  `json:"owner"`
	Content    string `json:"content"`
	ClassID    int64  `json:"classId"`
	Permission int64  `json:"permission"`
	TextName   string `json:"textName"`
	TextPoster string `json:"textPoster"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
}

type Count struct {
	UID         int64 `json:"uId"`
	FileNum     int64 `json:"fileNum"`
	VideoNum    int64 `json:"videoNum"`
	PicNum      int64 `json:"picNum"`
	StorageSize int64 `json:"storageSize"`
	CreateTime  int64 `json:"createTime"`
	UpdateTime  int64 `json:"updateTime"`
}

type ResComment struct {
	ID         int64  `json:"id"`
	Owner      int64  `json:"owner"`
	ResourceID int64  `json:"resourceId"`
	Content    string `json:"content"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
}

type MenuItem struct {
	ClassID          int64      `json:"classId"`
	ClassParentID    int64      `json:"classParentId"`
	ClassName        string     `json:"className"`
	ClassResourceNum int64      `json:"classResourceNum"`
	Children         []MenuItem `json:"children"`
}

type ClassificationTreeMenuResp struct {
	Classifications []MenuItem `json:"classifications"`
}

type SearchClassificationSubDataReq struct {
	ClassificationID int64  `json:"classificationID,optional" validate:"gte=0"`
	IsUser           bool   `json:"isUser"`
	KeyWord          string `json:"keyWord,optional"`
	ResType          int64  `json:"resType,optional" validate:"gte=0"`
	Page             int64  `json:"page" validate:"required,gt=0"`
	Limit            int64  `json:"limit" validate:"required,gt=0"`
}

type SearchClassificationSubDataResp struct {
	Files      []File       `json:"files"`
	FilesTotal int64        `json:"filesTotal"`
	OnlineText []OnlineText `json:"onlineText"`
	TextsTotal int64        `json:"textsTotal"`
}

type GetCountUiDResp struct {
	UserFileCount Count `json:"userFileCount"`
}

type UploadFileReq struct {
	Name          string `form:"name" validate:"required"`
	Class         int64  `form:"class" validate:"required" validate:"required,gt=0"`
	DownloadAllow int64  `form:"downloadAllow" validate:"required" validate:"required,gte=0,lte=1"`
}

type DelFileReq struct {
	Id int64 `form:"id" validate:"required,gt=0"`
}

type SearchFileConditionReq struct {
	Page     int64  `json:"page" validate:"required,gt=0"`
	Limit    int64  `json:"limit" validate:"required,gt=0"`
	Name     string `json:"name"`
	Owner    int64  `json:"owner" validate:"gt=0"`
	FileType int64  `json:"fileType" validate:"gte=0,lte=3"`
	Class    int64  `json:"class"`
}

type SearchFileConditionResp struct {
	Files []File `json:"files"`
}

type DownLoadFileReq struct {
	ResourceFileId int64 `form:"resourceFileId" validate:"required,gte=0"`
}

type FileResInfoReq struct {
	FileResId int64 `form:"fileResId" validate:"required,gte=0"`
}

type FileResInfoResp struct {
	File File `json:"file"`
}

type UploadTextReq struct {
	TypeSuffix int64  `form:"typeSuffix" validate:"gte=0"`
	Content    string `form:"content" validate:"required"`
	ClassID    int64  `form:"classId" validate:"required,gte=0"`
	Permission int64  `form:"permission" validate:"gte=0,lte=1"`
	TextName   string `form:"textName" validate:"required"`
}

type DelTextReq struct {
	Id int64 `form:"id" validate:"required"`
}

type SearchOnlineConditionTextReq struct {
	Page       int64 `json:"page" validate:"required,gt=0"`
	Limit      int64 `json:"limit" validate:"required,gt=0"`
	Owner      int64 `json:"owner,optional" validate:"omitempty,gt=0"`
	ClassID    int64 `json:"classId,optional" validate:"omitempty,gt=0"`
	Permission int64 `json:"permission" validate:"gte=0,lte=1"`
}

type SearchOnlineTextConditionResp struct {
	OnlineText []OnlineText `json:"onlineText"`
}

type TextResInfoReq struct {
	TextResId int64 `form:"textResId" validate:"required,gte=0"`
}

type TextResInfoResp struct {
	OnlineText OnlineText `json:"onlineText"`
}

type AddResCommentReq struct {
	ResourceID int64  `json:"resourceId" validate:"required,gt=0"`
	Content    string `json:"content" validate:"required"`
}

type DelCommentReq struct {
	ID int64 `form:"id" validate:"required,gt=0"`
}

type GetCommentByIdReq struct {
	ID int64 `form:"id" validate:"required,gt=0"`
}

type GetCommentByIdResp struct {
	Comment ResComment `json:"comment"`
}

type ResCommentByUserOrResIdReq struct {
	Page       int64 `json:"page" validate:"required,gt=0"`
	Limit      int64 `json:"limit" validate:"required,gt=0"`
	ResourceID int64 `json:"resourceId" validate:"required,gt=0"`
}

type ResCommentByUserOrResIdResp struct {
	Comments []ResComment `json:"comments"`
}
