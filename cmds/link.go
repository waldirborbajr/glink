package cmds

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

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

	// Load ignore patterns from .ignore file
	ignorePatterns, err := loadIgnorePatterns(linkSourcePath)
	if err != nil {
		fmt.Println("Error loading ignore patterns:", err)
		return
	}

	filesToLink := listFilestoLink(linkSourcePath)

	for _, file := range filesToLink {

		fileName := file.Name()

		// ignoring the .glink-ignore file
		if file.Name() == ".glink-ignore" || fileName == "glink-ignore" {
			continue
		}

		if isIgnored(fileName, ignorePatterns) {
			fmt.Printf("Ignoring file: %s\n", fileName)
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

// loadIgnorePatterns loads patterns from an ignore file.
func loadIgnorePatterns(sourcePath string) (map[string]struct{}, error) {
	patterns := make(map[string]struct{})

	_, err := os.Stat(sourcePath + "/.glink-ignore")

	if errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	ignoreFile := sourcePath + "/.glink-ignore"

	file, err := os.Open(ignoreFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			patterns[line] = struct{}{}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return patterns, nil
}

// isIgnored checks if the file matches any ignore pattern.
func isIgnored(path string, ignorePatterns map[string]struct{}) bool {
	filename := filepath.Base(path)
	_, ignored := ignorePatterns[filename]
	return ignored
}
