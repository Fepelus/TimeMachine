Time Machine
============

This is a little utility for backing up a hard-drive.

What does it do?
----------------

Creates a directory for your snapshot that holds a directory structure
that matches your source directory structure. Every file in the destination
directory structure is a hard link to a file in the object store.

The object store is a directory full of every version of every file you store
where the name is the SHA1 hash of the data-content of the file.

If in two different snapshots, made by running this utility two different times,
you save the same file twice, it will actually be stored in the object store only
once and the hard links in the snapshot directory tree will both point to the
same file.

To do
-----

- There is a memory leak that needs plugging.

License
-------

MIT
