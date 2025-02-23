package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
	"github.com/wb2-cli/config"
	"github.com/wb2-cli/utils"
)

func Create() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "Create a new project",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "sdk",
				Aliases: []string{"s"},
				Usage:   "Set the SDK path",
			},
		},
		Action: doCreate,
	}
}

func doCreate(ctx *cli.Context) error {
	projectName := inputString("Create New Project:")
	projectPATH := projectName + "/" + projectName

	_, confErr := config.ReadConfig(config.GetLocations())
	sdk := ctx.String("sdk")
	if sdk == "" {
		if confErr != nil {
			utils.Exit1With("SDK is not configured, run 'wb2-cli init'")
			return nil
		}
	}

	err := os.MkdirAll(projectPATH, 0755) // 创建多级目录
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		// 写入文件内容
		fileMakefile := projectName + "/Makefile"
		fileConfig := projectName + "/proj_config.mk"
		filemk := projectPATH + "/bouffalo.mk"
		filemain := projectPATH + "/main.c"
		CopyFile(projectName, fileMakefile, ctx)
		CopyFile(projectName, fileConfig, ctx)
		CopyFile(projectName, filemk, ctx)
		CopyFile(projectName, filemain, ctx)
		fmt.Printf("%s Created Successfully\r\n", projectName)
	}
	return nil
}

func getTemplate(pattern string) string {
	// 创建映射
	templateVars := map[string]string{
		"Makefile":       utils.TempMakeFile,
		"proj_config.mk": utils.TempProjectConfig,
		"bouffalo.mk":    utils.TempBouffalo,
		"main.c":         utils.TempMain,
	}

	// 根据 pattern 选择对应的模板
	if template, exists := templateVars[pattern]; exists {
		return template
	}
	return "No matching template found."
}

func CopyFile(proj_name string, filepath string, ctx *cli.Context) {
	fileName := getFileName(filepath)
	fileTemp := getTemplate(fileName)
	conf, confErr := config.ReadConfig(config.GetLocations())
	sdk := ctx.String("sdk")
	if sdk == "" {
		if confErr != nil {
			utils.Exit1With("SDK is not configured, run 'wb2-cli init'")
			return
		}
		sdk = conf.SDK
	}

	filecontect := strings.ReplaceAll(fileTemp, "$(PROJECTNAME)", proj_name)
	filecontect = strings.ReplaceAll(filecontect, "$(SDK)", sdk)
	err := os.WriteFile(filepath, []byte(filecontect), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func getFileName(input string) string {
	return filepath.Base(input)
}
