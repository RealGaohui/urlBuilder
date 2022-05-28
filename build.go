package urlBuilder

import (
	"regexp"
	"strings"
)

var (
	charSet = "&#+=/? "
	defaultChar = "koko"
	url      string
	defaultURL  = "your base or parameters invalid"
	defaultBuilder = Builder{url: &defaultURL}
)

type Builder struct {
	url  *string
	args []*string
	char map[string]string
}

type builder interface {
	SetBase(base string) *Builder
	SetPath(path string) *Builder
	SetParameter(key, value string) *Builder
	ToString() string
	valid(para string) bool
	generateParameter(args []string)
	replaceSpecialCharacter(arg string) string
}

func URLBuilder() *Builder {
	return &Builder{
		url: &url,
		args: []*string{},
		char: registry(),
	}
}

func (b *Builder) SetBase(base string) *Builder {
	b.args = append(b.args, &base)
	if b.valid(base) {
		return &Builder{
			url: func() *string {
				tmp := make([]string, 5)
				for _, val := range b.args{
					tmp = append(tmp, *val)
				}
				url = strings.Join(tmp, "")
				return &url
			}(),
			args: b.args,
			char: b.char,
		}
	}else {
		return &defaultBuilder
	}
}

func (b *Builder) SetPath(path string) *Builder {
	b.args = append(b.args, &path)
		return &Builder{
			url: func() *string {
				tmp := make([]string, 5)
				for _, val := range b.args {
					tmp = append(tmp, *val)
				}
				url = strings.Join(tmp, "")
				return &url
			}(),
			args: b.args,
			char: b.char,
		}
}

func (b *Builder)SetParameter(args ...string) *Builder{
	format := args[:]
	b.generateParameter(format)
	var builder strings.Builder
	if len(args)%2 != 0{
		return &defaultBuilder
	}
	tmp := []string{}
	for i:=0; i<len(format); i++{
		if i%2 == 0 {
			arg := format[i] + "=" + format[i+1]
			tmp = append(tmp, arg)
		}
	}
	param := strings.Join(tmp, "&")
	if strings.Contains(*b.url, "?"){
		return &Builder{
			url: func() *string {
				builder.WriteString(*b.url)
				builder.WriteString("&")
				builder.WriteString(param)
				url = builder.String()
				return &url
			}(),
		}
	}else {
		return &Builder{
			url: func() *string {
				builder.WriteString(*b.url)
				builder.WriteString("?")
				builder.WriteString(param)
				url = builder.String()
				return &url
			}(),
		}
	}
}

func (b *Builder) valid(para string) bool {
	reg1 := regexp.MustCompile("^/[0-9a-zA-Z]+$")
	reg2 := regexp.MustCompile("^(https|http|ftp|rtsp|mms)?://[0-9a-z.:]*$")
	if reg1.MatchString(para) || reg2.MatchString(para) {
		return true
	} else {
		return false
	}
}

func (b *Builder)generateParameter(args []string) {
	for index, _ := range args {
		if strings.ContainsAny(args[index], charSet) {
			args[index] = b.replaceSpecialCharacter(args[index])
		}
	}
}

func (b*Builder)replaceSpecialCharacter(arg string) string {
	for key, val := range b.char{
		if strings.Contains(arg, key){
			arg = strings.Replace(arg, key, val, -1)
		} else {
			arg = strings.Replace(arg, key, defaultChar, -1)
		}
	}
	return arg
}

func (b *Builder) ToString() string {
	return *b.url
}

func registry() map[string]string  {
	char := make(map[string]string)
	char["&"] = "%26"
	char["+"] = "%2B"
	char[" "] = "%20"
	char["/"] = "%2F"
	char["?"] = "%3F"
	char["#"] = "%23"
	char["="] = "%3D"
	return char
}

