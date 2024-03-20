// Copyright 2023-2033 Ryan Du <duruyao@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package arg

import (
	"flag"
	"github.com/duruyao/gotest/conf"
)

func init() {
	flag.StringVar(&a.dir, "dir", conf.DefaultDir, "Directory of dataset")
	flag.StringVar(&a.host, "host", conf.DefaultHost, "Host address to listen")
	flag.BoolVar(&a.wantHelp, "h", false, "Display this help message")
	flag.BoolVar(&a.wantHelp, "help", false, "Display this help message")
	flag.BoolVar(&a.wantVersion, "v", false, "Print version information and quit")
	flag.BoolVar(&a.wantVersion, "version", false, "Print version information and quit")
}
