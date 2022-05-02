// @Author ljn 2022/4/29 20:40:00
package common

import (
	"fmt"
	"strings"
	"testing"
)

func similar(a, b string) int {
	if a[:4] == b[:4] {
		return 0
	}
	if a[:4] < b[:4] {
		return 1
	}
	return -1
}

func test(html string) string {
	html = strings.Replace(html, `id="`, `id="toc_`, -1)
	hs := regToc.FindAllString(html, -1)
	if len(hs) > 1 {
		sb := strings.Builder{}
		sb.WriteString(`<div class="toc"><ul>`)
		level := 0
		for i := 0; i < len(hs)-1; i++ {
			fg := similar(hs[i], hs[i+1])
			if fg == 0 {
				sb.WriteString(regH.ReplaceAllString(hs[i], `<li><a href="#$1">$2</a></li>`))
			} else if fg == 1 {
				level++
				sb.WriteString(regH.ReplaceAllString(hs[i], `<li><a href="#$1">$2</a><ul>`))
			} else {
				level--
				sb.WriteString(regH.ReplaceAllString(hs[i], `<li><a href="#$1">$2</a></li></ul></li>`))
			}
		}
		fg := similar(hs[len(hs)-2], hs[len(hs)-1])
		if fg == 0 {
			sb.WriteString(regH.ReplaceAllString(hs[len(hs)-1], `<li><a href="#$1">$2</a></li>`))
		} else if fg == 1 {
			level++
			sb.WriteString(regH.ReplaceAllString(hs[len(hs)-1], `<li><a href="#$1">$2</a><ul>`))
		} else {
			level--
			sb.WriteString(regH.ReplaceAllString(hs[len(hs)-1], `<li><a href="#$1">$2</a></li></ul></li>`))
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
		sb.WriteString(regH.ReplaceAllString(hs[0], `<li><a href="#$1">$2</a></li>`))
		sb.WriteString(`</ul></div>`)
		return sb.String() + html
	}
	return html
}

func TestGetTocHtml(t *testing.T) {

	html := `<h1><a id="ljn"></a>wj</h1>`
	fmt.Println(test(html))

}
