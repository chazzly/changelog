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

package commands

import (
	"fmt"
	"github.com/urfave/cli"
	"strings"
	"io/ioutil"
	"os"
	"regexp"
	"errors"
)

const IAMCONVERTED = "<!-- Converted -->"

func Convert() []cli.Command {
	return []cli.Command{
		{
			Name:        "convert",
			Usage:        "Convert old style changelog to new format",
			Action:        convert,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "input, i", Value: "CHANGELOG.md", Usage: "Name of old change log to convert"},
				cli.StringFlag{Name: "file, f", Value: "CHANGELOG.md", Usage: "Which file to read the current changelog from and prepend the new changelog's contents to"},
			},
		},
	}
}


func convert(c *cli.Context) {
	// Load current
	var existingContent string
	contentBytes, err := ioutil.ReadFile(c.String("input"))
	if err != nil {
		fmt.Println("ERROR reading file to convert: " + c.String("input") + " - " + err.Error())
		os.Exit(1)
	}
	existingContent = string(contentBytes)

	// parse into commits - using existing Commit struct
	entries, header, err := ParseOldLog(existingContent)
	if err != nil {
		fmt.Println("ERROR: File to convert: " + c.String("file") + " " + err.Error())
		os.Exit(1)
	}

	// send to write the log
	writesConvertedChangelog(c.String("file"), header, entries, c)
}

type OldEntry struct {
	Version	string
	Author	string
	Body	string
//	Date	string // TODO:  Do not know if any have a date, but in case.
}

func ParseOldLog(oldContent string) (entries []OldEntry, header string, err error) {
	if oldContent == "" {
		return nil, "", errors.New("is empty")
	}

	oldContent = strings.TrimPrefix(oldContent, "\n")
	oldContent = strings.TrimSuffix(oldContent, "\n")
	lines := strings.Split(oldContent, "\n")

	if lines[len(lines)-1] == IAMCONVERTED {
		return nil, "", errors.New("has already been converted")
	}

	newEntry := OldEntry{}
	var headfound,firstVer bool
	for i := 0 ; i < len(lines) ; i++ {
		bd, _ := regexp.MatchString("^(- )+-*$", lines[i])
		hm, _ := regexp.MatchString("====*", lines[i])
		vm, _ := regexp.MatchString("^---*", lines[i])
		vs, _ := regexp.MatchString("^#*\\s*\\d+\\.\\d+\\.\\d+\\s*", lines[i])
		lines[i] = strings.Trim(lines[i], "- *#")  // Remove unnecessary leading dashes, asterisks, spaces, & hashes
		ao, _ := regexp.MatchString("^\\(.*\\)$", lines[i])
		bl, _ := regexp.MatchString("^\\s*$", lines[i])
		ap, _ := regexp.MatchString("^\\(.*\\)", lines[i])
		switch {
		case hm:
			headfound = true
			header += lines[i] + "\n"
		case !headfound:
			header += lines[i] + "\n"
		case bl, vm, bd:
		case vs:
			if newEntry.Version != "" {
				entries = append(entries, newEntry)
				newEntry = OldEntry{}
			}
			lines[i] = strings.Trim(lines[i], ":")
			newEntry.Version= strings.Replace(lines[i], "\n", "", 0)
			firstVer = true
			fmt.Println(newEntry.Version)
		case !firstVer:
			header += lines[i] + "\n"
		case ao:
			newEntry.Author = strings.Trim(lines[i], "()")
		case newEntry.Version != "":
			var ln []string
			if ap {
				ln = strings.SplitN(strings.Trim(lines[i], "("), ")", 2)
			} else {
				ln = strings.SplitN(lines[i], " - ", 2)
			}
			// TODO: need to account for multiple changes per entry with multiple authors and multi-line body
			fmt.Println(ln)
			switch {
			case len(ln) == 1:
				newEntryBodyadd(&newEntry, ln[0])
			case newEntry.Author == "":
				newEntry.Author = ln[0]
				newEntryBodyadd(&newEntry, ln[1])
			default:
				newEntryBodyadd(&newEntry, ln[0] + " - " + ln[1])
			}
		default:
		}
	}
	entries = append(entries, newEntry)
	return
}


func newEntryBodyadd(b *OldEntry, t string) *OldEntry{
	if b.Body == "" {
		b.Body = t
	} else {
		b.Body += "\n" + t
	}
	return b
}

func writesConvertedChangelog(filename, header string, entries []OldEntry, c *cli.Context) {
	fmt.Printf("Parsed %d old entries\n", len(entries))

	_, err := os.Stat(filename)

	if err != nil && !os.IsNotExist(err) {
		fmt.Println("ERROR getting file stats: " + filename + " - " + err.Error())
		os.Exit(1)
	}

	// create changelog file if not exists
	if os.IsNotExist(err) {
		_, err = os.Create(filename)

		if err != nil {
			fmt.Println("ERROR creating file " + filename + " - " + err.Error())
			os.Exit(1)
		}
	}

	newContent, err := GenerateConvertedChangelogContent(header, entries)
	err = ioutil.WriteFile(filename, []byte(newContent), 0644)

	if err != nil {
		fmt.Println("ERROR writing new file content " + filename + " - " + err.Error())
		os.Exit(1)
	}
}

func GenerateConvertedChangelogContent(header string, entries []OldEntry) (newContent string, err error) {
	fmt.Printf("Converting %d Change log entries\n", len(entries))

	if header == "" {
		header = "CHANGELOG\n========\nList changes on a release by release basis.\n"
	}
	// TODO: Add cookbook name to Header

	var entryContent string
	for i := len(entries); i > 0; i-- {
		tentry := entries[i-1]
		if tentry.Body == "" { continue } // if body is empty, skip that version entirely
		if tentry.Author == "" { tentry.Author = "Unknown"}
		entryContent += tentry.Version
		entryContent += "\n------\n"
		entryContent += tentry.Author + " - " + tentry.Body + "\n\n"
	}

	newContent = header + "\n" + entryContent + "\n" + IAMCONVERTED
	return
}