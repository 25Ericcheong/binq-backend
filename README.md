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
- Decided to go with Domain Driven Development with an emphasis on my project reflecting the domain's needs and requirements
- [Understanidng receiver value methods and receiver pointer methods](https://go.dev/wiki/MethodSets#:~:text=The%20method%20set%20of%20the,must%20have%20a%20unique%2). Best read about understanding receivable values and receivable pointers and how it relates to interfaces [highlights how values are stored in interfaces when they are assigned to an interface](https://npf.io/2014/05/intro-to-go-interfaces/) and the interface's methodsa are called within another method. The problem is compiler not able to automatically find the pointer for the stored value on the interface (explains why interfaces are not addressable) - it needs this pointer for the value because the method is expecting a receivable pointer type. It also highlights how a pointer to a nil can be assigned to an interface's value (which makes the interface value output as not nil) which is a bug when we expect the output to be a nil but because the nil pointer is stored as the interface's value - the interface itself (the data specifically) is storing a pointer to a nil; making the interface type as not nil.
