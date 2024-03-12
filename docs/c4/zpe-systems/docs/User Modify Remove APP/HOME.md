# User Modify Remove APP

**Level 1: System Context diagram**

"User Modify Deleter APP is a web service designed to handle user management operations with different access levels"

_See more in our [Readme](https://github.com/wstiehler/zpe-project)_

**Scope**: The User Modify Deleter APP is a web service designed to handle user management operations with different access levels. It provides functionalities to modify, delete, and view users based on the user's role and permissions.

**Primary elements**: User Modify Deleter APP.

**Intended audience**: Everyone, technical and non-technical, inside and outside the software development team.

**Functional Requirement**
The User Modify Deleter APP should provide an API to modify user roles and delete users from the system. It should adhere to the following rules:

* A user can have multiple roles:
* Admin can be Modifier and/or Watcher.
* Modifier can also be a Watcher.
* Watcher cannot assume any other role.* If user details are provided, return details for the specified user.
* The API should allow updating or adding roles for a user based on the rules above.
* If the rules are followed during role modification, the API should return success; otherwise, it should return failure with an error message.
* The API should provide functionality to delete any user from the system.
* If the user to be deleted does not exist, the API should return failure with an error message.

***Success Scenarios***

* The API allows updating or adding roles for a user according to the specified rules.
* The user's roles are successfully modified or added, and the API returns a success status.
* The API allows deleting a user from the system, and the user is successfully deleted.
* The user to be deleted exists in the system, and the deletion operation is successful.

***Failure Scenarios***

* The API encounters an error while updating or adding roles for a user.
* The provided user roles do not adhere to the specified rules, and the API returns a failure status with an error message.
* The API encounters an error while deleting a user from the system.
* The user to be deleted does not exist in the system, and the API returns a failure status with an error message.