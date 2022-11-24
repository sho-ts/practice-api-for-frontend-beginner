const createUseCase = require('./.console/createUseCase');
const createRepository = require('./.console/createRepository');
const args = process.argv.filter((_, i) => i > 1);
const [action] = args;

switch (action) {
  case 'usecase':
  case 'u':
    createUseCase(args);
    break;
  case 'repository':
  case 'r':
    createRepository(args);
    break;
}