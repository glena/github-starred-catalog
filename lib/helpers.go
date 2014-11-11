package lib

func ArrayPush(array []interface{}, element interface{}) []interface{} {

    newArray := make([]interface{}, len(array)+1)

    for i:=range(array) {
      newArray[i] = array[i]
    }

    newArray[len(newArray)-1] = element
    return newArray
}
