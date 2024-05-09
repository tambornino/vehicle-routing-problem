# Vehicle Routing Problem (VRP)

### How do I run this?
From the command line, simply use the binary executable plus an arg of the data set file path.\
The existing samples are in the data directory.\
For example:
`./bin data/problem1.txt`

### What if I want to make changes to the executable?
Simply run `go build -o bin` and the binary will be rebuilt.

### What am I seeing on the printout?
Each line presents a driver's list of loads completed.\
The number of lines is the number of drivers used.

### Can you explain the problem in more detail?
The VRP specifies a set of loads to be completed efficiently by an unbounded number of drivers.

Each load has a pickup location and a dropoff location, each specified by a Cartesian point. A driver completes a load by driving to the pickup location, picking up the load, driving to the dropoff, and dropping off the load. The time required to drive from one point to another, in minutes, is the Euclidean distance between them. That is, to drive from (x1, y1) to (x2, y2) takes sqrt((x2-x1)^2 + (y2-y1)^2) minutes.

As an example, suppose a driver located at (0,0) starts a load that picks up at (50,50) and delivers at (100,100). This would take 2*sqrt(2*50^2) = ~141.42 minutes of drive time to complete: sqrt((50-0)^2 + (50-0)^2) minutes to drive to the pickup, and sqrt((100-50)^2 + (100-50)^2) minutes to the dropoff.

Each driver starts and ends his shift at a depot located at (0,0). A driver may complete multiple loads on his shift, but may not exceed 12 hours of total drive time. That is, the total Euclidean distance of completing all his loads, including the return to (0,0), must be less than 12*60.

A VRP solution contains a list of drivers, each of which has an ordered list of loads to be completed. All loads must be assigned to a driver.

### What is the basis for this solution?
It was inspired by the 'Nearest Neighbor' solution explained here: https://iopscience.iop.org/article/10.1088/1742-6596/2421/1/012027/pdf
