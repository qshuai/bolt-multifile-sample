# bolt-multifile-sample

[bolt](https://github.com/boltdb/bolt) using a single file to store k/v pairs with same bucket.

A blockchain named [Sia](https://github.com/NebulousLabs/Sia) using bolt db only has a single file for storing blockchain data. With data becoming big, the db performance will be down. 

So this repository give an example to store data with same bucket and path using different files.

### Usage:

The test file has given you best practice

### Todo:

- monitor the size of the last *.db, and create a new bolt db reference when reach the specified limit
- data migration when the size of a *.db file is to small


