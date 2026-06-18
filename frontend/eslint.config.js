import js from '@eslint/js'
import globals from 'globals'
import tseslint from 'typescript-eslint'
import react from 'eslint-plugin-react'
import reactHooks from 'eslint-plugin-react-hooks'
import reactRefresh from 'eslint-plugin-react-refresh'
import importPlugin from 'eslint-plugin-import'
import jsxA11y from 'eslint-plugin-jsx-a11y'
import prettier from 'eslint-plugin-prettier'
import { defineConfig, globalIgnores } from 'eslint/config'

export default defineConfig([
  globalIgnores(['dist']),
  js.configs.recommended,
  {
    extends: [...tseslint.configs.recommendedTypeChecked],
    files: ['**/*.{ts,tsx}'],
    plugins: {
      react,
      'react-hooks': reactHooks,
      'react-refresh': reactRefresh,
      import: importPlugin,
      'jsx-a11y': jsxA11y,
      prettier,
    },
    languageOptions: {
      globals: globals.browser,
      parserOptions: {
        project: ['./tsconfig.app.json', './tsconfig.node.json'],
        tsconfigRootDir: import.meta.dirname,
      },
    },
    settings: {
      react: {
        version: 'detect',
      },
      'import/resolver': {
        typescript: true,
      },
    },
    rules: {
      //
      // TypeScript
      //
      'no-unused-vars': 'off',
      '@typescript-eslint/no-unused-vars': [
        'error',
        {
          ignoreRestSiblings: true,
          argsIgnorePattern: '^_',
          varsIgnorePattern: '^_',
        },
      ],
      '@typescript-eslint/consistent-type-imports': [
        'error',
        {
          prefer: 'type-imports',
        },
      ],
      '@typescript-eslint/no-floating-promises': 'error',
      '@typescript-eslint/no-misused-promises': 'error',
      '@typescript-eslint/await-thenable': 'error',
      '@typescript-eslint/no-unnecessary-condition': 'error',
      '@typescript-eslint/switch-exhaustiveness-check': 'error',
      //
      // React
      //
      'react/react-in-jsx-scope': 'off',
      'react/prop-types': 'off',

      'react/jsx-no-useless-fragment': [
        'error',
        { allowExpressions: true },
      ],

      'react/no-unstable-nested-components': [
        'warn',
        { allowAsProps: true },
      ],
      'react/no-array-index-key': 'warn',
      'react-hooks/rules-of-hooks': 'error',
      'react-hooks/exhaustive-deps': 'warn',
      //
      // Import order
      //
      'import/order': [
        'error',
        {
          groups: [
            ['builtin', 'external'],
            'internal',
            ['parent', 'sibling'],
            'index',
          ],
          'newlines-between': 'always',
          alphabetize: {
            order: 'asc',
            caseInsensitive: true,
          },
        },
      ],
      'import/no-cycle': ['error', { maxDepth: 1 }],
      //
      // Accessibility
      //
      'jsx-a11y/alt-text': 'error',
      'jsx-a11y/anchor-has-content': 'error',
      'jsx-a11y/aria-props': 'error',
      'jsx-a11y/aria-role': 'error',
      'jsx-a11y/aria-unsupported-elements': 'error',
      'jsx-a11y/label-has-associated-control': 'error',
      //
      // General
      //
      curly: ['error', 'all'],
      'no-console': [
        'warn',
        {
          allow: ['warn', 'error'],
        },
      ],
      'no-param-reassign': [
        'error',
        {
          props: false,
        },
      ],
      'no-unneeded-ternary': 'error',
      //
      // Formatting (prettier)
      //
      'prettier/prettier': [
        'error',
        {
          singleQuote: true,
          semi: false,
          arrowParens: 'avoid',
          trailingComma: 'es5',
          printWidth: 100
        },
      ], 
    },
  },
])