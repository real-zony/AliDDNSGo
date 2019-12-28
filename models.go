package main

type CommandModel struct {
	FilePath *string
	Interval *int
}

type ConfigurationModel struct {
	AccessId   string            // 阿里云的 Access Id
	AccessKey  string            // 阿里云的 Access Key
	MainDomain string            // 需要更新的主域名，例如 sample.com
	SubDomains *[]SubDomainModel // 需要更新的具体子域名。
}

type SubDomainModel struct {
	Type     string // 子域名记录的类型。
	Name     string `json:"SubDomain"` // 子域名的名称，例如 sub1.sample.com
	Interval int64  // 子域名记录的 TTL 时间，单位是秒。
}
