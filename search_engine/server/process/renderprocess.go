package process

import (
	"fmt"
	"strings"
)

/*
 功能：负责数据渲染
*/
type Renderprocess struct{}

/*
功能：传入数据  分词数据 进行渲染
参数：strs查询的分词 docs文档集
实现：
都转为字节匹配
1.分词遍历
2.文档kmp匹配
返回：渲染后的数据集
优化方案：
AC自动机多模式匹配
这里没实现
*/
func (this *Renderprocess) Render_Data(keys []string,test string) (result string) {
	/*
	preTag := "<span style='color:red'>"
	posTag := "</span>"
	key := "年后"
	test := "年后号号"
	test = strings.ReplaceAll(test, key, fmt.Sprintf("%s%s%s", preTag, key, posTag))
	*/
	preTag := "<span style='color:red'>"
	posTag := "</span>"
	for _, key := range keys {
		test = strings.ReplaceAll(test, key, fmt.Sprintf("%s%s%s", preTag, key, posTag))
	}
	result=test;
	return
}

/*
const int N=1e6+10;
int n,m;
char s[N],t[N];
int ne[N],mch[N];
int main(){
	cin>>(s+1)>>(t+1);
	n=strlen(s+1);m=strlen(t+1);
	ne[0]=ne[1]=0;
	for(int i=2,j=0;i<=m;++i){
		while(j&&t[i]!=t[j+1])j=ne[j];
		if(t[i]==t[j+1])j++;ne[i]=j;
	}
	for(int i=1,j=0;i<=n;++i){
		while(j&&s[i]!=t[j+1])j=ne[j];
		if(s[i]==t[j+1])j++;
		if(j==m)mch[i]=1,j=ne[j];
	}
	for(int i=1;i<=n;++i)if(mch[i])cout<<i-m+1<<endl;
	for(int i=1;i<=m;++i)cout<<ne[i]<<" ";
	return 0;
}
*/

/*
功能：传入一个词 数据进行染红色
参数:(传入一个词)
实现：数据标识进行染色
返回：染色后的词
*/
func (this *Renderprocess) Render_Red(str string) string {
	return "<span style='color: #c00;'>" + str + "</span>"
}
