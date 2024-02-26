package cmds

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
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

		err := os.Symlink(linkSourcePath+"/"+file.Name(), userHomeDir+"/"+file.Name())
		if err != nil {
			fmt.Println("Error creating symlink ", file.Name())
		}
	}
}

// Get user home path
func userHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Unable to get home directory")
	}

	return home
}

// Get current path
func sourceLinkPath() string {
	linkSourcePath, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to get current directory")
	}

	return linkSourcePath
}

// Get list of files and folder to create symlink
func listFilestoLink(sourcePath string) []fs.DirEntry {
	files, err := os.ReadDir(sourcePath)
	if err != nil {
		log.Fatal("Unable to get a list of files")
	}

	return files
}

// Do not create symlink from .glink-ignore
func ignoreList(sourcePath string) []string {
	_, err := os.Stat(sourcePath + "/.glink-ignore")

	if errors.Is(err, os.ErrNotExist) {
		return make([]string, 0)
	}

	ignoreFile := sourcePath + "/.glink-ignore"

	fileToIgnore, err := os.OpenFile(ignoreFile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal("Unable to read .glink-ignore file")
	}

	defer fileToIgnore.Close()

	content := bufio.NewReader(fileToIgnore)
	for {
		line, err := content.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("read file line error: %v", err)
		}
		_ = line // GET the line string
	}

	// TODO implement ignore list
	return make([]string, 0)
}
