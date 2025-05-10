# taskManagerCLI

Task Manager CLI coded on Go lang with PostgreSQL database.
Using CLI with this flags:
-title "Title"
-desc "Description" (Optional)
-priority (low, medium, high) default "low"
-edit id(0,1,2...)
-complete id 
-del id
-list

Filter using:
-filter priority -priority "high" | lists all tasks with priority "high"
-filter completed | lists all completed tasks
