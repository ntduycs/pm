module.exports = {
  root: true,
  env: { browser: true, es2020: true },
  settings: {
    react: {
      version: 'detect',
    },
  },
  extends: [
    'eslint:recommended',
    'plugin:react/recommended',
    'plugin:@typescript-eslint/recommended',
    'plugin:react-hooks/recommended',
    'plugin:prettier/recommended',
    'plugin:@tanstack/eslint-plugin-query/recommended',
  ],
  ignorePatterns: ['dist', '.eslintrc.cjs'],
  parser: '@typescript-eslint/parser',
  parserOptions: {
    ecmaFeatures: { jsx: true },
    ecmaVersion: 2020,
    sourceType: 'module',
    project: './tsconfig.json',
  },
  plugins: ['react-refresh', '@tanstack/query'],
  rules: {
    'no-unused-vars': 'off',
    'react-refresh/only-export-components': [
      'warn',
      { allowConstantExport: true },
    ],
    'react/react-in-jsx-scope': 'off',
    '@typescript-eslint/no-unused-vars': 'off',
    'react/no-unescaped-entities': 'error',
    'react/no-unknown-property': 'error',
    '@typescript-eslint/no-unsafe-assignment': 'off',
    'react/display-name': 'error',
    '@typescript-eslint/consistent-type-assertions': ['off',
      {
        'assertionStyle': 'as' | 'angle-bracket',
        'objectLiteralTypeAssertions': 'allow' | 'allow-as-parameter',
      },
    ],
    '@typescript-eslint/no-redundant-type-constituents': 'error',
    '@typescript-eslint/naming-convention': 'off',
    '@typescript-eslint/no-floating-promises': 'error',
    '@typescript-eslint/restrict-template-expressions': ['error', {
      'allowNumber': true,
      'allowBoolean': true,
    }],
    '@typescript-eslint/unified-signatures': 'error',
    'react-hooks/exhaustive-deps': 'warn',
    'react-hooks/rules-of-hooks': 'error',
    '@tanstack/query/prefer-query-object-syntax': 'off',
  },
};
