package main

type DomainWhitelist struct {
	Whitelist map[string][]string
}

func (dw *DomainWhitelist) Get(user string) []string {
	return dw.Whitelist[user]
}

func (dw *DomainWhitelist) Set(user string, regex string) {
	if _, ok := dw.Whitelist[user]; !ok {
		dw.Whitelist[user] = make([]string, 0)
	}
	dw.Whitelist[user] = append(dw.Whitelist[user], regex)
}

func NewDomainWhitelist() *DomainWhitelist {
	dw := new(DomainWhitelist)
	dw.Whitelist = make(map[string][]string)

	return dw
}
