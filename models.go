package main

type CommandModel struct {
	FilePath *string
	Interval *int
}

type ConfigurationModel struct {
	AccessId   string
	AccessKey  string
	MainDomain string
	SubDomains *[]SubDomainModel
}

type SubDomainModel struct {
	Type     string
	Name     string
	Interval int
}
