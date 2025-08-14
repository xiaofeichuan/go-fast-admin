package consts

type ComponentType string

const (
	ComponentTypeInput          ComponentType = "Input"          //输入框
	ComponentTypeTextarea                     = "Textarea"       //文本域
	ComponentTypeSwitch                       = "Switch"         //开关
	ComponentTypeSelect                       = "Select"         //字典选择
	ComponentTypeImageUpload                  = "ImageUpload"    //图片上传
	ComponentTypeInputNumber                  = "InputNumber"    //数字输入框
	ComponentTypeDateTimePicker               = "DateTimePicker" //日期时间选择器
)
