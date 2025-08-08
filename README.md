**Describe the bug**
The lsp seems to break when using struct methods in templates.

**To Reproduce**

- [full example repo](https://github.com/dontWatchMeCode/templ-issue-template/tree/struct-method-lps-issue)
- [example .templ file](https://raw.githubusercontent.com/dontWatchMeCode/templ-issue-template/refs/heads/struct-method-lps-issue/templates/example.templ)

**Expected behavior**
A clear and concise description of what you expected to happen.

**Screenshots**

struct example:

<img src="https://github.com/dontWatchMeCode/templ-issue-template/blob/struct-method-lps-issue/screenshots/struct_1.png?raw=true" width="500">

<img src="https://github.com/dontWatchMeCode/templ-issue-template/blob/struct-method-lps-issue/screenshots/struct_2.png?raw=true" width="500">

function example:

<img src="https://github.com/dontWatchMeCode/templ-issue-template/blob/struct-method-lps-issue/screenshots/function_1.png?raw=true" width="500">

<img src="https://github.com/dontWatchMeCode/templ-issue-template/blob/struct-method-lps-issue/screenshots/function_2.png?raw=true" width="500">

**Logs**

- [gopls](https://github.com/dontWatchMeCode/templ-issue-template/blob/struct-method-lps-issue/logs/gopls.log)
- [templ](https://github.com/dontWatchMeCode/templ-issue-template/blob/struct-method-lps-issue/logs/templ.log)

**`templ info` output**

```
(✓) os [ goos=linux goarch=amd64 ]
(✓) go [ location=/usr/bin/go version=go version go1.24.5 linux/amd64 ]
(✓) gopls [ location=/home/me/go/bin/gopls version=golang.org/x/tools/gopls v0.20.0 ]
(✓) templ [ location=/home/me/go/bin/templ version=v0.3.924 ]
```

**Desktop (please complete the following information):**

- OS: Linux fedora 6.15.9
- templ CLI version v0.3.924
- Go version go1.24.5
- `gopls` version v0.20.0
