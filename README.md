# go-shopify-graphql-model

This is a simple library to help you use Shopify's GraphQL objects in your Go code.

## Getting started

0. Install dependencies

    ```bash
    yarn
    ```

1. Fetch the Shopify graphql schema

    ```bash
    STORE=my-store ACCESS_TOKEN=shpat_xxxxx API_VERSION=2023-07 yarn fetch
    ```

2. Remove the following declaration from `schema.graphql` so that models can be generated

    ```graphql
    schema {
        query: QueryRoot
        mutation: Mutation
    }
    ```

3. Generate models

    ```bash
    go run main.go
    ```
