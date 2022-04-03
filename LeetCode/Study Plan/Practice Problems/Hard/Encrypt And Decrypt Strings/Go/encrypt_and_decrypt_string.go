type Encrypter struct {
    key_val map[string]string
    val_key map[string]map[string]bool
    dic_hash map[string][]string
    
}

func get_encrypted_word(key string, word map[string]string)string{
    str := ""
    for i:=0;i<len(key);i++{
        str+= word[string(key[i])]
    }
    return str
}

func Constructor(keys []byte, values []string, dic []string) Encrypter {
    key_val := make(map[string]string)
    val_key := make(map[string]map[string]bool)
    n := len(keys)
    for i:=0;i<n;i++{
        key_val[string(keys[i])] = values[i]
        if _,ok:=val_key[values[i]];!ok{
            val_key[values[i]] = make(map[string]bool)
        }
        val_key[values[i]][string(keys[i])] = true 
    }
    n = len(dic)
    dic_hash := make(map[string][]string)
    for i:=0;i<n;i++ {
        val := get_encrypted_word(dic[i], key_val)
        if _,ok:=dic_hash[val];!ok{
            dic_hash[val] = make([]string,0)
        }
        dic_hash[val] = append(dic_hash[val], dic[i])
    }
    return Encrypter{key_val, val_key, dic_hash}
}


func (this *Encrypter) Encrypt(word1 string) string {
    str := ""
    n := len(word1)
    for i:=0;i<n;i++{
        str += this.key_val[string(word1[i])]
    }
    return str
}


func (this *Encrypter) Decrypt(word2 string) int {
    return len(this.dic_hash[word2])
}


/**
 * Your Encrypter object will be instantiated and called as such:
 * obj := Constructor(keys, values, dictionary);
 * param_1 := obj.Encrypt(word1);
 * param_2 := obj.Decrypt(word2);
 */