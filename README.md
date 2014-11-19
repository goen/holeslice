holeslice
=========

Higher order slices (HOS) provide an idiomatic pattern for implementing
a linear ordered set with a fast random deletion and insertion. 

Example:

	`var foo [2][2][]byte`
	`var bar []byte`

Foo is a quadratic slice, composed of 4 individual slices.
Foo can contain up to 3 holes, and may be used anywhere a bar can be used.

HOS would behave like a normal slice. You can access by offset,
append, reslice them and copy.

The additional features provided by the HOS are insert / delete to any location.

HOS are suitable to anywhere a programmer would need a linear ordered set, e.g.
to use a generic data structure, such as a balancing tree/trie or a list? rope?.

The best performance would be for an occasional inserts/deletes clustered around
a few specific spots.

The random access /read,write/ would be fast, just like in a regular array.

Use cases:
 * like a gap buffer
 * manipulating abstract syntax trees
 * a row in a text editor (cursor == the hole)
 * fifo, lifo, deque, ring buffer

Misuse cases
 * Priority Queue /heap/ (due to very frequent random insertion)

Performance specific additional operations
 * high degree HOS may need to be rebalanced
 * small or fragmented subslices or unused holes may need to be merged
