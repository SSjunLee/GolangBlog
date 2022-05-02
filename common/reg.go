// @Author ljn 2022/4/29 17:22:00
package common

import (
	"regexp"
	"strings"
)

var RegImg = regexp.MustCompile(`<img src="([^" ]+)" alt="([^" ]*)"\s?\/?>`)
var RegToc = regexp.MustCompile("<h[1-6]>.*?</h[1-6]>")
var RegH = regexp.MustCompile(`<h[1-6]><a id="(.*?)"></a>(.*?)</h[1-6]>`)

func GetTocHtml(html string) string {
	html = strings.Replace(html, `id="`, `id="toc_`, -1)
	hs := RegToc.FindAllString(html, -1)
	if len(hs) > 1 {
		sb := strings.Builder{}
		sb.WriteString(`<div class="toc"><ul>`)
		level := 0
		for i := 0; i < len(hs)-1; i++ {
			fg := similar(hs[i], hs[i+1])
			if fg == 0 {
				sb.WriteString(RegH.ReplaceAllString(hs[i], `<li><a href="#$1">$2</a></li>`))
			} else if fg == 1 {
				level++
				sb.WriteString(RegH.ReplaceAllString(hs[i], `<li><a href="#$1">$2</a><ul>`))
			} else {
				level--
				sb.WriteString(RegH.ReplaceAllString(hs[i], `<li><a href="#$1">$2</a></li></ul></li>`))
			}
		}
		fg := similar(hs[len(hs)-2], hs[len(hs)-1])
		if fg == 0 {
			sb.WriteString(RegH.ReplaceAllString(hs[len(hs)-1], `<li><a href="#$1">$2</a></li>`))
		} else if fg == 1 {
			level++
			sb.WriteString(RegH.ReplaceAllString(hs[len(hs)-1], `<li><a href="#$1">$2</a><ul>`))
		} else {
			level--
			sb.WriteString(RegH.ReplaceAllString(hs[len(hs)-1], `<li><a href="#$1">$2</a></li></ul></li>`))
		}
		for level > 0 {
			sb.WriteString(`</ul></li>`)
			level--
		}
		sb.WriteString(`</ul></div>`)
		return sb.String() + html
	}
	if len(hs) == 1 {
		sb := strings.Builder{}
		sb.WriteString(`<div class="toc"><ul>`)
		sb.WriteString(RegH.ReplaceAllString(hs[0], `<li><a href="#$1">$2</a></li>`))
		sb.WriteString(`</ul></div>`)
		return sb.String() + html
	}
	return html
}

func similar(a, b string) int {
	if a[:4] == b[:4] {
		return 0
	}
	if a[:4] < b[:4] {
		return 1
	}
	return -1
}
