# BranchKV

Copy-on-Write Virtual Memory Runtime for Tree-of-Thought AI Inference.

## Features

- O(1) sequence branching
- Copy-on-write token pages
- Atomic ref-counted memory
- Radix trie prefix compression
- Zero-allocation hot paths
- Deterministic isolation

## Architecture

BranchKV virtualizes token memory pages similarly to an operating system MMU.

Logical branches share ancestral pages until divergence occurs.

## Stack

- Go
- sync/atomic
- unsafe.Pointer
- radix trie
- custom memory pools

