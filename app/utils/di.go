package utils


var di = make(map[string]interface{})

const DiNameDefaultDb = "DefaultDb";

func RegisterDi(name string, value interface{}){
    di[name] = value
}

func GetDi(name string) interface{} {
    return di[name]
}