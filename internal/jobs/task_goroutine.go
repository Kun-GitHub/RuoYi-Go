package jobs

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

// TaskDemo 实现 Task 接口的一个示例
type TaskGoroutine struct {
	logger      *zap.Logger
	pageNumChan chan int
}

// NewTaskDemo
func NewTaskGoroutine(l *zap.Logger) *TaskGoroutine {
	pageNumChan := make(chan int, 1)
	pageNumChan <- 1

	//无缓冲的chan
	//pageNumChan := make(chan int)
	//// 初始化pageNum为0
	//go func() {
	//	pageNumChan <- 1
	//}()

	return &TaskGoroutine{
		logger:      l,
		pageNumChan: pageNumChan,
	}
}

func (this *TaskGoroutine) Run() {
	//this.logger.Info("TaskGoroutine is running")
	select {
	case pageNum := <-this.pageNumChan:
		this.logger.Info(fmt.Sprintf("当前页码为：%d", pageNum))
		go this.fetchAndInsertData(pageNum)
	default:
		this.logger.Info(fmt.Sprintf("default"))
		//defaultPageNum := 1
		//this.fetchAndInsertData(defaultPageNum)
	}
}

func (this *TaskGoroutine) fetchAndInsertData(pageNum int) {
	// 获取数据
	// 创建一个客户端请求对象
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	t := time.Now()
	dateStr := t.Format("2006-01-02")

	// 设置请求 URL
	req.SetRequestURI("http://deal.ggzy.gov.cn/ds/deal/dealList_find.jsp?" +
		"TIMEBEGIN_SHOW=" + dateStr + "&TIMEEND_SHOW=" + dateStr +
		"&TIMEBEGIN=" + dateStr + "&TIMEEND=" + dateStr +
		"&SOURCE_TYPE=1&DEAL_TIME=02&DEAL_CLASSIFY=01&DEAL_STAGE=0101&DEAL_PROVINCE=530000&DEAL_CITY=0&DEAL_PLATFORM=0&BID_PLATFORM=0&DEAL_TRADE=0&isShowAll=1&PAGENUMBER=" + strconv.Itoa(pageNum) + "&FINDTXT=")
	// 设置请求方法为 POST
	req.Header.SetMethod("POST")

	// 创建一个响应对象
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// 发起请求
	if err := fasthttp.Do(req, resp); err != nil {
		this.pageNumChan <- pageNum
		this.logger.Error("An error occurred while making the request: %v", zap.Error(err))
		return
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		this.pageNumChan <- pageNum
		this.logger.Error(fmt.Sprintf("Request failed with status code: %d", resp.StatusCode()))
		return
	}

	// 获取响应体
	body := string(resp.Body())
	// 输出结果
	body = strings.ReplaceAll(body, "&nbsp;", "")
	body = strings.ReplaceAll(body, "\n", "")
	body = strings.ReplaceAll(body, "\r", "")
	//this.logger.Info(fmt.Sprintf("Response body: %s", body))

	var response Response
	err := json.Unmarshal(resp.Body(), &response)
	if err != nil {
		this.pageNumChan <- pageNum
		this.logger.Error("Error unmarshalling JSON: %v", zap.Error(err))
		return
	}
	if response.Success && response.TtlRow > 0 {
		this.logger.Info(fmt.Sprintf("TtlPage: %d", response.TtlPage))
		if pageNum < response.TtlPage {
			this.pageNumChan <- pageNum + 1
		} else {
			this.logger.Info("数据抓取完毕")
			this.pageNumChan <- 1
		}
	} else {
		this.pageNumChan <- pageNum
	}
}

type Response struct {
	TtlPage     int           `json:"ttlpage"`
	TtlRow      int           `json:"ttlrow"`
	Data        []interface{} `json:"data"`
	UseTime     int           `json:"usetime"`
	CurrentPage int           `json:"currentpage"`
	Success     bool          `json:"success"`
	Error       string        `json:"error"`
}
