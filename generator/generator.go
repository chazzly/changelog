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

	newContent = "HIIII" + existingContent
	return
}
