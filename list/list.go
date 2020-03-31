package list

func Contains(target interface{}, values ...interface{}) bool {
    for _, v := range values {
        if v == target {
            return true
        }
    }
    return false
}