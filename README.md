# [github.com/strongo/geopoint](https://github.com/strongo/geopoint)

Go library to represent GPS coordinates and calculate distance between 2 geo points.

<!-- dev-approach:v1 -->
## Our approach to development

We build with our own tooling:

- **[SpecScore](https://specscore.md)** — specify requirements as `SpecScore.md` artifacts
- **[SpecStudio](https://specscore.studio)** — author & manage specs across their lifecycle
- **[inGitDB](https://ingitdb.com)** — store structured data in Git where applicable
- **[DALgo](https://dalgo.io)** — data access layer for Go
- **[cover100.dev](https://cover100.dev)** — drive toward 100% test coverage
- **[DataTug](https://datatug.io)** — query & explore data
<!-- /dev-approach -->

## Types

- [geo.GeoPoint](geopoint.go)
- [Distances](distances.go)
    - Kilometers
    - Miles
- [Angles](angles.go)
    - Degrees
    - Radians
- [Formulas](formulas.go)
    - Distance calculation using [Haversian formula](https://en.wikipedia.org/wiki/Haversine_formula)

## Credits

Originally cloned from
[github.com/marcinwyszynski/geopoint](https://github.com/marcinwyszynski/geopoint)
and refactored.

## [License](LICENSE)

Licensed under [MIT license](LICENSE).