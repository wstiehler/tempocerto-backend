
**Level 1: System Context diagram**

"Tempo Certo is an application that offers a scheduling service for available time slots for companies."

_See more in our [Readme](https://github.com/wstiehler/tempocerto-backend)_

**Scope**: Tempo Certo is an application that offers a scheduling service for available time slots for companies. The goal is to allow clients to schedule appointments with companies easily and efficiently.

**Primary elements**: TempoCerto-API app.

**Intended audience**: Everyone, technical and non-technical, inside and outside the software development team.

**Functional Requirement**
The TempoCerto-APP is responsible for managing companies and schedules for users in the system. It should provide the following functionalities:

* Companies can register on the platform by providing their name and CNPJ.
* Companies can define weekly available time slots for scheduling, specifying the start and end period as well as the weekdays when the slots are available.
* Companies can fill specific slots by date and time with details about the appointment, such as title and company ID.
* Clients can view available slots for scheduling.
* Companies can view an agenda of all scheduled appointments, including details about the clients and the time slots.

***Success Scenarios***

* A company successfully registers on the platform, providing its name and CNPJ. A record is created in the database, and a unique ID is assigned to the company.
* A company defines weekly available time slots for scheduling. The slots are created in the database based on the provided information and are available for clients to schedule appointments.
* A company fills a specific slot by date and time with details of the appointment. The slot is updated in the database as unavailable and contains information about the appointment.
* A client queries available slots for scheduling. The application returns a list of available slots with information about the company and the time slots.
* A company queries its agenda to see all scheduled appointments. The application returns a list of all appointments with details about the clients and the time slots.

***Failure Scenarios***

* The registration of a company fails due to invalid data validation. The application returns an error message indicating the registration failure.
* Whenn defining weekly available times, a company specifies times that are outside the standard working hours (e.g., before 8:00 AM or after 6:00 PM). The application validates the times and returns an error message indicating that the times are outside working hours.
* A company tries to fill a slot by date and time, but an error occurs when updating the slot in the database. The application returns an error message indicating the failure to fill the slot.
* A client tries to query available slots, but an error occurs when retrieving data from the database. The application returns an error message indicating the failure to query available slots.
* A company tries to query its agenda, but an error occurs when retrieving data from the database. The application returns an error message indicating the failure to query the agenda.
