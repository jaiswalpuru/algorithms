package main

type Codec struct {
    hash_map map[int]string
    i int
}


func Constructor() Codec {
    return Codec{hash_map:make(map[int]string), i:0 }
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	this.i++
    this.hash_map[this.i] = longUrl
    return "https://tinyurl.com/" + strconv.Itoa(this.i)
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
    v := int(shortUrl[len(shortUrl)-1]-'0')
    return this.hash_map[v]
}

func main(){}