package cmds

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	cli "github.com/urfave/cli/v2"
	"github.com/waldirborbajr/glink/internal/util"
)

func Link() *cli.Command {
	return &cli.Command{
		Name:                   "link",
		Aliases:                []string{"l"},
		Usage:                  "Create symlinks",
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "force",
				Aliases: []string{"f"},
				Usage:   "force overwrite symlinkl",
			},
		},
		Action: func(cCtx *cli.Context) error {
			createSymlink()
			fmt.Println("symlinks created")
			return nil
		},
	}
}

// Create the Symlink for file or directory
func createSymlink() {
	userHomeDir := userHomeDir()

	linkSourcePath := sourceLinkPath()

	// TODO: implement logic to ignore files and directories
	// ignoreList(linkSourcePath)

	filesToLink := listFilestoLink(linkSourcePath)

	for _, file := range filesToLink {

		fileName := file.Name()

		// ignoring the .glink-ignore file
		if file.Name() == ".glink-ignore" || fileName == "glink-ignore" {
			continue
		}

		// Check if the file is a directory and has content
		if isDirectory(fileName) {
			if !hasContent(fileName) {
				continue
			}

			newUserHomeDirectory := userHomeDir + "/" + fileName

			if isTargetExists(newUserHomeDirectory) {

				newSourcePath := fileName

				pwd, _ := os.Getwd()
				filesIntoDiretory := listFilestoLink(pwd + "/" + newSourcePath)

				for _, fileIntoDirectory := range filesIntoDiretory {
					fileFromDiretory := fileIntoDirectory.Name()

					if !isTargetExists(newUserHomeDirectory + "/" + fileFromDiretory) {

						pwd, _ := os.Getwd()
						os.Chdir(newUserHomeDirectory)
						if err := makeSymlink("../"+filepath.Base(pwd)+"/"+fileName+"/"+fileFromDiretory, fileFromDiretory); err != nil {
							util.ExitWithError("Error creating symlink", err)
						}
						os.Chdir(pwd)
					}
				}
			} else {
				if err := makeSymlink(linkSourcePath+"/"+fileName, userHomeDir+"/"+fileName); err != nil {
					util.ExitWithError("Error creating symlink", err)
				}
			}
			continue

		} else {
			if !isTargetExists(userHomeDir + "/" + fileName) {
				if err := makeSymlink(linkSourcePath+"/"+fileName, userHomeDir+"/"+fileName); err != nil {
					util.ExitWithError("Error creating symlink", err)
				}
			}
		}
	}
}

// makeSymlink creates a symlink and return error
func makeSymlink(source string, target string) error {
	return os.Symlink(source, target)
}

// isDirectory determines if a file represented
// by `path` is a directory or not
func isDirectory(path string) bool {
	fileInfo, _ := os.Stat(path)

	return fileInfo.IsDir()
}

// hasContent determines if a directory has content to be created as a symlink
func hasContent(path string) bool {
	contents, err := os.ReadDir(path)
	if err != nil {
		return false
	}

	return len(contents) != 0
}

// Validate if diretory exists on target symlink
func isTargetExists(target string) bool {
	_, err := os.Stat(target)

	return !errors.Is(err, os.ErrNotExist)
}

// Get user home path
func userHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		util.ExitWithError("Unable to get $HOME directory", err)
	}

	return home
}

// Get current path
func sourceLinkPath() string {
	linkSourcePath, err := os.Getwd()
	if err != nil {
		util.ExitWithError("Unable to get current directory", err)
	}

	// linkSourcePath = filepath.Base(linkSourcePath)

	return linkSourcePath
}

// Get list of files and folder to create symlink
func listFilestoLink(sourcePath string) []fs.DirEntry {
	files, err := os.ReadDir(sourcePath)
	if err != nil {
		util.ExitWithError("Unable to get a list of files", err)
	}

	return files
}

// TODO: implement ignore list
// Do not create symlink from .glink-ignore
func ignoreList(sourcePath string) []string {
	_, err := os.Stat(sourcePath + "/.glink-ignore")

	if errors.Is(err, os.ErrNotExist) {
		return make([]string, 0)
	}

	ignoreFile := sourcePath + "/.glink-ignore"

	fileToIgnore, err := os.OpenFile(ignoreFile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		util.ExitWithError("Unable to read .glink-ignore file", err)
	}

	defer fileToIgnore.Close()

	content := bufio.NewReader(fileToIgnore)
	for {
		line, err := content.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			util.ExitWithError("read file line error: %v", err)
		}
		_ = line // GET the line string
	}

	return make([]string, 0)
}
