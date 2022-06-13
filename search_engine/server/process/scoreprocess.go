package process

import "math"

/*
 用于打分算法处理器
*/
type ScorerPrcoss struct{}

/*
功能：实现TF-IDF打分
参数：
实现:
TF: 词频(Term Frequency)，表示词条在文档d中出现的频率，这个比较简单。
IDF可以由文档库中的总文档数（numDocs）除以包含该词条的文档数量（docFreq），再将得到的商取以10为底的对数得到，即：
返回：分数
*/
func (this *ScorerPrcoss) TF_IDF(tf uint32, numDocs uint64, docFreq uint64) (Score float64) {
	idf := math.Log10(float64(numDocs) / float64(docFreq+1))
	return float64(idf) * float64(tf)
}
