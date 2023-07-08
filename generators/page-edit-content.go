package generators

import (
	"fmt"
	"janic0/gemserv/redis"
	"janic0/gemserv/utils"
	"strings"
)

func GetPageEditContent(path string) ([]byte, error) {
	messages := make([]string, 0)
	messages = append(
		messages,
		utils.CraftAdminLink("/admin", "Back to overview"),
		"",
		"# "+path,
		"------------",
	)
	items, err := redis.Client.LRange("page:"+path, 0, -1).Result()
	escapedPath := utils.EscapePathSegments(path)
	if err != nil {
		return nil, err
	} else if len(items) == 0 {
		return nil, fmt.Errorf("No blocks found.")
	}
	for i, item := range items {
		if i == 0 {
			if item == "enabled" {
				messages = append(messages, utils.CraftAdminLink("/admin/page"+escapedPath+"/disable", "Disable this page"))
			} else {
				messages = append(messages, utils.CraftAdminLink("/admin/page"+escapedPath+"/enable", "Enable this page"))
			}
			messages = append(messages,
				utils.CraftAdminLink("/admin/page"+escapedPath+"/move", "Move this page"),
				utils.CraftAdminLink("/admin/page"+escapedPath+"/duplicate", "Duplicate this page"),
				utils.CraftAdminLink("/admin/page"+escapedPath+"/delete", "Delete this page"),
				"------------",
				"")
		} else {
			messages = append(messages,
				// LTR Character for escaping
				string([]byte{226, 128, 142})+item,
				utils.CraftAdminLink(fmt.Sprintf("/admin/page%s/blocks/%d/edit", escapedPath, i), "Edit"),
				utils.CraftAdminLink(fmt.Sprintf("/admin/page%s/blocks/%d/insert-above", escapedPath, i), "Insert Above"),
				utils.CraftAdminLink(fmt.Sprintf("/admin/page%s/blocks/%d/remove", escapedPath, i), "Delete"),
				"---",
			)
		}
	}
	messages = append(messages, "", utils.CraftAdminLink(fmt.Sprintf("/admin/page%s/add-block", escapedPath), "+ Add block"))
	return []byte(strings.Join(messages, "\r\n")), nil
}
