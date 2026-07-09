package model

type SysCache struct {
	CacheName  string `json:"cacheName"`
	CacheKey   string `json:"cacheKey,omitempty"`
	CacheValue string `json:"cacheValue,omitempty"`
	Remark     string `json:"remark,omitempty"`
}

type CacheInfoResult struct {
	Info         map[string]string `json:"info"`
	DbSize       int64             `json:"dbSize"`
	CommandStats []CommandStat     `json:"commandStats"`
}

type CommandStat struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
