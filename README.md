# sensitive-word-detect
sensitive words detection

Golang version of [observerss/textfilter](https://github.com/observerss/textfilter)

+ toy work for ByteDance backend trainingcamp

### Example
```golang
package utils

import (
	"github.com/Karshilov/sensitive-word-detect/automaton"
	"github.com/Karshilov/sensitive-word-detect/utils"
)

var SensitiveWordsFilter = func() automaton.ACAutomaton {
	keywords, _ := utils.GetKeywords()
	ac := automaton.ACAutomaton{}
	ac.Reserve(1000000)
	for _, v := range keywords {
		ac.Insert(v)
	}
	ac.Build()
	return ac
}()
```
