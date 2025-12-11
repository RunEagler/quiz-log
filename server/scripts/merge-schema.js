const { mergeTypeDefs } = require('@graphql-tools/merge');
const { loadFilesSync } = require('@graphql-tools/load-files');
const { print } = require('graphql');
const fs = require('fs');
const path = require('path');

// Load all GraphQL schema files
const typesArray = loadFilesSync(path.join(__dirname, '../graph/schema'), {
  extensions: ['graphqls']
});

// Merge all type definitions
const mergedSchema = mergeTypeDefs(typesArray);

// Convert to string
const printedSchema = print(mergedSchema);

// Write to shared/schema directory
const outputPath = path.join(__dirname, '../../web/schema/schema.graphql');
const outputDir = path.dirname(outputPath);

// Ensure directory exists
if (!fs.existsSync(outputDir)) {
  fs.mkdirSync(outputDir, { recursive: true });
}

fs.writeFileSync(outputPath, printedSchema);

console.log(`✅ Schema merged successfully to ${outputPath}`);
