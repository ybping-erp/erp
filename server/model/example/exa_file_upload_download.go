package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type ExaFileUploadAndDownload struct {
	global.GVA_MODEL
	Name      string `json:"name" gorm:"comment:文件名"` // 文件名
	Url       string `json:"url" gorm:"comment:文件地址"` // 文件地址
	Tag       string `json:"tag" gorm:"comment:文件标签"` // 文件标签
	Key       string `json:"key" gorm:"comment:编号"`   // 编号
	Size      int64  `json:"size" gorm:"comment:编号"`  // 文件大小 in Byte
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

func (ExaFileUploadAndDownload) TableName() string {
	return "exa_file_upload_and_downloads"
}
