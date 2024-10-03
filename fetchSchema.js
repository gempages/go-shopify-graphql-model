import { promises } from "fs"
import fetch from "node-fetch"
import { getIntrospectionQuery, printSchema, buildClientSchema } from "graphql"

async function main() {
    const introspectionQuery = getIntrospectionQuery()

    const response = await fetch(`https://${process.env.STORE}.myshopify.com/admin/api/${process.env.API_VERSION}/graphql.json`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "X-Shopify-Access-Token": process.env.ACCESS_TOKEN,
        },
        body: JSON.stringify({ query: introspectionQuery }),
    })

    const result = await response.json()
    if (result.errors) {
        console.error(result.errors)
        return
    }

    const schema = buildClientSchema(result.data)

    const outputFile = "./schema.graphql"

    await promises.writeFile(outputFile, printSchema(schema))
    await promises.writeFile(`./${process.env.API_VERSION}.json`, JSON.stringify(result)) // Use this schema file for GemPages v6
}

main()
