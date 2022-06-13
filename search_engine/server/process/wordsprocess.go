package process

import (
	"Search_Engines/search_engines/commom/static"
	"fmt"
	"regexp"

	"github.com/wangbin/jiebago"
)

var Wdp *WordsParticiple

func init() {
	Wdp = &WordsParticiple{}
	fmt.Println("加载自定义词典!!!1")
	Wdp.seg.LoadDictionary(static.LOAD_DICTIONARY)
	Wdp.simple_test()
}

type WordsParticiple struct {
	seg jiebago.Segmenter
}

func NewWordParticiple() *WordsParticiple {
	return Wdp
}

// RemovePunctuation 移除所有的标点符号
func (this *WordsParticiple) RemovePunctuation(str string) string {
	reg := regexp.MustCompile(`\p{P}+`)
	return reg.ReplaceAllString(str, "")
}

// RemoveSpace 移除所有的空格
func (this *WordsParticiple) RemoveSpace(str string) string {
	reg := regexp.MustCompile(`\s+`)
	return reg.ReplaceAllString(str, "")
}

/*
功能：实现分词并返回分词结果
实现：
1.调用api分词
2.map分词结果过滤
3.返回分词数据
参数：words:中文要分词的句子
返回值:[]string
*/
func (this *WordsParticiple) Participle(words string) (str []string) { //分词
	ch := Wdp.seg.CutForSearch(words, true)
	// str= make([]string, 1)
	cnt := make(map[string]int, len(ch))
	for word := range ch {
		word = Wdp.RemoveSpace(word)
		word = Wdp.RemovePunctuation(word)
		/*
		   当要查询的 key 不在 map 里，
		   带 comma 的用法会返回一个 bool 型变量提示 key 是否在 map 中；
		   而不带 comma 的语句则会返回一个 value 类型的零值。
		   如果 value 是 int 型就会返回 0，如果 value 是 string 类型，就会返回空字符串。
		   // 不带 comma 用法
		       age1 := ageMap["stefno"]
		       fmt.Println(age1)
		       // 带 comma 用法
		       age2, ok := ageMap["stefno"]
		       fmt.Println(age2, ok)
		*/
		_, ok := cnt[word]
		if len(word) > 0 && !ok {
			str = append(str, word)
		}
		cnt[word] = 1 //插入
	}
	return str
}

/*
功能：实现对敏感词过滤
参数：str1分词句子  str2敏感词
实现：
1.对str1，str2分词
2.遍历扫描str1分词结果中有没有str2中数据
3.返回结果
返回：
*/
func (this *WordsParticiple)Participle_Filter(str1 string,str2 string)(ret []string){
	str1s:=this.Participle(str1)
	str2s:=this.Participle(str2)
	for _,v1:=range str1s{
		sum:=0
		for _,v2:=range str2s{
			if v1==v2{
				sum=sum+1
			}
		}
		if sum==0{
			ret=append(ret,v1);
		}
	}
	return 
}
/*
功能：实现分词并返回分词结果
实现：
1.调用api分词
2.返回分词数据
参数：words:中文要分词的句子
返回值:[]string 分词结果不过滤
*/
func (this *WordsParticiple) Participle_No_Filtering(words string) (str []string) { //分词
	ch := Wdp.seg.CutForSearch(words, true)
	for word := range ch {
		word = Wdp.RemoveSpace(word)
		word = Wdp.RemovePunctuation(word)
		if len(word) > 0 {
			str = append(str, word)
		}
	}
	return str
}

func (this *WordsParticiple) simple_test() {
	//test_001
	str := "小明硕士毕业于中国科学院计算所，后在日本京都B站大学深造 小明硕士毕业于中国科学院计算所，后在日本京都B站大学深造"
	ch := this.Participle(str)
	if len(ch) != 18 {
		fmt.Println("WordsParticiple Error!!!!")
		fmt.Println(ch)
	} else {
		fmt.Println("加载自定义词典成功 ACEPT 001!!!!")
	}
}

/*
功能：完成url和文当内容分离
实现：
1.找到逗号
2.切片
3.去掉result两端的引号
参数：str：*string(url,caption) 以逗号分割url，文档内容的字符串
返回值：[]string 长度为2 0：为url 1：为文档
*/

func (this *WordsParticiple) Url_Caption(str *string) (result []string) {
	return
}
