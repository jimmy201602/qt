package templater

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
)

func GenModule(m, target string, mode int) {
	if os.Getenv("QT_DUMP") == "true" {
		defer parser.Dump()
	}

	if !parser.ShouldBuildForTarget(m, target) {
		utils.Log.WithField("module", m).Debug("skip generation")
		return
	}
	utils.Log.WithField("module", m).Debug("generating")

	var suffix string
	switch m {
	case "AndroidExtras":
		suffix = "_android"

	case "Sailfish":
		suffix = "_sailfish"
	}

	if mode == NONE {
		utils.RemoveAll(utils.GoQtPkgPath(strings.ToLower(m)))
		utils.MkdirAll(utils.GoQtPkgPath(strings.ToLower(m)))
	}

	if m == "AndroidExtras" {
		utils.Save(utils.GoQtPkgPath(strings.ToLower(m), "utils-androidextras_android.go"), utils.Load(filepath.Join(strings.TrimSpace(utils.GoListOptional("{{.Dir}}", "github.com/therecipe/qt/internal", "-find", "get files dir")), "/binding/files/utils-androidextras_android.go")))
	} else if m == "Qml" {
		utils.Save(utils.GoQtPkgPath(strings.ToLower(m), "utils-qml.go"), utils.Load(filepath.Join(strings.TrimSpace(utils.GoListOptional("{{.Dir}}", "github.com/therecipe/qt/internal", "-find", "get files dir")), "/binding/files/utils-qml.go")))
	}

	if mode == MINIMAL {
		if suffix != "" {
			return
		}

		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+"-minimal.cpp"), CppTemplate(m, mode, target, ""))
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+"-minimal.h"), HTemplate(m, mode, ""))
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+"-minimal.go"), GoTemplate(m, false, mode, m, target, ""))

		if !UseStub(false, "Qt"+m, mode) {
			CgoTemplate(m, "", target, mode, m, "")
		}

		return
	}

	if !UseStub(false, "Qt"+m, mode) {
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+suffix+".cpp"), CppTemplate(m, mode, target, ""))
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+suffix+".h"), HTemplate(m, mode, ""))
	}

	//always generate full
	if suffix != "" {
		utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+suffix+".go"), GoTemplate(m, false, mode, m, target, ""))
	}

	//may generate stub
	utils.SaveBytes(utils.GoQtPkgPath(strings.ToLower(m), strings.ToLower(m)+".go"), GoTemplate(m, suffix != "", mode, m, target, ""))

	if !UseStub(false, "Qt"+m, mode) {
		CgoTemplate(m, "", target, mode, m, "")
	}

	if utils.QT_DOCKER() {
		if idug, ok := os.LookupEnv("IDUG"); ok {
			if path := utils.GoQtPkgPath(strings.ToLower(m)); utils.UseGOMOD(path) {
				utils.RunCmd(exec.Command("chown", "-R", idug, filepath.Dir(utils.GOMOD(path))), "chown files to user")
			} else {
				utils.RunCmd(exec.Command("chown", "-R", idug, path), "chown files to user")
			}
		}
	}
}
