# User Create APP


**Level 1: System Context diagram**

"User Create API is a robust web service designed to handle user management operations with various levels of access"

_See more in our [Readme](https://github.com/wstiehler/zpe-project)_

**Scope**: The User Create API is a robust web service designed to handle user management operations with various levels of access. It provides functionalities to create with each associated with different types of access permissions.

**Primary elements**: User Create-API app.

**Intended audience**: Everyone, technical and non-technical, inside and outside the software development team.


**Functional Requirement**

The User Create APP should provide an API to create a new user with user details, including Id, Name, Role, and Email. The API should be able to handle the following cases:

* Receive a request to create a new user.
* Validate the mandatory fields: Name, Role, and Email.
* Create a new user in the system with the provided details.
* Return a status code and a message indicating the result of the user creation.

***Success Scenarios***

* The user provides all mandatory fields: Name, Role, and Email.
* The user provides valid and unique information for the new user.
* The system successfully creates the new user.
* The API returns a status code 200 (OK) and a message indicating that the user has been successfully created.

***Failure Scenarios***

* The user does not provide all mandatory fields: Name, Role, and Email.
* The user provides invalid or duplicate information for the new user.
* The system encounters a conflict when trying to create the new user because the user already exists in the system.
* The API returns an error status code (e.g., 400 or 409) and a message indicating the reason for the failure in creating the user.