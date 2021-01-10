# Todo app

## Tables Description
### Task

One task has many subtask

Primary key Id
| Id | Title | Content | Created_at | Updated_at | Deleted_at |
|:--:|:------|:-------:|:----------:|:----------:|:-----------:|
| int | varchar | varchar | timestamp | timestamp | timestamp |

### SubTask

One subtask has only one task

Primary key Id

Foreign key Task_id
| Id | Task_id | Title | Content | Created_at | Updated_at | Deleted_at |
|:--:|:--------|:-----:|:-------:|:----------:|:-----------:|:---------:|
| int | int | varchar | varchar | timestamp | timestamp | timestamp |


