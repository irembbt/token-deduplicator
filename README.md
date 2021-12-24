# Token Deduplicator

Generates 10 million (customizable) tokens consisting 7 letters which may contain duplicates.
This projects investigates 2 methods to deduplicate a long list of tokens and saves it to a relational database using gorm.

## How to Run

Run `docker-compose up` from the root of the project to build the source code into a docker image, and deploy it alongside a database.

## Performance

Two different deduplication algorithms are implemented.

- SortDeduplication runs in (nlogn) time but it uses constant memory. Since duplicated tokens should be a very small subset of all tokens I am assuming it's size to be constant.

- HashDeduplication runs in O(n) time but has to store every token in a map therefore uses O(n) space as well.

A small sized channel is used to communicate between deduplicator and database loader.

- This ensures that only a subset of the unique tokens are stored in memory at once, savinng memory.

`gorm.CreateInBatches` is used to parallelize insert statements, therefore increasing throughput, while also batching rows to insert.

# Database

There is only Tokens table to store deduplicate tokens including Token column. Token column is also primary key since they have to be unique.
