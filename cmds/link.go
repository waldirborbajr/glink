package cmds

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"

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

func createSymlink() {
	userHomeDir := userHomeDir()

	linkSourcePath := sourceLinkPath()

	ignoreList(linkSourcePath)

	filesToLink := listFilestoLink(linkSourcePath)

	for _, file := range filesToLink {
		fmt.Println(file)

		// ignoring the .glink-ignore file
		if file.Name() == ".glink-ignore" {
			continue
		}

		// If it is a directory, must validate if exists on target
		// Existing must enter into and create symlink from the content inside
		// Not existing, create directory as symlink
		if isDirectory(file.Name()) {

			newUserHomeDirectory := userHomeDir + "/" + file.Name()

			if isTargetExists(newUserHomeDirectory) {

				newSourcePath := sourceLinkPath() + "/" + file.Name()

				filesFromDiretory := listFilestoLink(newSourcePath)

				for _, file := range filesFromDiretory {
					err := os.Symlink(newSourcePath+"/"+file.Name(), newUserHomeDirectory+"/"+file.Name())
					if err != nil {
						util.ExitWithError("Error creating symlink", err)
					}
				}

			} else {

				err := os.Symlink(linkSourcePath+"/"+file.Name(), userHomeDir+"/"+file.Name())
				if err != nil {
					util.ExitWithError("Error creating symlink", err)
				}
			}
			continue
		}

		err := os.Symlink(linkSourcePath+"/"+file.Name(), userHomeDir+"/"+file.Name())
		if err != nil {
			util.ExitWithError("Error creating symlink", err)
		}
	}
}

// isDirectory determines if a file represented
// by `path` is a directory or not
func isDirectory(path string) bool {
	fileInfo, _ := os.Stat(path)

	return fileInfo.IsDir()
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

// Do not create symlink from .glink-ignore
// TODO implement ignore list
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

	// TODO implement ignore list
	return make([]string, 0)
}
