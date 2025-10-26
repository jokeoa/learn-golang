## Pattern Overview

The Strategy pattern defines a family of algorithms, encapsulates each one, and makes them interchangeable. It lets the algorithm vary independently from clients that use it.

### Architecture Components

**1. SortingAlgo Interface**
- Defines the contract all strategies must follow
- Single method: `sort(s *Sorter)`
- Enables polymorphism and loose coupling

**2. Sorter (Context)**
- Holds the data to be sorted
- Maintains a reference to the current strategy
- Delegates sorting work to the strategy
- Provides methods to switch strategies dynamically

**3. Concrete Strategies**
- **BubbleSort**: O(n²) comparison-based algorithm for small datasets
- **QuickSort**: O(n log n) average time, pivot-based partitioning
- **MergeSort**: O(n log n) guaranteed time, divide-and-conquer approach

## Pattern Trade-offs

### Advantages
- High flexibility and extensibility
- Clean code organization
- Easy testing of individual strategies
- Reduced conditional complexity

### Disadvantages
- Increases number of classes/objects
- Clients must understand different strategies
- Communication overhead between context and strategy
- May be overkill for simple cases
