# Binq Backend
Current purpose of having a backend is to allow customers to queue digitally. This allows branches to better manage customers. Current success list for this project will be:

- Creating a simple wireframe to understand domain requirements
- Allow customers to create a ticket to queue. After ticket is generated, display queue number and position to customer for reference later on
- Allow internal staff to have overview of queues (each branch will have its own internal dashboard view)
- Internal dashboard should display information submitted by customers
- Internal staff should be able to modify existing tickets on dashboard
- Internal staff should be able to remove any tickets on dashboard
- There should be a refresh button that pulls all tickets currently existing for up-to-date information for staff reference

# Key notes
These are notes that I found while researching on how to best maintain certain aspects of the backend
- Datahase migration will rely solely on golang-migrate package. Will use a script to run the migrations and prior to actual deployment to production; will attempt to build docker container with postgresql and test if migration has taken place
- [Postgres row data conversion to golang struct](https://stackoverflow.com/questions/17265463/how-do-i-convert-a-database-row-into-a-struct) - is there a better way of doing this without relying on external libraries?
- ~Timezone for local dev database [has been set](https://stackoverflow.com/questions/6663765/postgres-default-timezone) to `Asia/Kuala_Lumpur`. Will need to do the same for production database too~ Have reset back to default setting
