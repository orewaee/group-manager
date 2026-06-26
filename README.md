# Group Manager

REST API service for managing hierarchical groups of people with recursive member counting. The service provides:

* Create, update, and delete groups with a hierarchical tree structure (parent/child groups)
* Create, update, and delete people (firstname, lastname, birthday) — assign to a group on creation, reassign on update
* List all groups with direct and total (including all descendants via recursive CTE) member counts
* List group members - either direct only or including all members from descendant subgroups
