# Changes

## [1.6.0](https://github.com/googleapis/google-cloud-go/compare/retail/v1.5.0...retail/v1.6.0) (2022-09-06)


### Features

* **retail:** release Control and ServingConfig serivces to v2 version feat: release AttributesConfig APIs to v2 version feat: release CompletionConfig APIs to v2 version feat: add local inventories info to the Product resource docs: Improved documentation for Fullfillment and Inventory API in ProductService docs: minor documentation format and typo fixes ([204b856](https://github.com/googleapis/google-cloud-go/commit/204b85632f2556ab2c74020250850b53f6a405ff))

## [1.5.0](https://github.com/googleapis/google-cloud-go/compare/retail/v1.4.0...retail/v1.5.0) (2022-08-02)


### Features

* **retail:** new model service to manage recommendation models feat: support case insensitive match on search facets feat: allow disabling spell check in search requests feat: allow adding labels in search requests feat: allow returning min/max values on search numeric facets feat: allow using serving configs as an alias of placements feat: allow enabling recommendation filtering on custom attributes feat: return output BigQuery table on product / event export response feat: allow skiping default branch protection when doing product full import docs: keep the API doc up-to-date with recent changes ([338d9c3](https://github.com/googleapis/google-cloud-go/commit/338d9c38b9c7f1b5e75493a2e3437c50785c561c))
* **retail:** support case insensitive match on search facets feat: allow disabling spell check in search requests feat: allow adding labels in search requests feat: allow returning min/max values on search numeric facets feat: allow using serving configs as an alias of placements feat: allow enabling recommendation filtering on custom attributes feat: return output BigQuery table on product / event export response docs: keep the API doc up-to-date with recent changes ([1d6fbcc](https://github.com/googleapis/google-cloud-go/commit/1d6fbcc6406e2063201ef5a98de560bf32f7fb73))
* **retail:** support case insensitive match on search facets feat: allow to return min/max values on search numeric facets feat: allow to use serving configs as an alias of placements docs: keep the API doc up-to-date with recent changes ([338d9c3](https://github.com/googleapis/google-cloud-go/commit/338d9c38b9c7f1b5e75493a2e3437c50785c561c))

## [1.4.0](https://github.com/googleapis/google-cloud-go/compare/retail/v1.3.0...retail/v1.4.0) (2022-06-01)


### Features

* **retail:** allow users to disable spell check in search requests feat: allow users to add labels in search requests docs: deprecate indexable/searchable on the product level custom attributes docs: keep the API doc up-to-date with recent changes ([5ed25c5](https://github.com/googleapis/google-cloud-go/commit/5ed25c5e2e40c7602802c35d61742631b619ed3c))

## [1.3.0](https://github.com/googleapis/google-cloud-go/compare/retail/v1.2.0...retail/v1.3.0) (2022-05-24)


### Features

* **retail:** start generating apiv2alpha and apiv2beta ([#6073](https://github.com/googleapis/google-cloud-go/issues/6073)) ([ec90f7b](https://github.com/googleapis/google-cloud-go/commit/ec90f7b224c67a02eb293224916c019054f5749d))

## [1.2.0](https://github.com/googleapis/google-cloud-go/compare/retail/v1.1.0...retail/v1.2.0) (2022-02-23)


### Features

* **retail:** set versionClient to module version ([55f0d92](https://github.com/googleapis/google-cloud-go/commit/55f0d92bf112f14b024b4ab0076c9875a17423c9))

## [1.1.0](https://github.com/googleapis/google-cloud-go/compare/retail/v1.0.0...retail/v1.1.0) (2022-02-14)


### Features

* **retail:** add file for tracking version ([17b36ea](https://github.com/googleapis/google-cloud-go/commit/17b36ead42a96b1a01105122074e65164357519e))

## 1.0.0

Stabilize GA surface.

## v0.1.0

This is the first tag to carve out retail as its own module. See
[Add a module to a multi-module repository](https://github.com/golang/go/wiki/Modules#is-it-possible-to-add-a-module-to-a-multi-module-repository).
