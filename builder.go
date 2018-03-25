package mysqlib

// Instance 构建器实例
type Instance struct {
	options    *Options              //配置
	modelCache map[string]*modelInfo //模型结构缓存
}

// Options 构建器实例配置选项
type Options struct {
	TagName           string //标记名
	TableNameField    string //表名字段名
	DisableModelCache bool   //禁用模型缓存（默认开启）
}

// New 实例化
func New(option ...*Options) *Instance {
	//实例化配置
	var opt Options
	//如果有指定配置
	if len(option) > 0 {
		opt = *option[0]
	}
	//默认配置
	if opt.TagName == "" {
		opt.TagName = "sql"
	}
	if opt.TableNameField == "" {
		opt.TableNameField = "tableName"
	}

	//创建构建器实例
	var instance Instance
	instance.options = &opt

	//如果没有禁用模型缓存
	if opt.DisableModelCache == false {
		//初始化模型缓存
		instance.modelCache = make(map[string]*modelInfo)
	}

	return &instance
}

// Model 通过传入的模型创建会话实例Session
func (instance *Instance) Model(m interface{}) *Session {
	//创建会话
	var sess Session
	sess.builder = instance   //存入构建器指针
	sess.modelValue.Value = m //存入模型实例
	return &sess
}
