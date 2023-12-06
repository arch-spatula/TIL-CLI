package gitignoreWriter

import (
	"fmt"
	"os"
)

func WriteGitIgnore() {
	gitignoreFileName := ".gitignore"

	if _, err := os.Stat(gitignoreFileName); os.IsNotExist(err) {
		gitignoreFile, err := os.Create(gitignoreFileName)
		if err != nil {
			fmt.Printf("Unable to write file: %v\n", err)
		}
		defer gitignoreFile.Close()
		fmt.Fprintln(gitignoreFile, string(`# Ignore all
*

# Unignore all with extensions
!*.*

# Unignore all dirs
!*/

### Above combination will ignore all files without extension ###

# Ignore files with extension .class & .sm
*.class
*.sm

# Ignore bin dir
bin/
# or
*/bin/*

# Unignore all .jar in bin dir
!*/bin/*.jar

# Ignore all library.jar in bin dir
*/bin/library.jar

# Ignore a file with extension
relative/path/to/dir/filename.extension

# Ignore a file without extension
template.md
til-config.json`))

		fmt.Println(gitignoreFileName, "을 만들어두겠습니다.")
	} else {
		fmt.Println(gitignoreFileName, "이 이미 만들어졌습니다.")
	}

}
