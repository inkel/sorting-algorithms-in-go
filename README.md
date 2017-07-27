# Sorting Algorithms Benchmarks

Based on the [Sorting Algorithms Primer](https://hackernoon.com/sorting-algorithms-primer-374b83f3ba09) by [Richar Knop](https://twitter.com/richardknop).

## Running the benchmarks

You can select how many elements the array to be sorted has by changing the `SIZE` environment variable. If left without value it will use the same array as in Richard's article.

## Known issues
* With `SIZE=1` many of the algorithms fail.
* `MergeSort` doesn't sort the slice in place but returns a new one.
* Days only have 24 hours and I've a job, you know?

## License
Do what you want. These are sorting algorithms benchmarks.