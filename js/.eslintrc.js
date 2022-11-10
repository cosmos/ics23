module.exports = {
  env: {
    es6: true,
    jasmine: true,
    node: true,
    worker: true,
  },
  parser: "@typescript-eslint/parser",
  parserOptions: {
    // ecmaVersion: 2018,
    project: "./tsconfig.json",
    tsconfigRootDir: __dirname,
  },
  plugins: ["@typescript-eslint", "prettier", "simple-import-sort", "import"],
  extends: [
    "eslint:recommended",
    "plugin:@typescript-eslint/recommended",
    "prettier",
    "plugin:prettier/recommended",
    "plugin:import/typescript",
  ],
  rules: {
    curly: ["warn", "multi-line", "consistent"],
    "no-bitwise": "warn",
    "no-console": ["warn", { allow: ["error", "info", "table", "warn"] }],
    "no-param-reassign": "warn",
    "no-shadow": "off", // disabled in favour of @typescript-eslint/no-shadow, see https://github.com/typescript-eslint/typescript-eslint/blob/master/packages/eslint-plugin/docs/rules/no-shadow.md
    "no-unused-vars": "off", // disabled in favour of @typescript-eslint/no-unused-vars, see https://github.com/typescript-eslint/typescript-eslint/blob/master/packages/eslint-plugin/docs/rules/no-unused-vars.md
    "prefer-const": "warn",
    radix: ["warn", "always"],
    "spaced-comment": ["warn", "always", { line: { markers: ["/ <reference"] } }],
    "import/no-cycle": "warn",
    "simple-import-sort/imports": "warn",
    "simple-import-sort/exports": "warn",
    "@typescript-eslint/array-type": ["warn", { default: "array-simple" }],
    "@typescript-eslint/await-thenable": "warn",
    "@typescript-eslint/ban-types": "warn",
    "@typescript-eslint/explicit-function-return-type": ["warn", { allowExpressions: true }],
    "@typescript-eslint/explicit-member-accessibility": "warn",
    "@typescript-eslint/naming-convention": [
      "warn",
      {
        selector: "default",
        format: ["strictCamelCase"],
      },
      {
        selector: "typeLike",
        format: ["StrictPascalCase"],
      },
      {
        selector: "enumMember",
        format: ["StrictPascalCase"],
      },
      {
        selector: "variable",
        format: ["strictCamelCase"],
        leadingUnderscore: "allow",
      },
      {
        selector: "parameter",
        format: ["strictCamelCase"],
        leadingUnderscore: "allow",
      },
    ],
    "@typescript-eslint/no-dynamic-delete": "warn",
    "@typescript-eslint/no-empty-function": "off",
    "@typescript-eslint/no-empty-interface": "off",
    "@typescript-eslint/no-explicit-any": "off",
    "@typescript-eslint/no-floating-promises": "warn",
    "@typescript-eslint/no-parameter-properties": "warn",
    "@typescript-eslint/no-shadow": "warn",
    "@typescript-eslint/no-unused-vars": ["warn", { argsIgnorePattern: "^_", varsIgnorePattern: "^_" }],
    "@typescript-eslint/no-unnecessary-type-assertion": "warn",
    // This one is used in CosmJS but would require a lot ot code change here
    // "@typescript-eslint/no-use-before-define": "off",
    "@typescript-eslint/prefer-readonly": "warn",
    // Re-activate once the codebase is migrates to ts-proto
    "@typescript-eslint/no-non-null-assertion": "off",
  },
  overrides: [
    {
      files: "**/*.js",
      rules: {
        "@typescript-eslint/no-var-requires": "off",
        "@typescript-eslint/explicit-function-return-type": "off",
        "@typescript-eslint/explicit-member-accessibility": "off",
        "@typescript-eslint/explicit-module-boundary-types": "off",
      },
    },
  ],
};
