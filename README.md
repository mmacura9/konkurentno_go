# Concurrent programming in Go

This is a simple concurrent program. The Producer reads the data from a file and pushes it to the shared buffer. Consumers take an object from the shared buffer and process it. After processing the data, each consumer updates another shared object which is periodically printed by the Printer. After all consumers finish the data processing, Printer prints the final solution. 
