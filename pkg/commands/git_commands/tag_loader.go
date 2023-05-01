package git_commands

import (
	"regexp"

	"github.com/jesseduffield/generics/slices"
	"github.com/BSteffaniak/lazyaws/pkg/commands/models"
	"github.com/BSteffaniak/lazyaws/pkg/commands/oscommands"
	"github.com/BSteffaniak/lazyaws/pkg/common"
	"github.com/BSteffaniak/lazyaws/pkg/utils"
)

type TagLoader struct {
	*common.Common
	cmd oscommands.ICmdObjBuilder
}

func NewTagLoader(
	common *common.Common,
	cmd oscommands.ICmdObjBuilder,
) *TagLoader {
	return &TagLoader{
		Common: common,
		cmd:    cmd,
	}
}

func (self *TagLoader) GetTags() ([]*models.Tag, error) {
	// get remote branches, sorted  by creation date (descending)
	// see: https://git-scm.com/docs/git-tag#Documentation/git-tag.txt---sortltkeygt
	tagsOutput, err := self.cmd.New(`git tag --list -n --sort=-creatordate`).DontLog().RunWithOutput()
	if err != nil {
		return nil, err
	}

	split := utils.SplitLines(tagsOutput)

	lineRegex := regexp.MustCompile(`^([^\s]+)(\s+)?(.*)$`)

	tags := slices.Map(split, func(line string) *models.Tag {
		matches := lineRegex.FindStringSubmatch(line)
		tagName := matches[1]
		message := ""
		if len(matches) > 3 {
			message = matches[3]
		}

		return &models.Tag{
			Name:    tagName,
			Message: message,
		}
	})

	return tags, nil
}
