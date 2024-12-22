package urls

func GetDomain(url string) string {
    var domain string
    for index := range url {
        if url[index] == '/' && url[index-1] == '/' {
            index++
            for {
                if url[index] == '/' { break }
                domain += string(url[index])
                index++
            }
            break
        }
    }
    return domain
}

