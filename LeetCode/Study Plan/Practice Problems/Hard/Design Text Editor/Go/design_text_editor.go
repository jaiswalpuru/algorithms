package main

//--------------------------------brute force-----------------------------(won't work because chaning contents of a string on the go takes time)
type TextEditor struct {
	cur  int
	text string
}

func Constructor() TextEditor {
	return TextEditor{cur: 0, text: ""}
}

func (this *TextEditor) AddText(text string) {
	temp := this.text[:this.cur]
	temp += text
	this.text = temp + this.text[this.cur:]
	this.cur += len(text)
	// fmt.Println("add", this.text, this.cur)
}

func (this *TextEditor) DeleteText(k int) int {
	// fmt.Println("delete : ",this.cur, k,this.text)
	if this.cur == 0 {
		return 0
	}
	temp := this.text[this.cur:]
	temp_2 := this.text[:this.cur]
	d := 0
	if this.cur-k < 0 {
		d = this.cur
		this.cur = 0
	} else {
		d = k
		this.cur -= k
	}
	temp_2 = this.text[:this.cur]
	this.text = temp_2 + temp
	return d
}

func (this *TextEditor) CursorLeft(k int) string {
	// fmt.Println("left :",this.cur, k,this.text)
	if this.cur-k < 0 {
		this.cur = 0
	} else {
		this.cur -= k
	}
	if this.cur-10 < 0 {
		return this.text[:this.cur]
	} else {
		return this.text[this.cur-10 : this.cur]
	}
}

func (this *TextEditor) CursorRight(k int) string {
	if this.cur+k > len(this.text) {
		this.cur = len(this.text)
	} else {
		this.cur += k
	}
	if this.cur-10 < 0 {
		return this.text[:this.cur]
	} else {
		return this.text[this.cur-10 : this.cur]
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//---------------------------------------------------------------------------------------

//------------------------------------effficient approach--------------------------------
type TextEditor struct{ left, right []byte }

func Constructor() TextEditor { return TextEditor{left: make([]byte, 0), right: make([]byte, 0)} }

func (this *TextEditor) AddText(text string) {
	str := []byte(text)
	for i := 0; i < len(text); i++ {
		this.left = append(this.left, str[i])
	}
}

func (this *TextEditor) DeleteText(k int) int {
	cnt := 10
	for len(this.left) > 0 && k > 0 {
		this.left = this.left[:len(this.left)-1]
		cnt++
		k--
	}
	return cnt
}

func (this *TextEditor) CursorLeft(k int) string {
	for len(this.left) > 0 && k > 0 {
		char := this.left[len(this.left)-1]
		this.left = this.left[:len(this.left)-1]
		this.right = append(this.right, char)
		k--
	}
    res := shift(&this.left)
    return res
}

func (this *TextEditor) CursorRight(k int) string {
	for len(this.right) > 0 && k > 0 {
		char := this.right[len(this.right)-1]
		this.right = this.right[:len(this.right)-1]
		this.left = append(this.left, char)
		k--
	}
    res := shift(&this.left)
    return res
}

func shift(l *[]byte) string {
	min_cnt := 10
	str := []byte{}
	for len(*l) > 0 && min_cnt > 0 {
		char := (*l)[len(*l)-1]
		*l = (*l)[:len(*l)-1]
		str = append(str, char)
		min_cnt--
	}

	for i := 0; i < len(str)/2; i++ {
		str[i], str[len(str)-i-1] = str[len(str)-i-1], str[i]
	}

	for i := 0; i < len(str); i++ {
		*l = append(*l, str[i])
	}

	return string(str)
}

//---------------------------------------------------------------------------------------

/**
 * Your TextEditor object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddText(text);
 * param_2 := obj.DeleteText(k);
 * param_3 := obj.CursorLeft(k);
 * param_4 := obj.CursorRight(k);
 */
