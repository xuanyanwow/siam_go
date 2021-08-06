package core

// 实例化一个通过字符串映射函数切片的map
var eventByName = make(map[string][]func(interface{}))


const EventNameBeforeStartListen = "BeforeStartListen"
const EventNameOnShutDown = "OnShutDown"
const EventNameAfterRequest = "AfterRequest"


func RegisterEvent(name string, callback func(interface{})){
	// 通过名字查找事件列表
	list := eventByName[name]

	// 在列表切片中添加函数
	list = append(list, callback)

	// 将修改的事件列表切片保存回去
	eventByName[name] = list
}

// 调用事件
func CallEvent(name string, param ...interface{}) {
	// 通过名字找到事件列表
	list := eventByName[name]

	// 遍历这个事件的所有回调
	for _, callback := range list {
		// 传入参数调用回调
		callback(param)
	}
}

func HasEvent(name string) bool {
	list := eventByName[name]
	return len(list) > 0;
}