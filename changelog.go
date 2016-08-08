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

package main

import (
	"./commands"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "cars-changelog"
	app.Usage = "Cars.com cookbook changelog generator, based on changelog generator by Sebastian Müller"
	app.Authors = []cli.Author{ { "Chaz Ruhl", "cruhl@cars.com"}, {"Sebastian Müller", "info@sebastian-mueller.net"}}
	app.Version = "0.1.0"
	app.Commands = Commands()
	app.Run(os.Args)
}

func Commands() (comms []cli.Command) {
	comms = append(commands.Convert(), commands.Generate())
	return
}
