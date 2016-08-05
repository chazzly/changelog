// Changelog
//
// Copyright (c) 2014 Sebastian Müller <info@sebastian-mueller.net>
//
// https://github.com/SebastianM/changelog
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package generator

import (
	"fmt"
	"../git"
	"strings"
	"regexp"
)

const (
	VERSION          = "## %s%s"
	PATCH_VERSION    = "### %s%s"
	LINK_ISSUE       = "[#%s](%s/issues/%s)"
	ISSUE            = "(#%s)"
	LINK_COMMIT      = "[%s](%s/commit/%s)"
	COMMIT           = "(%s)"
	EMPTY_COMPONENT  = "$$"
	PLAIN_HEADER_TPL = `<a name="%s"></a>\n%s (%s)\n\n`
	LINK_HEADER_TPL  = "%s (%s)\n\n"
)

type sections struct {
}

func GenerateNewChangelogContent(existingContent string, commits []*git.Commit, version string) (newContent string, err error) {
	fmt.Printf("Generating content for %d commits\n", len(commits))

	existingContent = strings.TrimPrefix(existingContent, "\n")
	existingContent = strings.TrimSuffix(existingContent, "\n")
	lines := strings.Split(existingContent, "\n")

	var header, oldContent string
	var headfound, firstVer bool
	for i := 0; i < len(lines); i++ {
		hm, _ := regexp.MatchString("====*", lines[i])
		vm, _ := regexp.MatchString("---*", lines[i])
		vs, _ := regexp.MatchString("\\s*\\d\\.\\d\\.\\d\\s*", lines[i])
		switch {
		case hm:
			headfound = true
			header += lines[i] + "\n"
		case !headfound:
			header += lines[i] + "\n"
		case vs, vm:
			firstVer = true
			oldContent += lines[i] + "\n"
		case !firstVer:
			header += lines[i] + "\n"
		default:
			oldContent += lines[i] + "\n"
		}
	}

	if header == "" {
		header = "CHANGELOG\n========\nList changes on a release by release basis.\n\n"
		// TODO: Add cookbook name to Header
	}

	var freshContent string = version + "\n------\n"
	for _,tcommit := range commits {
		var entry string
		entry = tcommit.Date + " - "+ tcommit.Author + ": " + tcommit.Subject + "\n"
		freshContent = freshContent + entry
	}

	newContent = header + freshContent + "\n" + oldContent
	return
}