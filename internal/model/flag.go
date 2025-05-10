package model

import (
	"flag"
)

type Flags struct {
	Title    *string
	Desc     *string
	Priority *Priority
	Add      *bool
	Edit     *int
	Del      *int
	Complete *int
	Filter   *string
	List     *bool
}

// Parsing CLI flags for actions with tasks
func InitFlags() Flags {
	title := flag.String("title", "", "Title of the task")
	desc := flag.String("desc", "", "Descriptions of the task")
	priority := flag.String("priority", "low", "Priority: low, medium, high")
	add := flag.Bool("add", false, "Add new task")
	del := flag.Int("del", -1, "Delete task by id")
	complete := flag.Int("complete", -1, "Delete task by id")
	filter := flag.String("filter", "", "Print filtered task by completed, -priority")
	list := flag.Bool("list", false, "Print list of all tasks")
	edit := flag.Int("edit", -1, "Edit specify id new title and description")

	flag.Parse()

	return Flags{
		Title:    title,
		Desc:     desc,
		Priority: ParsePriority(*priority),
		Add:      add,
		Edit:     edit,
		Del:      del,
		Complete: complete,
		Filter:   filter,
		List:     list,
	}
}
