import nextJest from 'next/jest.js'

const createJestConfig = nextJest({
  dir: './',
})

const customJestConfig = {
  setupFilesAfterEnv: ['<rootDir>/jest.setup.js'],
  testEnvironment: 'jsdom',
  testpathIgnorePatterns: ['<rootDir>/.next.js', '<rootDir>/node_modules/'],
  moduleNameMapper: {
    '^@(.*)$': '<roorDir>/src/$1',
  },
  collectCoverageFrom: [
    'src/**/*.{js, jsx, ts, tsx}',
    '!src/**/*.d.ts',
    '!src/**/*.stories.{js, jsx, ts, tsx}',
  ],
  converageThreshold: {
    global: {
      branches: 50,
      functions: 50,
      lines: 50,
      statements: 50,
    },
  },
}

export default createJestConfig(customJestConfig)

